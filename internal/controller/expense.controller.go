package controller

import (
	"exp-go/internal/services"

	"github.com/go-playground/validator/v10"
)

type ExpenseController struct {
	service   services.ExpenseService
	Validator *validator.Validate
}

func NewExpenseController(service services.ExpenseService, Validator *validator.Validate) ExpenseController {
	return ExpenseController{service: service, Validator: Validator}
}

// AddExpense adds a new expense
