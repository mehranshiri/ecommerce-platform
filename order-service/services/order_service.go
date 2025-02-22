package services

import (
	"errors"
	"order-service/config"
	"order-service/database"
	"order-service/models"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// RegisterUserService handles business logic for user registration
func RegisterUserService(input models.User) (models.User, string, error) {
	input.Email = strings.TrimSpace(input.Email)
	input.Password = strings.TrimSpace(input.Password)

	var existingUser models.User
	if err := database.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		return models.User{}, "", errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, "", errors.New("failed to hash password")
	}
	input.Password = string(hashedPassword)

	if err := database.DB.Create(&input).Error; err != nil {
		return models.User{}, "", errors.New("failed to save user")
	}

	token, err := config.GenerateJWT(input.ID, input.Email)
	if err != nil {
		return models.User{}, "", errors.New("failed to generate token")
	}

	return input, token, nil
}
