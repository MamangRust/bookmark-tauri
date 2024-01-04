package models

type User struct {
	ID       uint   `gorm:"primaryKey;unique"`
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Posts    []Post `gorm:"foreignKey:UserID"`
}
