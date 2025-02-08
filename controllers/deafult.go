package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DefaultPage handler
func DefaultPage(c *gin.Context) {
	println("DefaultPage invoked")
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the Home Page",
	})
}

// HealthCheck handler
func HealthCheck(c *gin.Context) {
	println("HealthCheck invoked")
	c.JSON(http.StatusOK, gin.H{
		"message": "Server is up and running",
	})
}
