package routes

import (
	"go-gin-boilerplate/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// InitRoutes initializes all routes
func InitRoutes(r *gin.Engine) {
	println("InitRoutes started")

	// Root route
	r.GET("/", func(c *gin.Context) { controllers.DefaultPage(c) })

	// Health-check route
	r.GET("/health-check", func(c *gin.Context) { controllers.HealthCheck(c) })

	r.NoRoute(func(c *gin.Context) {
		println("Fallback route triggered for:", c.Request.URL.Path)
		c.JSON(http.StatusNotFound, gin.H{"message": "Route not found"})
	})

	println("InitRoutes completed")
}
