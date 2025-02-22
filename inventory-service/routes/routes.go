package routes

import (
	"inventory-service/handlers"
	"inventory-service/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) *gin.Engine {

	// Register public and protected routes
	setupPublicRoutes(router)
	setupProtectedRoutes(router)

	return router
}

// setupPublicRoutes configures routes that don't require authentication
func setupPublicRoutes(router *gin.Engine) {
	// router.POST("/login", handlers.LoginUser)
	// router.POST("/inventory", handlers.CreateInventory)
}

// setupProtectedRoutes configures routes that require authentication
func setupProtectedRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/inventory", handlers.CreateInventory)
	}
}
