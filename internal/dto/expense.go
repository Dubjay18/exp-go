package dto

import (
	"errors"
	"time"
)

type AddExpenseRequest struct {
	Amount   float64 `json:"amount" binding:"required"`
	Category string  `json:"category" binding:"required,oneof=Groceries Leisure Electronics Utilities Clothing Health Others"`
	Note     string  `json:"note"`
	Date     string  `json:"date" binding:"required"`
}

type AddExpenseResponse struct {
	Message string `json:"message"`
}

type GetExpenseResponse struct {
	ID       uint    `json:"id"`
	Amount   float64 `json:"amount"`
	Category string  `json:"category"`
	Note     string  `json:"note"`
	Date     string  `json:"date"`
}

type GetUserExpensesResponse struct {
	Expenses []GetExpenseResponse `json:"expenses"`
}

// ParseDate parses and validates the date
func (input *AddExpenseRequest) ParseDate() (time.Time, error) {
	const layout = "2006-01-02" // Defining the date format layout (YYYY-MM-DD)
	parsedDate, err := time.Parse(layout, input.Date)
	if err != nil {
		return time.Time{}, errors.New("invalid date format, use YYYY-MM-DD")
	}
	return parsedDate, nil
}
