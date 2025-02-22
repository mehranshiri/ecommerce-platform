package main

import (
	"inventory-service/database"
	"inventory-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	database.ConnectDB()

	routes.SetupRouter(router)

	// Start the Gin server
	router.Run(":8081")
}
