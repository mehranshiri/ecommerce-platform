package models

import "time"

type Payment struct {
	ID            uint      `gorm:"primaryKey;autoIncrement"`
	OrderID       uint      `gorm:"not null"`
	PaymentMethod string    `gorm:"not null"`
	Amount        float64   `gorm:"type:decimal(10,2);not null"`
	Currency      string    `gorm:"default:'USD'"`
	Status        string    `gorm:"default:'pending'"`
	TransactionID string    `gorm:"size:100"`
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
