package services

import (
	"exp-go/internal/dto"
	"exp-go/internal/models"
	"exp-go/internal/repositories"
	"exp-go/internal/utils"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	RegisterUser(c *gin.Context, req *utils.RegisterUserRequest) (*dto.RegisterUserResponse, error)
}

type DefaultUserService struct {
	repo repositories.UserRepository
}

func (s *DefaultUserService) RegisterUser(c *gin.Context, req *utils.RegisterUserRequest) (*dto.RegisterUserResponse, error) {
	user := models.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}

	if err := user.HashPassword(req.Password); err != nil {
		return nil, err
	}
	user.ID = utils.GenerateUUID()
	err := s.repo.CreateUser(&user)
	if err != nil {
		return nil, err
	}
	token := utils.GenerateJWT(user.ID, user.Username)

	return &dto.RegisterUserResponse{
		Message: "User registered successfully",
		Token:   token,
	}, nil
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &DefaultUserService{
		repo: repo,
	}
}
