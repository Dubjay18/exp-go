package repositories

import (
	"exp-go/internal/database"
	"exp-go/internal/database/postgresql"
	"exp-go/internal/models"
)

type ExpenseRepository interface {
	CreateExpense(expense *models.Expense) error
	GetExpenseByID(id string) (*models.Expense, error)
	GetExpensesByUserID(userId string) ([]models.Expense, error)
	UpdateExpense(expense *models.Expense) error
	DeleteExpense(id uint) error
	GetTotalExpenses(filters map[string]interface{}) (float64, error)
}

type DefaultExpenseRepository struct {
	db database.Service
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
func (r *DefaultExpenseRepository) GetExpenseByID(id string) (*models.Expense, error) {
	expense := &models.Expense{}
	_, err := postgresql.SelectOneFromDb(r.db.Getpdb(), expense, id)
	if err != nil {
		return nil, err
	}

	return expense, nil
}

// GetExpensesByUserID gets all expenses by a user's ID
func (r *DefaultExpenseRepository) GetExpensesByUserID(userId string) ([]models.Expense, error) {
	expenses := []models.Expense{}
	err := postgresql.SelectAllFromDb(r.db.Getpdb(), "desc", &expenses, userId)
	if err != nil {
		return nil, err
	}
	return expenses, nil
}

// UpdateExpense updates an expense
func (r *DefaultExpenseRepository) UpdateExpense(expense *models.Expense) error {
	err := expense.Update(r.db.Getpdb())
	if err != nil {
		return err
	}
	return nil
}

// DeleteExpense deletes an expense
func (r *DefaultExpenseRepository) DeleteExpense(id uint) error {
	expense := &models.Expense{}
	expense.ID = id
	err := expense.Get(r.db.Getpdb())
	if err != nil {
		return err
	}
	err = postgresql.DeleteRecordFromDb(r.db.Getpdb(), expense)
	if err != nil {
		return err
	}
	return nil
}

func (r *DefaultExpenseRepository) GetTotalExpenses(filters map[string]interface{}) (float64, error) {
	totalExpenses := 0.0
	count, err := postgresql.CountSpecificRecords(r.db.Getpdb(), &models.Expense{}, filters)
	if err != nil {
		return 0, err
	}
	totalExpenses = float64(count)
	return totalExpenses, nil
}

func NewExpenseRepository(service database.Service) ExpenseRepository {
	return &DefaultExpenseRepository{
		db: service,
	}
}
