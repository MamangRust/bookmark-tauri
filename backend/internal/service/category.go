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

type categoryService struct {
	categoryMapper mapper.CategoryMapping
	file           fileService
	folder         folderService
	repository     repository.CategoryRepository
	logger         logger.Logger
}

func NewCategoryService(categoryMapper mapper.CategoryMapping, file fileService, folder folderService, repository repository.CategoryRepository, logger logger.Logger) *categoryService {
	return &categoryService{categoryMapper: categoryMapper, file: file, folder: folder, repository: repository, logger: logger}
}

func (s *categoryService) GetAll() (*response.ServiceResponse, *response.ServiceError) {
	res, err := s.repository.FindAllCategory()

	if err != nil {
		s.logger.Error("Error fetching all categories", zap.Error(err))
		return nil, &response.ServiceError{Err: err, Description: "Error fetching all categories"}
	}

	categoriesResponse := s.categoryMapper.MapToCategoriesResponse(res)

	return &response.ServiceResponse{
		Data: categoriesResponse,
	}, nil
}

func (s *categoryService) GetByID(id int) (*response.ServiceResponse, *response.ServiceError) {
	res, err := s.repository.FindCategoryByID(id)

	if err != nil {
		s.logger.Error("Error fetching category by ID", zap.Error(err), zap.Int("CategoryID", id))
		return nil, &response.ServiceError{Err: err, Description: fmt.Sprintf("Error fetching category by ID: %d", id)}
	}

	categoryResponse := s.categoryMapper.MapToCategoryResponse(res)

	return &response.ServiceResponse{Data: categoryResponse}, nil
}

func (s *categoryService) Create(request request.CreateCategoryRequest) (*response.ServiceResponse, *response.ServiceError) {
	createRequest := s.categoryMapper.MapToCategoryModel(request)

	res, err := s.repository.CreateCategory(createRequest)

	if err != nil {
		s.logger.Error("Error creating category", zap.Error(err), zap.Any("Request", request))
		return nil, &response.ServiceError{Err: err, Description: "Error creating category"}
	}

	categoryResponse := s.categoryMapper.MapToCategoryResponse(res)

	return &response.ServiceResponse{Data: categoryResponse}, nil
}

func (s *categoryService) Update(request request.UpdateCategoryRequest) (*response.ServiceResponse, *response.ServiceError) {
	res, err := s.repository.FindCategoryByID(request.CategoryID)

	if err != nil {
		s.logger.Error("Error fetching category by ID", zap.Error(err), zap.Int("CategoryID", request.CategoryID))
		return nil, &response.ServiceError{Err: err, Description: fmt.Sprintf("Error fetching category by ID: %d", request.CategoryID)}
	}

	if err := func() error {

		if err := s.file.DeleteFile(res.Image); err != nil {
			s.logger.Error("Error Delete File:", zap.Error(err))
			return err
		}

		if err := s.folder.DeleteFolder(res.Name); err != nil {
			s.logger.Error("Error Update Folder: ", zap.Error(err))
			return err
		}

		return nil
	}(); err != nil {
		return nil, &response.ServiceError{Err: err, Description: "Error handling file and folder operations"}
	}

	updateRequest := s.categoryMapper.MapToUpdateCategoryModel(request)

	res_update, err := s.repository.UpdateCategory(updateRequest)

	if err != nil {
		s.logger.Error("Error updating category", zap.Error(err), zap.Any("Request", request))
		return nil, &response.ServiceError{Err: err, Description: "Error updating category"}
	}

	categoryResponse := s.categoryMapper.MapToCategoryResponse(res_update)

	return &response.ServiceResponse{Data: categoryResponse}, nil
}

func (s *categoryService) Delete(id int) (*response.ServiceResponse, *response.ServiceError) {
	res, err := s.repository.FindCategoryByID(id)

	if err != nil {
		s.logger.Error("Error fetching category by ID", zap.Error(err), zap.Int("CategoryID", id))
		return nil, &response.ServiceError{Err: err, Description: fmt.Sprintf("Error fetching category by ID: %d", id)}
	}

	err = s.repository.DeleteCategory(id)

	if err != nil {
		s.logger.Error("Error deleting category", zap.Error(err), zap.Int("CategoryID", id))
		return nil, &response.ServiceError{Err: err, Description: fmt.Sprintf("Error deleting category: %d", id)}
	}

	err = s.folder.DeleteFolder(res.Name)

	if err != nil {
		s.logger.Error("Error deleting category", zap.Error(err), zap.Int("CategoryID", id))
		return nil, &response.ServiceError{Err: err, Description: fmt.Sprintf("Error deleting category: %d", id)}
	}

	return &response.ServiceResponse{Data: "Category deleted"}, nil
}
