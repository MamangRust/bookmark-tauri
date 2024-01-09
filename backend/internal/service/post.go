package service

import (
	"bookmark-backend/internal/domain/request"
	"bookmark-backend/internal/domain/response"
	"bookmark-backend/internal/repository"
	"bookmark-backend/pkg/logger"
	"bookmark-backend/pkg/mapper"
	"fmt"

	"go.uber.org/zap"
)

type postService struct {
	file               fileService
	postMapper         mapper.PostMapping
	repository         repository.PostsRepository
	categoryRepository repository.CategoryRepository
	userRepository     repository.UserRepository
	logger             logger.Logger
}

func NewPostService(file fileService, postMapper mapper.PostMapping, repository repository.PostsRepository, categoryRepository repository.CategoryRepository, userRepository repository.UserRepository, logger logger.Logger) *postService {
	return &postService{file: file, postMapper: postMapper, repository: repository, categoryRepository: categoryRepository, userRepository: userRepository, logger: logger}
}

func (s *postService) FindAllPosts() (*response.ServiceResponse, *response.ServiceError) {
	res, err := s.repository.FindAllPosts()

	if err != nil {
		s.logger.Error("Error fetching all posts", zap.Error(err))
		return nil, &response.ServiceError{Err: err, Description: "Error fetching all posts"}
	}

	postsResponse := s.postMapper.MapToPostsResponse(res)

	return &response.ServiceResponse{Data: postsResponse}, nil
}

func (s *postService) FindPostByID(postID int) (*response.ServiceResponse, *response.ServiceError) {
	res, err := s.repository.FindPostByID(postID)

	if err != nil {
		s.logger.Error("Error fetching post by ID", zap.Error(err), zap.Int("PostID", postID))

		return nil, &response.ServiceError{Err: err, Description: fmt.Sprintf("Error fetching post by ID: %d", postID)}

	}

	postResponse := s.postMapper.MapToPostResponse(res)

	return &response.ServiceResponse{Data: postResponse}, nil
}

func (s *postService) FindPostByTitle(title string) (*response.ServiceResponse, *response.ServiceError) {
	res, err := s.repository.FindPostByTitle(title)

	if err != nil {
		s.logger.Error("Error fetching post by title", zap.Error(err), zap.String("Title", title))

		return nil, &response.ServiceError{Err: err, Description: fmt.Sprintf("Error fetching post by title: %s", title)}
	}

	postResponse := s.postMapper.MapToPostResponse(res)

	return &response.ServiceResponse{Data: postResponse}, nil
}

func (s *postService) Create(userId int, requests request.CreatePostRequest) (*response.ServiceResponse, *response.ServiceError) {

	user, err := s.userRepository.FindUserByID(userId)
	if err != nil {
		s.logger.Error("Error fetching user", zap.Error(err), zap.Int("UserID", userId))

		return nil, &response.ServiceError{Err: err, Description: fmt.Sprintf("Error fetching user: %d", userId)}
	}

	category, err := s.categoryRepository.FindCategoryByID(requests.CategoryID)
	if err != nil {
		s.logger.Error("Error fetching category", zap.Error(err), zap.Int("CategoryID", requests.CategoryID))

		return nil, &response.ServiceError{Err: err, Description: fmt.Sprintf("Error fetching category: %d", requests.CategoryID)}

	}

	requestCreate := s.postMapper.MapToPostModel(requests, user.ID, category.ID)

	res, err := s.repository.CreatePost(requestCreate)
	if err != nil {
		s.logger.Error("Error creating post", zap.Error(err))
		return nil, &response.ServiceError{Err: err, Description: "Error creating post"}
	}

	_, err_file := s.file.CreateFile(request.CreateFileRequest{
		Folder:  category.Name,
		Title:   res.Title,
		Content: res.Content,
	})

	if err_file != nil {
		s.logger.Error("Error creating file", zap.Error(err_file))
		return nil, &response.ServiceError{Err: err, Description: "Error creating file"}
	}

	postResponse := s.postMapper.MapToPostResponse(res)

	return &response.ServiceResponse{Data: postResponse}, nil
}

func (s *postService) Update(userId int, requests request.UpdatePostRequest) (*response.ServiceResponse, *response.ServiceError) {

	user, err := s.userRepository.FindUserByID(userId)
	if err != nil {
		s.logger.Error("Error fetching user", zap.Error(err), zap.Int("UserID", userId))

		return nil, &response.ServiceError{Err: err, Description: fmt.Sprintf("Error fetching user: %d", userId)}
	}

	category, err := s.categoryRepository.FindCategoryByID(requests.CategoryID)
	if err != nil {
		s.logger.Error("Error fetching category", zap.Error(err), zap.Int("CategoryID", requests.CategoryID))

		return nil, &response.ServiceError{Err: err, Description: fmt.Sprintf("Error fetching category: %d", requests.CategoryID)}

	}

	post, err := s.repository.FindPostByID(requests.PostID)

	if err != nil {
		s.logger.Error("Error fetching post", zap.Error(err), zap.Int("PostID", requests.PostID))

		return nil, &response.ServiceError{Err: err, Description: fmt.Sprintf("Error fetching post: %d", requests.PostID)}
	}

	updateRequest := s.postMapper.MapToUpdatePostModel(requests, user.ID, category.ID)

	res, err := s.repository.UpdatePost(updateRequest)
	if err != nil {
		s.logger.Error("Error updating post", zap.Error(err), zap.Int("PostID", requests.PostID))

		return nil, &response.ServiceError{Err: err, Description: fmt.Sprintf("Error updating post: %d", requests.PostID)}
	}

	_, err_file := s.file.UpdateFile(request.UpdateFileRequest{
		Folder:   category.Name,
		OldTitle: post.Title,
		NewTitle: res.Title,
		Content:  res.Content,
	})

	if err_file != nil {
		s.logger.Error("Error updating file", zap.Error(err_file))

		return nil, &response.ServiceError{Err: err, Description: "Error updating file"}
	}

	postResponse := response.PostResponse{
		ID:      res.ID,
		Title:   res.Title,
		Content: res.Content,
	}

	return &response.ServiceResponse{Data: postResponse}, nil
}

func (s *postService) Delete(id int) (*response.ServiceResponse, *response.ServiceError) {
	err := s.repository.DeletePost(id)

	if err != nil {
		s.logger.Error("Error deleting post", zap.Error(err), zap.Int("PostID", id))
		return nil, &response.ServiceError{Err: err, Description: fmt.Sprintf("Error deleting post: %d", id)}
	}

	return &response.ServiceResponse{Data: "Successfully Post deleted"}, nil
}
