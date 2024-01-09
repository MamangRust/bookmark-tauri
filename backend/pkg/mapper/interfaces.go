package mapper

import (
	"bookmark-backend/internal/domain/request"
	"bookmark-backend/internal/domain/response"
	"bookmark-backend/internal/models"
)

type PostMapping interface {
	MapToPostModel(req request.CreatePostRequest, userID uint, categoryID uint) models.Post
	MapToUpdatePostModel(req request.UpdatePostRequest, userID uint, categoryID uint) models.Post
	MapToPostResponse(post *models.Post) response.PostResponse
	MapToPostsResponse(posts *[]models.Post) []response.PostsResponse
}

type CategoryMapping interface {
	MapToCategoryModel(req request.CreateCategoryRequest) models.Category
	MapToUpdateCategoryModel(req request.UpdateCategoryRequest) models.Category
	MapToCategoryResponse(category *models.Category) response.CategoriesResponse
	MapToCategoriesResponse(categories *[]models.Category) []response.CategoriesResponse
}
