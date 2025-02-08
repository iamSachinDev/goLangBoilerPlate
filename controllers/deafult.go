package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DefaultPage handler
func DefaultPage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the Home Page",
	})
}

// HealthCheck handler
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Server is up and running",
	})
}
