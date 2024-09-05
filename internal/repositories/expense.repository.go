package repositories

import (
	"exp-go/internal/database"
	"exp-go/internal/models"
)

type ExpenseRepository interface {
	CreateExpense(expense *models.Expense) error
}

type DefaultExpenseRepository struct {
	db database.Service
}

func NewExpenseRepository(service database.Service) ExpenseRepository {
	return &DefaultExpenseRepository{
		db: service,
	}
}

// CreateExpense creates a new expense
func (r *DefaultExpenseRepository) CreateExpense(expense *models.Expense) error {
	err := expense.Create(r.db.Getpdb())
	if err != nil {
		return err
	}
	return nil
}

// GetExpenseByID gets an expense by its ID
func (r *DefaultExpenseRepository) GetExpenseByID() error {
	return nil
}

// GetExpensesByUserID gets all expenses by a user's ID
func (r *DefaultExpenseRepository) GetExpensesByUserID() error {
	return nil
}

// UpdateExpense updates an expense
func (r *DefaultExpenseRepository) UpdateExpense() error {
	return nil
}

// DeleteExpense deletes an expense
func (r *DefaultExpenseRepository) DeleteExpense() error {
	return nil
}

// GetTotalExpenses gets the total expenses of a user
func (r *DefaultExpenseRepository) GetTotalExpenses() error {
	return nil
}

// GetTotalExpensesByCategory gets the total expenses of a user by category
func (r *DefaultExpenseRepository) GetTotalExpensesByCategory() error {
	return nil
}

// GetTotalExpensesByDate gets the total expenses of a user by date
func (r *DefaultExpenseRepository) GetTotalExpensesByDate() error {
	return nil
}

// GetTotalExpensesByMonth gets the total expenses of a user by month
func (r *DefaultExpenseRepository) GetTotalExpensesByMonth() error {
	return nil
}

// GetTotalExpensesByYear gets the total expenses of a user by year
func (r *DefaultExpenseRepository) GetTotalExpensesByYear() error {
	return nil
}
