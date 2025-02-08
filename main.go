package main

import (
	"log"
	"os"

	"go-gin-boilerplate/server"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Load environment variables
func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using default values")
	}
}

func main() {
	// Load environment variables
	loadEnv()

	// Set Gin mode from .env
	mode := os.Getenv("GIN_MODE")
	println("mode: ", mode)
	if mode == "" {
		mode = gin.DebugMode // Default to debug mode if not set
	}
	gin.SetMode(mode)

	log.Printf("Running in %s mode", mode)

	// Start the server
	server.StartServer()
}
