package main

import (
	"database/sql"
	"fmt"
	"log"
	"user-service/config"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

func registerUser(c *gin.Context) {
	// Handler for user registration
}

func main() {
	config.LoadConfig()

	dbHost := config.GetEnv("DB_HOST", "db")
	dbPort := config.GetEnv("DB_PORT", "5432")
	dbUser := config.GetEnv("DB_USER", "postgres")
	dbPassword := config.GetEnv("DB_PASSWORD", "password")
	dbName := config.GetEnv("DB_NAME", "ecommerce_db")

	// PostgreSQL Connection String
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	// Open connection to PostgreSQL
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	fmt.Println("✅ Successfully connected to PostgreSQL!")

	router := gin.Default()

	router.POST("/register", registerUser)
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
