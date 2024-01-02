package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID    uint   `gorm:"primaryKey"`
	Username  string `gorm:"not null"`
	Email     string `gorm:"not null;unique"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	Posts     []Post `gorm:"foreignKey:UserID"`
}
