package response

type UserResponse struct {
	ID       uint   `gorm:"primaryKey;unique"`
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
}
