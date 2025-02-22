package models

import "time"

type OrderAuditLog struct {
	ID        uint       `gorm:"primaryKey;autoIncrement"`
	OrderID   uint       `gorm:"not null"`
	Action    string     `gorm:"not null"`
	ChangedAt *time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
