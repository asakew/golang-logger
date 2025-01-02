package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// Create a new Gin router
	r := gin.Default()

	// Define a simple route
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Gin! 🌟")
	})

	log.Println("http://localhost:8080")

	// Start the server
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
