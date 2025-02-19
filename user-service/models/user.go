package models

import (
	"gorm.io/gorm"
)

// User represents a user in the system
type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}
