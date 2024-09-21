package models

import (
	"gorm.io/gorm"
	"time"
)

type Expense struct {
	gorm.Model
	UserID   string    `json:"user_id"`
	Amount   float64   `json:"amount"`
	Category string    `json:"category"`
	Note     string    `json:"note"`
	Date     time.Time `json:"date"`
}

// Category constants for validation
var ValidCategories = []string{
	"Groceries", "Leisure", "Electronics", "Utilities", "Clothing", "Health", "Others",
}

type ExpenseFilter struct {
	StartDate time.Time
	EndDate   time.Time
	Category  string
}

// Create
func (expense *Expense) Create(db *gorm.DB) error {
	err := db.Create(expense).Error
	if err != nil {
		return err
	}
	return nil
}

//update

func (expense *Expense) Update(db *gorm.DB) error {
	err := db.Save(expense).Error
	if err != nil {
		return err
	}
	return nil
}

// GEt
func (expense *Expense) Get(db *gorm.DB) error {
	err := db.First(expense, expense.ID).Error
	if err != nil {
		return err
	}
	return nil
}
