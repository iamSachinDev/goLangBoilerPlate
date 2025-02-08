package routes

import (
	"go-gin-boilerplate/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// InitRoutes initializes all routes
func InitRoutes(route *gin.Engine) {

	// Group for API Home and Health Check
	DefaultApis := route.Group("default")
	{
		// Home Page
		DefaultApis.GET("/", func(c *gin.Context) { controllers.DefaultPage(c) })
		// Health-check route
		DefaultApis.GET("/health-check", func(c *gin.Context) { controllers.HealthCheck(c) })
	}

	route.NoRoute(func(c *gin.Context) {
		println("Fallback route triggered for:", c.Request.URL.Path)
		c.JSON(http.StatusNotFound, gin.H{"message": "Route not found"})
	})

	println("InitRoutes completed")
}
