package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model

	PostID     uint   `gorm:"primaryKey"`
	Title      string `gorm:"not null"`
	Content    string
	UserID     uint     `gorm:"not null"`
	User       User     `gorm:"foreignKey:UserID"`
	CategoryID uint     `gorm:"not null"`
	Category   Category `gorm:"foreignKey:CategoryID"`
}
