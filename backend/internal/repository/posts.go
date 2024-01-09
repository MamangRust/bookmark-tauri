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

func (r *postsRepository) FindAllPosts() (*[]models.Post, error) {
	var posts []models.Post

	db := r.db.Model(&posts)

	checkPosts := db.Debug().Preload("User").Find(&posts)

	if checkPosts.RowsAffected < 1 {
		return nil, gorm.ErrRecordNotFound
	}

	return &posts, nil
}

func (r *postsRepository) FindPostByID(postID int) (*models.Post, error) {
	var post models.Post

	checkPost := r.db.Debug().Preload("User").Where("id = ?", postID).Find(&post)

	if checkPost.RowsAffected < 1 {
		return &post, gorm.ErrRecordNotFound
	}

	return &post, nil
}

func (r *postsRepository) FindPostByTitle(title string) (*models.Post, error) {
	var post models.Post

	db := r.db.Model(&post)

	checkPost := db.Debug().Where("title = ?", title).Find(&post)

	if checkPost.RowsAffected < 1 {
		return &post, gorm.ErrRecordNotFound
	}

	return &post, nil
}

func (r *postsRepository) CreatePost(request models.Post) (*models.Post, error) {
	var newPost models.Post

	db := r.db.Model(&newPost)

	newPost.Title = request.Title
	newPost.Content = request.Content
	newPost.UserID = uint(request.UserID)
	newPost.CategoryID = uint(request.CategoryID)

	checkPost := db.Debug().Where("title = ?", newPost.Title).Find(&newPost)

	if checkPost.RowsAffected > 0 {
		return nil, gorm.ErrDuplicatedKey
	}

	addPost := db.Debug().Create(&newPost).Commit()

	if addPost.RowsAffected < 1 {
		return nil, gorm.ErrRecordNotFound
	}

	return &newPost, nil

}

func (r *postsRepository) UpdatePost(request models.Post) (*models.Post, error) {
	var post models.Post

	db := r.db.Model(&post)

	checkPost := db.Debug().Where("id = ?", request.ID).First(&post)

	if checkPost.RowsAffected < 1 {
		return nil, gorm.ErrRecordNotFound
	}

	post.Title = request.Title
	post.Content = request.Content
	post.UserID = uint(request.UserID)
	post.CategoryID = uint(request.CategoryID)

	updatePost := db.Debug().Updates(&post)

	if updatePost.RowsAffected < 1 {
		return nil, gorm.ErrRecordNotFound
	}

	return &post, nil
}

func (r *postsRepository) DeletePost(id int) error {
	var post models.Post

	db := r.db.Model(&post)

	checkPost := db.Debug().Where("id = ?", id).First(&post)

	if checkPost.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	deletePost := db.Debug().Delete(&post)

	if deletePost.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}
	return nil

}
