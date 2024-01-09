package mapper

import (
	"bookmark-backend/internal/domain/request"
	"bookmark-backend/internal/domain/response"
	"bookmark-backend/internal/models"
)

type CategoryMapper struct{}

func NewCategoryMapper() *CategoryMapper {
	return &CategoryMapper{}
}

func (cm *CategoryMapper) MapToCategoryModel(req request.CreateCategoryRequest) models.Category {
	return models.Category{
		Name:        req.Name,
		Image:       req.Image,
		Description: req.Description,
	}
}

func (cm *CategoryMapper) MapToUpdateCategoryModel(req request.UpdateCategoryRequest) models.Category {

	return models.Category{
		ID:          uint(req.CategoryID),
		Name:        req.Name,
		Image:       req.Image,
		Description: req.Description,
	}
}

func (cm *CategoryMapper) MapToCategoryResponse(category *models.Category) response.CategoriesResponse {
	var postResponses []response.PostResponse

	for _, post := range category.Posts {
		postResponse := response.PostResponse{
			ID:      post.ID,
			Title:   post.Title,
			Content: post.Content,
		}
		postResponses = append(postResponses, postResponse)
	}

	return response.CategoriesResponse{
		ID:          category.ID,
		Name:        category.Name,
		Image:       category.Image,
		Description: category.Description,
		Posts:       postResponses,
	}
}

func (cm *CategoryMapper) MapToCategoriesResponse(categories *[]models.Category) []response.CategoriesResponse {
	var categoryResponses []response.CategoriesResponse

	for _, category := range *categories {
		var postResponses []response.PostResponse

		for _, post := range category.Posts {
			postResponse := response.PostResponse{
				ID:      post.ID,
				Title:   post.Title,
				Content: post.Content,
			}

			postResponses = append(postResponses, postResponse)

		}

		categoryResponse := response.CategoriesResponse{
			ID:          category.ID,
			Name:        category.Name,
			Image:       category.Image,
			Description: category.Description,
			Posts:       postResponses,
		}
		categoryResponses = append(categoryResponses, categoryResponse)
	}

	return categoryResponses
}
