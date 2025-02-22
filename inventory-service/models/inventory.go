package models

import (
	"gorm.io/gorm"
)

// User represents a user in the system
type Inventory struct {
	gorm.Model
	SerialNo int16  `gorm:"unique;not null"`
	Title    string `gorm:"unique;not null"`
	Quantity int    `gorm:"not null"`
	Price    int    `gorm:"not null"`
}

func Create(db *gorm.DB, input *Inventory) error {
	return db.Create(input).Error
}
