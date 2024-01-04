package service

import (
	"bookmark-backend/internal/domain/request"
	"bookmark-backend/internal/domain/response"
	"bookmark-backend/internal/repository"
	"bookmark-backend/pkg/logger"
	"fmt"

	"go.uber.org/zap"
)

type categoryService struct {
	repository repository.CategoryRepository
	logger     logger.Logger
}

func NewCategoryService(repository repository.CategoryRepository, logger logger.Logger) *categoryService {
	return &categoryService{repository: repository, logger: logger}
}

func (s *categoryService) GetAll() (*response.ServiceResponse, *response.ServiceError) {
	res, err := s.repository.FindAllCategory()

	if err != nil {
		s.logger.Error("Error fetching all categories", zap.Error(err))
		return nil, &response.ServiceError{Err: err, Description: "Error fetching all categories"}
	}

	return &response.ServiceResponse{
		Data: res,
	}, nil
}

func (s *categoryService) GetByID(id int) (*response.ServiceResponse, *response.ServiceError) {
	res, err := s.repository.FindCategoryByID(id)

	if err != nil {
		s.logger.Error("Error fetching category by ID", zap.Error(err), zap.Int("CategoryID", id))
		return nil, &response.ServiceError{Err: err, Description: fmt.Sprintf("Error fetching category by ID: %d", id)}
	}

	return &response.ServiceResponse{Data: res}, nil
}

func (s *categoryService) Create(request request.CreateCategoryRequest) (*response.ServiceResponse, *response.ServiceError) {
	res, err := s.repository.CreateCategory(request)

	if err != nil {
		s.logger.Error("Error creating category", zap.Error(err), zap.Any("Request", request))
		return nil, &response.ServiceError{Err: err, Description: "Error creating category"}
	}

	return &response.ServiceResponse{Data: res}, nil
}

func (s *categoryService) Update(request request.UpdateCategoryRequest) (*response.ServiceResponse, *response.ServiceError) {
	res, err := s.repository.UpdateCategory(request)

	if err != nil {
		s.logger.Error("Error updating category", zap.Error(err), zap.Any("Request", request))
		return nil, &response.ServiceError{Err: err, Description: "Error updating category"}
	}

	return &response.ServiceResponse{Data: res}, nil
}

func (s *categoryService) Delete(id int) (*response.ServiceResponse, *response.ServiceError) {
	err := s.repository.DeleteCategory(id)

	if err != nil {
		s.logger.Error("Error deleting category", zap.Error(err), zap.Int("CategoryID", id))
		return nil, &response.ServiceError{Err: err, Description: fmt.Sprintf("Error deleting category: %d", id)}
	}

	return &response.ServiceResponse{Data: "Category deleted"}, nil
}
