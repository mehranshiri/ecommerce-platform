package main

import "github.com/gin-gonic/gin"

func registerUser(c *gin.Context) {

}

func main() {
	router := gin.Default()

	router.POST("/register", registerUser)

	router.Run("localhost:8080")

}
