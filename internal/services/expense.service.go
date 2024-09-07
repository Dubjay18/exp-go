package services

import (
	"errors"
	"exp-go/internal/dto"
	"exp-go/internal/models"
	"exp-go/internal/repositories"
	"exp-go/internal/utils"
	"fmt"

	"time"

	"github.com/gin-gonic/gin"
)

type ExpenseService interface {
	AddExpense(c *gin.Context, req dto.AddExpenseRequest) (*dto.AddExpenseResponse, error)
	GetExpense(c *gin.Context, id string) (*dto.GetExpenseResponse, error)
	GetUserExpenses(c *gin.Context, id string) (*dto.GetUserExpensesResponse, error)
}

type DefaultExpenseService struct {
	repo repositories.ExpenseRepository
}

func NewExpenseService(repo repositories.ExpenseRepository) ExpenseService {
	return &DefaultExpenseService{
		repo: repo,
	}
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
	fmt.Println(claims)
	if claims == nil {
		return nil, errors.New("user claims not found")
	}
	claimsMap, ok := claims.(*utils.Claims)
	if !ok {
		return nil, errors.New("user claims not found")
	}
	userID := claimsMap.UserID

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

func (s *DefaultExpenseService) GetExpense(c *gin.Context, id string) (*dto.GetExpenseResponse, error) {
	expense, err := s.repo.GetExpenseByID(id)
	if err != nil {
		return nil, err
	}

	return &dto.GetExpenseResponse{
		ID:       expense.ID,
		Amount:   expense.Amount,
		Category: expense.Category,
		Note:     expense.Note,
		Date:     expense.Date.Format("2006-01-02"),
	}, nil
}

func (s *DefaultExpenseService) GetUserExpenses(c *gin.Context, id string) (*dto.GetUserExpensesResponse, error) {
	expenses, err := s.repo.GetExpensesByUserID(id)
	if err != nil {
		return nil, err
	}

	var response dto.GetUserExpensesResponse
	for _, expense := range expenses {
		response.Expenses = append(response.Expenses, dto.GetExpenseResponse{
			ID:       expense.ID,
			Amount:   expense.Amount,
			Category: expense.Category,
			Note:     expense.Note,
			Date:     expense.Date.Format("2006-01-02"),
		})
	}

	return &response, nil
}
