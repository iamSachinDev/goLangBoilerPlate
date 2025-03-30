package routes

import (
	"go-gin-boilerplate/controllers"

	"github.com/gin-gonic/gin"
)

// defaultRoutes initializes routes under the "default" group
func DefaultRoutes(router *gin.Engine) {
	println("defaultRoutes")
	DefaultGroup := router.Group("default")
	{
		// Home Page
		DefaultGroup.GET("/", controllers.DefaultPage)
		// Health-check route
		DefaultGroup.GET("/health-check", controllers.HealthCheck)
	}
}
