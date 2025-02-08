package routes

import (
    "github.com/gin-gonic/gin"
    "go-gin-boilerplate/controllers"
)

// InitRoutes initializes all routes
func InitRoutes(r *gin.Engine) {
    r.GET("/", controllers.HomeController)
}
