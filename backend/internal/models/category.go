package models

type Category struct {
	ID          uint   `gorm:"primaryKey;unique"`
	Name        string `gorm:"unique;not null"`
	Image       string `gorm:"not null"`
	Description string `gorm:"not null"`
	Posts       []Post `gorm:"foreignKey:CategoryID"`
}
