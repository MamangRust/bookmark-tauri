package service

import (
	"bookmark-backend/internal/domain/request"
	"bookmark-backend/internal/domain/response"
	"mime/multipart"
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
	Create(userId int, requests request.CreatePostRequest) (*response.ServiceResponse, *response.ServiceError)
	Update(userId int, requests request.UpdatePostRequest) (*response.ServiceResponse, *response.ServiceError)
	Delete(id int) (*response.ServiceResponse, *response.ServiceError)
}

type FolderService interface {
	CreateFolder(name string) (string, error)

	CheckAndUpdateFolder(oldFolder string, newFolder string) error
	DeleteFolder(folder string) error
}

type FileService interface {
	CreateFileImage(file multipart.File, filePath string) (string, error)
	CreateFile(request request.CreateFileRequest) (*response.ServiceResponse, *response.ServiceError)
	FindFile(request request.FileRequest) (*response.ServiceResponse, *response.ServiceError)
	UpdateFile(request request.UpdateFileRequest) (*response.ServiceResponse, *response.ServiceError)

	DeleteFile(filePath string) error
}
