package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting NQI API server...")

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	err := router.Run(":8081")
	if err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
