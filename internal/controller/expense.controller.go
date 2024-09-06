package controller

import (
	"exp-go/internal/dto"
	"exp-go/internal/services"
	"exp-go/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type ExpenseController struct {
	service   services.ExpenseService
	Validator *validator.Validate
}

func NewExpenseController(service services.ExpenseService, Validator *validator.Validate) ExpenseController {
	return ExpenseController{service: service, Validator: Validator}
}

// AddExpense adds a new expense
func (e *ExpenseController) AddExpense(c *gin.Context) {
	var req dto.AddExpenseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.InvalidRequestData(c, utils.ParseValidationErrors(err))
		return
	}
	if err := e.Validator.Struct(req); err != nil {

		utils.InvalidRequestData(c, utils.ParseValidationErrors(err))
		return
	}
	resp, err := e.service.AddExpense(c, req)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	utils.NewSuccessResponse(c, http.StatusCreated, resp)

}
