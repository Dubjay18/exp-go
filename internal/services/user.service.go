package services

import (
	"exp-go/internal/repositories"
	"exp-go/internal/utils"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	RegisterUser(c *gin.Context, req *utils.RegisterUserRequest) (interface{}, error)
}

type DefaultUserService struct {
	repo *repositories.UserRepository
}

func (s *DefaultUserService) RegisterUser(c *gin.Context, req *utils.RegisterUserRequest) (interface{}, error) {
	return nil, nil
}

func NewUserService(repo *repositories.UserRepository) UserService {
	return &DefaultUserService{
		repo: repo,
	}
}
