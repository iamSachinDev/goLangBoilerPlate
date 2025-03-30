package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Register handler
func Register(c *gin.Context) {
	var user User
	println("Register", c.ShouldBindJSON(&user))
	c.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
	})
}

// Login handler
func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "User Login successfully",
	})
}

// Logout handler
func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "User Logout successfully",
	})
}
