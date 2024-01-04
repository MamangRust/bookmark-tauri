package models

type Post struct {
	ID         uint     `gorm:"primaryKey;unique"`
	Title      string   `gorm:"not null"`
	Content    string   `gorm:"not null"`
	CategoryID uint     `gorm:"not null"`
	Category   Category `gorm:"foreignKey:CategoryID"`
	UserID     uint     `gorm:"not null"`
	User       User     `gorm:"foreignKey:UserID"`
}
