package main

import (
	"user-service/database"
	"user-service/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	database.ConnectDB()

	router.POST("/register", handlers.RegisterUser)
	router.POST("/login", handlers.LoginUser)
	router.GET("/", func(ctx *gin.Context) {
		data := gin.H{
			"key": 23324,
		}
		ctx.JSON(200, gin.H{
			"message": "Hello, Gin!",
			"data":    data,
		})
	})

	// Start the Gin server
	router.Run(":8080")
}
