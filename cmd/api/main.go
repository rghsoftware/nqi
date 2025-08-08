package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rghsoftware/nqi/internal/db"
	"github.com/rghsoftware/nqi/internal/handler"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("No .env file found, using environment variables")
	}

	dbPool, err := db.NewConnection()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}
	defer dbPool.Close()

	fmt.Println("Starting NQI API server...")

	router := gin.Default()

	userHandler := &handler.UserHandler{DB: dbPool}

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	v1 := router.Group("/api/v1")
	{
		v1.POST("/users/register", userHandler.RegisterUser)
	}

	err = router.Run(":8081")
	if err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
