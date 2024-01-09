package repository

import (
	"bookmark-backend/internal/domain/request"
	"bookmark-backend/internal/models"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindAllUsers() (*[]models.User, error) {
	var users []models.User

	db := r.db.Model(&users)

	checkUsers := db.Debug().Find(&users)

	if checkUsers.RowsAffected < 1 {
		return nil, gorm.ErrRecordNotFound
	}

	return &users, nil
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User

	db := r.db.Model(&user)

	checkUser := db.Debug().Where("email = ?", email).First(&user)

	if checkUser.RowsAffected < 1 {
		return &user, gorm.ErrRecordNotFound
	}

	return &user, nil
}

func (r *userRepository) FindUserByID(userID int) (*models.User, error) {
	var user models.User

	db := r.db.Model(&user)

	checkUser := db.Debug().Where("id = ?", userID).First(&user)

	if checkUser.RowsAffected < 1 {
		return &user, gorm.ErrRecordNotFound
	}

	return &user, nil

}

func (r *userRepository) CreateUser(request request.CreateUserRequest) (*models.User, error) {
	var user models.User

	db := r.db.Model(&user)

	checkUser := db.Debug().Where("email = ?", request.Email).First(&user)

	if checkUser.RowsAffected > 0 {
		return nil, gorm.ErrDuplicatedKey
	}

	user.Username = request.Username
	user.Email = request.Email
	user.Password = request.Password

	addUser := db.Debug().Create(&user).Commit()

	if addUser.RowsAffected < 1 {
		return nil, gorm.ErrRecordNotFound
	}

	return &user, nil

}

func (r *userRepository) UpdateUser(id int, request request.UpdateUserRequest) (*models.User, error) {
	var user models.User

	db := r.db.Model(&user)

	checkUser := db.Debug().Where("id = ?", id).First(&user)

	if checkUser.RowsAffected < 1 {
		return nil, gorm.ErrRecordNotFound
	}

	user.Username = request.Username
	user.Email = request.Email
	user.Password = request.Password

	updateUser := db.Debug().Updates(&user)

	if updateUser.RowsAffected < 1 {
		return nil, gorm.ErrRecordNotFound
	}

	return &user, nil

}

func (r *userRepository) DeleteUser(id int) error {
	var user models.User

	db := r.db.Model(&user)

	checkUser := db.Debug().Where("id = ?", id).First(&user)

	if checkUser.RowsAffected < 1 {

		return gorm.ErrRecordNotFound
	}

	deleteUser := db.Debug().Delete(&user)

	if deleteUser.RowsAffected < 1 {

		return gorm.ErrRecordNotFound
	}

	return nil

}
