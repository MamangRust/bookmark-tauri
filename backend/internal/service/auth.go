package service

import (
	"bookmark-backend/internal/domain/request"
	"bookmark-backend/internal/domain/response"
	"bookmark-backend/internal/repository"
	"bookmark-backend/pkg/auth"
	"bookmark-backend/pkg/hash"
	"bookmark-backend/pkg/logger"
	"fmt"

	"go.uber.org/zap"
)

type authService struct {
	repository repository.UserRepository
	hash       hash.Hashing
	token      auth.TokenManager
	logger     logger.Logger
}

func NewAuthService(repository repository.UserRepository, hash hash.Hashing, token auth.TokenManager, logger logger.Logger) *authService {
	return &authService{repository: repository, hash: hash, token: token, logger: logger}
}

func (s *authService) Login(request request.LoginUserRequest) (*request.Token, *response.ServiceError) {
	user, err := s.repository.FindByEmail(request.Email)
	if err != nil {
		s.logger.Error("Error fetching user", zap.Error(err), zap.String("email", request.Email))
		return nil, &response.ServiceError{
			Err:         err,
			Description: fmt.Sprintf("Error fetching user: %s", request.Email),
		}
	}

	passwordErr := s.hash.ComparePassword(request.Password, user.Password)
	if passwordErr != nil {
		s.logger.Error("Error comparing password", zap.Error(passwordErr), zap.String("password", request.Password))
		return nil, &response.ServiceError{
			Err:         passwordErr,
			Description: "Error comparing password",
		}
	}

	token, tokenErr := s.createJwt(int(user.ID))
	if tokenErr != nil {
		s.logger.Error("Error creating token", zap.Error(tokenErr))
		return nil, &response.ServiceError{
			Err:         tokenErr,
			Description: "Error creating token",
		}
	}

	return token, nil
}

func (s *authService) Register(requests request.RegisterUserRequest) (*response.ServiceResponse, *response.ServiceError) {
	var createRequest request.CreateUserRequest

	hashing, err := s.hash.HashPassword(requests.Password)

	if err != nil {
		s.logger.Error("Error hashing password", zap.Error(err), zap.String("password", requests.Password))
		return nil, &response.ServiceError{
			Err:         err,
			Description: "Error hashing password",
		}
	}

	createRequest.Username = requests.Username
	createRequest.Email = requests.Email
	createRequest.Password = hashing

	user, err := s.repository.CreateUser(createRequest)

	if err != nil {
		s.logger.Error("Error creating user", zap.Error(err), zap.Any("request", requests))
		return nil, &response.ServiceError{
			Err:         err,
			Description: "Error creating user",
		}
	}

	return &response.ServiceResponse{
		Data: user,
	}, nil
}

func (s *authService) createJwt(id int) (*request.Token, *response.ServiceError) {
	var res request.Token
	tokenString, err := s.token.NewJwtToken(id)
	if err != nil {
		s.logger.Error("Error creating jwt token", zap.Error(err))
		return nil, &response.ServiceError{
			Err:         err,
			Description: "Error creating jwt token",
		}
	}
	res.Jwt = tokenString
	return &res, nil
}
