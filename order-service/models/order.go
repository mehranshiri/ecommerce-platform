package models

import (
	"time"

	"gorm.io/gorm"
)

// User represents a user in the system
type Order struct {
	gorm.Model
	OrderNumber string    `gorm:"unique;not null"`
	UserID      uint      `gorm:"not null"`
	Status      string    `gorm:"default:'pending'"`
	TotalAmount float64   `gorm:"type:decimal(10,2);not null"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	CancelledAt *time.Time
	DeliveredAt *time.Time

	// Relationships
	User       User        `gorm:"foreignKey:UserID"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderID"`
	Payment    []Payment   `gorm:"foreignKey:OrderID"`
}
