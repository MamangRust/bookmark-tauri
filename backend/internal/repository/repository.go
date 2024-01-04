package repository

import "gorm.io/gorm"

type Repositories struct {
	Category CategoryRepository
	Post     PostsRepository
	User     UserRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Category: NewCategoryRepository(db),
		Post:     NewPostsRepository(db),
		User:     NewUserRepository(db),
	}
}
