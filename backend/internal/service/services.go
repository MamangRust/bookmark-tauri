package service

import (
	"bookmark-backend/internal/repository"
	"bookmark-backend/pkg/auth"
	"bookmark-backend/pkg/hash"
	"bookmark-backend/pkg/logger"
	"bookmark-backend/pkg/mapper"
)

type Service struct {
	Auth     AuthService
	Category CategoryService
	Post     PostService
	User     UserService
	File     FileService
	Folder   FolderService
}

type Deps struct {
	Repository *repository.Repositories
	Logger     *logger.Logger
	Hash       *hash.Hashing
	Token      auth.TokenManager
	Mapper     *mapper.Mapper
}

func NewService(deps Deps) *Service {
	folder := NewFolderService()
	file := NewFileService()

	return &Service{
		Auth: NewAuthService(
			deps.Repository.User,
			*deps.Hash,
			deps.Token,
			*deps.Logger,
		),
		Category: NewCategoryService(deps.Mapper.CategoryMapper, *file, *folder, deps.Repository.Category, *deps.Logger),
		Post:     NewPostService(*file, deps.Mapper.PostMapper, deps.Repository.Post, deps.Repository.Category, deps.Repository.User, *deps.Logger),
		User:     NewUserService(deps.Repository.User, *deps.Hash, *deps.Logger),
		File:     file,
		Folder:   folder,
	}
}
