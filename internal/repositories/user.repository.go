package repositories

import (
	"errors"
	"exp-go/internal/database"
	"exp-go/internal/database/postgresql"
	"exp-go/internal/models"
)

type UserRepository interface {
	CreateUser(user *models.User) error
}

type DefaultUserRepository struct {
	db database.Service
}

func (r *DefaultUserRepository) CreateUser(user *models.User) error {
	if ok := postgresql.CheckExists(r.db.Getpdb(), user, nil, nil); ok {
		return errors.New("user already exists")
	}
	err := user.Create(r.db.Getpdb())
	if err != nil {
		return err
	}
	return nil
}

func NewUserRepository(service database.Service) UserRepository {
	return &DefaultUserRepository{
		db: service,
	}
}
