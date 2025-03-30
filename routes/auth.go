package routes

import (
	"go-gin-boilerplate/controllers"

	"github.com/gin-gonic/gin"
)

// AuthRoutes defines the routes for authentication
func AuthRoutes(router *gin.Engine) {
	AuthGroup := router.Group("/auth")
	{
		// Register a new user
		AuthGroup.POST("/register", controllers.Register)
		// Login user and issue token
		AuthGroup.POST("/login", controllers.Login)
		// Logout user and invalidate token
		AuthGroup.POST("/logout/:username", controllers.Logout)
	}
}
