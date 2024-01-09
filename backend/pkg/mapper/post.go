package mapper

import (
	"bookmark-backend/internal/domain/request"
	"bookmark-backend/internal/domain/response"
	"bookmark-backend/internal/models"
)

type PostMapper struct{}

func NewPostMapper() *PostMapper {
	return &PostMapper{}
}

func (pm *PostMapper) MapToPostModel(req request.CreatePostRequest, userID uint, categoryID uint) models.Post {
	return models.Post{
		Title:      req.Title,
		Content:    req.Content,
		UserID:     userID,
		CategoryID: categoryID,
	}
}

func (pm *PostMapper) MapToUpdatePostModel(req request.UpdatePostRequest, userID uint, categoryID uint) models.Post {
	return models.Post{
		ID:         uint(req.PostID),
		Title:      req.Title,
		Content:    req.Content,
		UserID:     userID,
		CategoryID: categoryID,
	}
}

func (pm *PostMapper) MapToPostResponse(post *models.Post) response.PostResponse {
	return response.PostResponse{
		ID:      post.ID,
		Title:   post.Title,
		Content: post.Content,
		User: response.UserResponse{
			ID:       post.User.ID,
			Username: post.User.Username,
			Email:    post.User.Email,
		},
	}
}

func (pm *PostMapper) MapToPostsResponse(posts *[]models.Post) []response.PostsResponse {
	var postResponses []response.PostsResponse

	for _, post := range *posts {
		postResponse := response.PostsResponse{
			ID:      post.ID,
			Title:   post.Title,
			Content: post.Content,
			User: response.UserResponse{
				ID:       post.UserID,
				Username: post.User.Username,
				Email:    post.User.Email,
			},
		}
		postResponses = append(postResponses, postResponse)
	}

	return postResponses
}
