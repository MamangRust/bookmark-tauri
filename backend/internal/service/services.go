package service

import (
	"bookmark-backend/internal/repository"
	"bookmark-backend/pkg/auth"
	"bookmark-backend/pkg/hash"
	"bookmark-backend/pkg/logger"
)

type Service struct {
	Auth     AuthService
	Category CategoryService
	Post     PostService
	User     UserService
}

type Deps struct {
	Repository *repository.Repositories
	Logger     logger.Logger
	Hash       hash.Hashing
	Token      auth.TokenManager
}

func NewService(deps Deps) *Service {

	return &Service{
		Auth: NewAuthService(
			deps.Repository.User,
			deps.Hash,
			deps.Token,
			deps.Logger,
		),
		Category: NewCategoryService(deps.Repository.Category, deps.Logger),
		Post:     NewPostService(deps.Repository.Post, deps.Repository.Category, deps.Repository.User, deps.Logger),
		User:     NewUserService(deps.Repository.User, deps.Hash, deps.Logger),
	}
}
