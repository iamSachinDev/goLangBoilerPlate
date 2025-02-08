package routes

import (
	"go-gin-boilerplate/controllers"

	"github.com/gin-gonic/gin"
)

// InitRoutes initializes all routes
// InitRoutes initializes the routes for the given gin engine.
// It sets up the route for the home controller.
//
// Parameters:
//   - r: A pointer to the gin.Engine instance where the routes will be registered.
func InitRoutes(r *gin.Engine) {
	println("InitRoutes")
	println("InitRoutes", r)
	r.GET("/", controllers.HomeController)
}
