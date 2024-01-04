package service

import (
	"bookmark-backend/internal/domain/request"
	"bookmark-backend/internal/domain/response"
	"bookmark-backend/internal/repository"
	"bookmark-backend/pkg/hash"
	"bookmark-backend/pkg/logger"
	"fmt"

	"go.uber.org/zap"
)

type userService struct {
	repository repository.UserRepository
	hash       hash.Hashing
	logger     logger.Logger
}

func NewUserService(repository repository.UserRepository, hash hash.Hashing, logger logger.Logger) *userService {
	return &userService{repository: repository, hash: hash, logger: logger}
}

func (s *userService) GetUserAll() (*response.ServiceResponse, *response.ServiceError) {
	users, err := s.repository.FindAllUsers()
	if err != nil {
		s.logger.Error("Error fetching users", zap.Error(err))

		return nil, &response.ServiceError{Err: err, Description: "Error fetching users"}

	}

	return &response.ServiceResponse{Data: users}, nil
}

func (s *userService) GetUserByID(id int) (*response.ServiceResponse, *response.ServiceError) {
	user, err := s.repository.FindUserByID(id)
	if err != nil {
		s.logger.Error("Error fetching user", zap.Error(err), zap.Int("id", id))

		return nil, &response.ServiceError{Err: err, Description: fmt.Sprintf("Error fetching user: %d", id)}
	}

	return &response.ServiceResponse{Data: user}, nil
}

func (s *userService) CreateUser(request request.CreateUserRequest) (*response.ServiceResponse, *response.ServiceError) {
	hashing, err := s.hash.HashPassword(request.Password)

	if err != nil {
		s.logger.Error("Error hashing password", zap.Error(err), zap.String("password", request.Password))
		return nil, &response.ServiceError{
			Err:         err,
			Description: "Error hashing password",
		}
	}

	request.Password = hashing

	user, err := s.repository.CreateUser(request)

	if err != nil {
		s.logger.Error("Error creating user", zap.Error(err))
		return nil, &response.ServiceError{
			Err:         err,
			Description: "Error creating user",
		}
	}

	return &response.ServiceResponse{Data: user}, nil
}

func (s *userService) UpdateUser(id int, request request.UpdateUserRequest) (*response.ServiceResponse, *response.ServiceError) {
	user, err := s.repository.UpdateUser(id, request)

	if err != nil {
		s.logger.Error("Error updating user", zap.Error(err))
		return nil, &response.ServiceError{
			Err:         err,
			Description: "Error updating user",
		}
	}

	return &response.ServiceResponse{Data: user}, nil
}

func (s *userService) DeleteUser(id int) (*response.ServiceResponse, *response.ServiceError) {
	err := s.repository.DeleteUser(id)

	if err != nil {
		s.logger.Error("Error deleting user", zap.Error(err))
		return nil, &response.ServiceError{
			Err:         err,
			Description: "Error deleting user",
		}
	}

	return &response.ServiceResponse{Data: "User deleted"}, nil
}
