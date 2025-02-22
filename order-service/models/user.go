package models

import (
	"gorm.io/gorm"
)

// User represents a user in the system
type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey;autoIncrement"` // Explicitly define the ID field
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`

	Orders []Order `gorm:"foreignKey:UserID"`
}
