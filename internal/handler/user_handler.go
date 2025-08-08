package handler

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/rghsoftware/nqi/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/argon2"
)

// Argon2Params holds the parameters for the Argon2id hashing algorithm.
// These are recommended baseline settings.
type Argon2Params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

type UserHandler struct {
	DB *pgxpool.Pool
}

type RegisterUserInput struct {
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	PublicKey string `json:"publicKey" binding:"required"`
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	var input RegisterUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	params := Argon2Params{
		memory: 64 * 1024, // 64 MB
		iterations: 1,
		parallelism: 4,
		saltLength: 16,
		keyLength: 32,
	}

	salt := make([]byte, params.saltLength)
	if _, err := rand.Read(salt); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate salt"})
		return
	}

	hash := argon2.IDKey([]byte(input.Password), salt, params.iterations, params.memory, params.parallelism, params.keyLength)
	b64Salt := base64.StdEncoding.EncodeToString(salt)
	b64Hash := base64.StdEncoding.EncodeToString(hash)
	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, params.memory, params.iterations, params.parallelism, b64Salt, b64Hash)

	newUser := model.User{
		ID: uuid.New().String(),
		Email: input.Email,
		Password: encodedHash,
		PublicKey: input.PublicKey,
	}

	query := `INSERT INTO users (id, email, password_hash, public_key) VALUES ($1, $2, $3, $4) RETURNING id, created_at`
	err := h.DB.QueryRow(context.Background(), query, newUser.ID, newUser.Email, newUser.Password, newUser.PublicKey).Scan(&newUser.ID, &newUser.CreatedAt)
	if err != nil {
		// TODO: Proper error handling should check for unique constraint violation etc.
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newUser)
}
