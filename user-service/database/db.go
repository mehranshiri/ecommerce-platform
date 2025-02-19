package database

import (
	"fmt"
	"log"

	"user-service/config"
	"user-service/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	config.LoadConfig()

	// Load environment variables
	dbHost := config.GetEnv("DB_HOST", "db")
	dbPort := config.GetEnv("DB_PORT", "5432")
	dbUser := config.GetEnv("DB_USER", "postgres")
	dbPassword := config.GetEnv("DB_PASSWORD", "password")
	dbName := config.GetEnv("DB_NAME", "ecommerce_db")

	// PostgreSQL DSN
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	// Connect to database
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	log.Println("✅ Successfully connected to PostgreSQL!")

	// Auto-migrate the User model
	DB.AutoMigrate(&models.User{})
}
