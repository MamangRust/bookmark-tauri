package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	CategoryID  uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Image       string `gorm:"not null"`
	Description string
	CreatedAt   time.Time
	Posts       []Post `gorm:"foreignKey:CategoryID"`
}
