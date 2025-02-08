package server

import (
	"log"
	"os"

	"go-gin-boilerplate/middlewares"
	"go-gin-boilerplate/routes"

	"github.com/gin-gonic/gin"
)

// StartServer initializes the Gin server
func StartServer() {
	r := gin.Default()

	// Register security middleware
	r.Use(middlewares.SecurityMiddleware())

	// Initialize routes
	routes.InitRoutes(r)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting server on port %s...", port)
	r.Run(":" + port)
}
