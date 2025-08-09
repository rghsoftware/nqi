package handler

import (
	"context"
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/rghsoftware/nqi/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/argon2"
)

// HashParams holds the parameters for the Argon2id hashing algorithm.
// These are recommended baseline settings.
type HashParams struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

// DefaultHashParams returns secure default parameters for Argon2
func DefaultHashParams() *HashParams {
	return &HashParams{
		Memory:      64 * 1024, // 64 MB
		Iterations:  3,         // 3 iterations
		Parallelism: 2,         // 2 threads
		SaltLength:  16,        // 16 bytes salt
		KeyLength:   32,        // 32 bytes key
	}
}

type UserHandler struct {
	DB *pgxpool.Pool
}

type RegisterUserInput struct {
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	PublicKey string `json:"publicKey" binding:"required"`
}

// HashPassword generates a salted hash of the password using Argon2id
func HashPassword(password string, params *HashParams) (string, error) {
	// Generate a random salt
	salt := make([]byte, params.SaltLength)
	if _, err := rand.Read(salt); err != nil {
		return "", fmt.Errorf("failed to generate salt: %w", err)
	}

	// Generate the hash using Argon2id
	hash := argon2.IDKey(
		[]byte(password),
		salt,
		params.Iterations,
		params.Memory,
		params.Parallelism,
		params.KeyLength,
	)

	// Encode the result as: $argon2id$v=19$m=memory,t=iterations,p=parallelism$salt$hash
	encodedSalt := base64.RawStdEncoding.EncodeToString(salt)
	encodedHash := base64.RawStdEncoding.EncodeToString(hash)

	format := "$argon2id$v=19$m=%d,t=%d,p=%d$%s$%s"
	return fmt.Sprintf(format, params.Memory, params.Iterations, params.Parallelism, encodedSalt, encodedHash), nil
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	var input RegisterUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	params := DefaultHashParams()
	encodedHash, err := HashPassword(input.Password, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	newUser := model.User{
		ID: uuid.New().String(),
		Email: input.Email,
		Password: encodedHash,
		PublicKey: input.PublicKey,
	}

	query := `INSERT INTO users (id, email, password_hash, public_key) VALUES ($1, $2, $3, $4) RETURNING id, created_at`
	err = h.DB.QueryRow(context.Background(), query, newUser.ID, newUser.Email, newUser.Password, newUser.PublicKey).Scan(&newUser.ID, &newUser.CreatedAt)
	if err != nil {
		// TODO: Proper error handling should check for unique constraint violation etc.
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newUser)
}

type LoginUserInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// VerifyPassword verifies if the provided password matches the stored hash
func VerifyPassword(password, encodedHash string) (bool, error) {
	// Parse the encoded hash
	params, salt, hash, err := parseEncodedHash(encodedHash)
	if err != nil {
		return false, fmt.Errorf("failed to parse hash: %w", err)
	}

	// Generate hash with the same parameters and salt
	testHash := argon2.IDKey(
		[]byte(password),
		salt,
		params.Iterations,
		params.Memory,
		params.Parallelism,
		params.KeyLength,
	)

	// Use constant-time comparison to prevent timing attacks
	return subtle.ConstantTimeCompare(hash, testHash) == 1, nil
}

// parseEncodedHash parses the encoded hash string and extracts parameters, salt, and hash
func parseEncodedHash(encodedHash string) (*HashParams, []byte, []byte, error) {
	parts := strings.Split(encodedHash, "$")
	if len(parts) != 6 {
		return nil, nil, nil, fmt.Errorf("invalid hash format")
	}

	if parts[1] != "argon2id" {
		return nil, nil, nil, fmt.Errorf("unsupported algorithm: %s", parts[1])
	}

	if parts[2] != "v=19" {
		return nil, nil, nil, fmt.Errorf("unsupported version: %s", parts[2])
	}

	// Parse parameters
	var memory, iterations uint32
	var parallelism uint8
	_, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &memory, &iterations, &parallelism)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to parse parameters: %w", err)
	}

	// Decode salt
	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to decode salt: %w", err)
	}

	// Decode hash
	hash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to decode hash: %w", err)
	}

	params := &HashParams{
		Memory:      memory,
		Iterations:  iterations,
		Parallelism: parallelism,
		SaltLength:  uint32(len(salt)),
		KeyLength:   uint32(len(hash)),
	}

	return params, salt, hash, nil
}


// LoginUser handles user authentication and JWT generation.
func (h *UserHandler) LoginUser(c *gin.Context) {
	var input LoginUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// --- DEBUGGING ---
	fmt.Printf("Attempting login for email: %s\n", input.Email)

	// 1. Fetch user from the database
	var user model.User
	query := `SELECT id, email, password_hash, public_key, created_at FROM users WHERE email = $1`
	err := h.DB.QueryRow(context.Background(), query, input.Email).Scan(&user.ID, &user.Email, &user.Password, &user.PublicKey, &user.CreatedAt)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

		// --- DEBUGGING ---
	fmt.Printf("Found user: %s. Verifying password...\n", user.Email)

	// 2. Verify the password
	match, err := VerifyPassword(input.Password, user.Password)
	if err != nil || !match {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// 3. Generate a JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID, // Subject (the user's ID)
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(), // Expires in 30 days
	})

	jwtSecret := os.Getenv("JWT_SECRET_KEY")
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// 4. Send the token back to the client
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}