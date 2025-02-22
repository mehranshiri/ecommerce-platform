package routes

import (
	"order-service/handlers"
	"order-service/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter initializes routes
func SetupRouter(router *gin.Engine) *gin.Engine {

	// Register public and protected routes
	setupPublicRoutes(router)
	setupProtectedRoutes(router)

	return router
}

// setupPublicRoutes configures routes that don't require authentication
func setupPublicRoutes(router *gin.Engine) {
	router.POST("/register", handlers.RegisterUser)
	router.POST("/login", handlers.LoginUser)
}

// setupProtectedRoutes configures routes that require authentication
func setupProtectedRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/users", handlers.GetUsers)
		auth.POST("/user", handlers.GetUser)
	}
}
