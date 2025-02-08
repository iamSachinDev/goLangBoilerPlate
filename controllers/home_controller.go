package controllers

import "github.com/gin-gonic/gin"

// HomeController handles the home route
func HomeController(c *gin.Context) {
    c.JSON(200, gin.H{"message": "Welcome to Go Boilerplate!"})
}
