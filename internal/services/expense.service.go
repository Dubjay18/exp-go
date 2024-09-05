package services

import (
	"errors"
	"exp-go/internal/dto"
	"exp-go/internal/models"
	"exp-go/internal/repositories"
	"time"

	"github.com/gin-gonic/gin"
)

type ExpenseService interface {
	AddExpense(c *gin.Context, req dto.AddExpenseRequest) (*dto.AddExpenseResponse, error)
}

type DefaultExpenseService struct {
	repo repositories.ExpenseRepository
}

func NewExpenseService(repo repositories.ExpenseRepository) ExpenseService {
	return &DefaultExpenseService{}
}

// CreateExpense creates a new expense
func (s *DefaultExpenseService) AddExpense(c *gin.Context, req dto.AddExpenseRequest) (*dto.AddExpenseResponse, error) {
	date, err := req.ParseDate()
	if err != nil {
		return nil, err
	}
	if date.IsZero() {
		date = time.Now()
	}
	claims, _ := c.Get("claims")
	if claims == nil {
		return nil, errors.New("user claims not found")
	}
	claimsMap, ok := claims.(map[string]interface{})
	if !ok {
		return nil, errors.New("user claims not found")
	}
	userID, ok := claimsMap["userID"].(string)
	if !ok {
		return nil, errors.New("user claims not found")
	}

	expense := models.Expense{
		UserID:   userID,
		Amount:   req.Amount,
		Category: req.Category,
		Note:     req.Note,
		Date:     date,
	}
	err = s.repo.CreateExpense(&expense)
	if err != nil {
		return nil, err
	}

	return &dto.AddExpenseResponse{
		Message: "Expense added successfully",
	}, nil
}
