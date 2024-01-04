package service

import (
	"bookmark-backend/internal/domain/request"
	"bookmark-backend/internal/domain/response"
	"bookmark-backend/internal/repository"
	"bookmark-backend/pkg/logger"
	"fmt"

	"go.uber.org/zap"
)

type postService struct {
	repository         repository.PostsRepository
	categoryRepository repository.CategoryRepository
	userRepository     repository.UserRepository
	logger             logger.Logger
}

func NewPostService(repository repository.PostsRepository, categoryRepository repository.CategoryRepository, userRepository repository.UserRepository, logger logger.Logger) *postService {
	return &postService{repository: repository, categoryRepository: categoryRepository, userRepository: userRepository, logger: logger}
}

func (s *postService) FindAllPosts() (*response.ServiceResponse, *response.ServiceError) {
	res, err := s.repository.FindAllPosts()

	if err != nil {
		s.logger.Error("Error fetching all posts", zap.Error(err))
		return nil, &response.ServiceError{Err: err, Description: "Error fetching all posts"}
	}

	return &response.ServiceResponse{Data: res}, nil
}

func (s *postService) FindPostByID(postID int) (*response.ServiceResponse, *response.ServiceError) {
	res, err := s.repository.FindPostByID(postID)

	if err != nil {
		s.logger.Error("Error fetching post by ID", zap.Error(err), zap.Int("PostID", postID))

		return nil, &response.ServiceError{Err: err, Description: fmt.Sprintf("Error fetching post by ID: %d", postID)}

	}

	return &response.ServiceResponse{Data: res}, nil
}

func (s *postService) FindPostByTitle(title string) (*response.ServiceResponse, *response.ServiceError) {
	res, err := s.repository.FindPostByTitle(title)

	if err != nil {
		s.logger.Error("Error fetching post by title", zap.Error(err), zap.String("Title", title))

		return nil, &response.ServiceError{Err: err, Description: fmt.Sprintf("Error fetching post by title: %s", title)}
	}

	return &response.ServiceResponse{Data: res}, nil
}

func (s *postService) Create(requests request.CreatePostRequest) (*response.ServiceResponse, *response.ServiceError) {
	var requestCreate request.CreatePostRequest

	user, err := s.userRepository.FindUserByID(requests.UserID)
	if err != nil {
		s.logger.Error("Error fetching user", zap.Error(err), zap.Int("UserID", requests.UserID))

		return nil, &response.ServiceError{Err: err, Description: fmt.Sprintf("Error fetching user: %d", requests.UserID)}
	}

	category, err := s.categoryRepository.FindCategoryByID(requests.CategoryID)
	if err != nil {
		s.logger.Error("Error fetching category", zap.Error(err), zap.Int("CategoryID", requests.CategoryID))

		return nil, &response.ServiceError{Err: err, Description: fmt.Sprintf("Error fetching category: %d", requests.CategoryID)}

	}

	requestCreate.Title = requests.Title
	requestCreate.Content = requests.Content
	requestCreate.UserID = int(user.ID)
	requestCreate.CategoryID = int(category.ID)

	res, err := s.repository.CreatePost(requestCreate)
	if err != nil {
		s.logger.Error("Error creating post", zap.Error(err))
		return nil, &response.ServiceError{Err: err, Description: "Error creating post"}
	}

	return &response.ServiceResponse{Data: res}, nil
}

func (s *postService) Update(requests request.UpdatePostRequest) (*response.ServiceResponse, *response.ServiceError) {
	var updateRequest request.UpdatePostRequest

	user, err := s.userRepository.FindUserByID(requests.UserID)
	if err != nil {
		s.logger.Error("Error fetching user", zap.Error(err), zap.Int("UserID", requests.UserID))

		return nil, &response.ServiceError{Err: err, Description: fmt.Sprintf("Error fetching user: %d", requests.UserID)}
	}

	category, err := s.categoryRepository.FindCategoryByID(requests.CategoryID)
	if err != nil {
		s.logger.Error("Error fetching category", zap.Error(err), zap.Int("CategoryID", requests.CategoryID))

		return nil, &response.ServiceError{Err: err, Description: fmt.Sprintf("Error fetching category: %d", requests.CategoryID)}

	}

	updateRequest.Title = requests.Title
	updateRequest.Content = requests.Content
	updateRequest.UserID = int(user.ID)
	updateRequest.CategoryID = int(category.ID)

	res, err := s.repository.UpdatePost(updateRequest)
	if err != nil {
		s.logger.Error("Error updating post", zap.Error(err), zap.Int("PostID", requests.PostID))

		return nil, &response.ServiceError{Err: err, Description: fmt.Sprintf("Error updating post: %d", requests.PostID)}
	}

	return &response.ServiceResponse{Data: res}, nil
}

func (s *postService) Delete(id int) (*response.ServiceResponse, *response.ServiceError) {
	err := s.repository.DeletePost(id)

	if err != nil {
		s.logger.Error("Error deleting post", zap.Error(err), zap.Int("PostID", id))
		return nil, &response.ServiceError{Err: err, Description: fmt.Sprintf("Error deleting post: %d", id)}
	}

	return &response.ServiceResponse{Data: "Post deleted"}, nil
}
