package service

import (
	"bookmark-backend/internal/domain/request"
	"bookmark-backend/internal/domain/response"
)

type AuthService interface {
	Login(request request.LoginUserRequest) (*request.Token, *response.ServiceError)
	Register(request request.RegisterUserRequest) (*response.ServiceResponse, *response.ServiceError)
}

type UserService interface {
	GetUserAll() (*response.ServiceResponse, *response.ServiceError)
	GetUserByID(id int) (*response.ServiceResponse, *response.ServiceError)
	CreateUser(request request.CreateUserRequest) (*response.ServiceResponse, *response.ServiceError)
	UpdateUser(id int, request request.UpdateUserRequest) (*response.ServiceResponse, *response.ServiceError)
	DeleteUser(id int) (*response.ServiceResponse, *response.ServiceError)
}

type CategoryService interface {
	GetAll() (*response.ServiceResponse, *response.ServiceError)
	GetByID(id int) (*response.ServiceResponse, *response.ServiceError)
	Create(request request.CreateCategoryRequest) (*response.ServiceResponse, *response.ServiceError)
	Update(request request.UpdateCategoryRequest) (*response.ServiceResponse, *response.ServiceError)
	Delete(id int) (*response.ServiceResponse, *response.ServiceError)
}

type PostService interface {
	FindAllPosts() (*response.ServiceResponse, *response.ServiceError)
	FindPostByID(postID int) (*response.ServiceResponse, *response.ServiceError)
	FindPostByTitle(title string) (*response.ServiceResponse, *response.ServiceError)
	Create(requests request.CreatePostRequest) (*response.ServiceResponse, *response.ServiceError)
	Update(requests request.UpdatePostRequest) (*response.ServiceResponse, *response.ServiceError)
	Delete(id int) (*response.ServiceResponse, *response.ServiceError)
}
