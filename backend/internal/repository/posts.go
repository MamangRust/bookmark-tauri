package repository

import (
	"bookmark-backend/internal/models"

	"gorm.io/gorm"
)

type postsRepository struct {
	db *gorm.DB
}

func NewPostsRepository(db *gorm.DB) *postsRepository {
	return &postsRepository{db: db}
}

func (r *postsRepository) FindAllPosts() ([]models.Post, error) {
	var posts []models.Post
	if err := r.db.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *postsRepository) FindPostByID(postID uint) (models.Post, error) {
	var post models.Post
	if err := r.db.Where("post_id = ?", postID).First(&post).Error; err != nil {
		return post, err
	}
	return post, nil
}

func (r *postsRepository) FindPostByTitle(title string) (models.Post, error) {
	var post models.Post
	if err := r.db.Where("title = ?", title).First(&post).Error; err != nil {
		return post, err
	}
	return post, nil
}

func (r *postsRepository) CreatePost(post models.Post) error {
	if err := r.db.Create(&post).Error; err != nil {
		return err
	}
	return nil
}

func (r *postsRepository) UpdatePost(post models.Post) error {
	if err := r.db.Model(&post).Updates(&post).Error; err != nil {
		return err
	}
	return nil
}

func (r *postsRepository) DeletePost(post models.Post) error {
	if err := r.db.Delete(&post).Error; err != nil {
		return err
	}
	return nil
}
