package routes

import (
	"github.com/gin-gonic/gin"
)

// SetupRoutes initializes all route groups
func InitRoutes(router *gin.Engine) {
	// Define the DefaultApis route group
	AuthRoutes(router)
	DefaultRoutes(router)
}
