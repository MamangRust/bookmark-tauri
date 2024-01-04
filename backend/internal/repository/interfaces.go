package repository

import (
	"bookmark-backend/internal/domain/request"
	"bookmark-backend/internal/models"
)

type PostsRepository interface {
	FindAllPosts() (*[]models.Post, error)
	FindPostByID(postID int) (*models.Post, error)
	FindPostByTitle(title string) (*models.Post, error)
	CreatePost(post request.CreatePostRequest) (*models.Post, error)
	UpdatePost(request request.UpdatePostRequest) (*models.Post, error)
	DeletePost(id int) error
}

type CategoryRepository interface {
	FindAllCategory() (*[]models.Category, error)
	FindCategoryByID(categoryID int) (*models.Category, error)
	FindCategoryByName(name string) (*models.Category, error)
	CreateCategory(request request.CreateCategoryRequest) (*models.Category, error)
	UpdateCategory(request request.UpdateCategoryRequest) (*models.Category, error)
	DeleteCategory(id int) error
}

type UserRepository interface {
	FindAllUsers() (*[]models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindUserByID(userID int) (*models.User, error)
	CreateUser(request request.CreateUserRequest) (*models.User, error)
	UpdateUser(id int, request request.UpdateUserRequest) (*models.User, error)
	DeleteUser(id int) error
}
