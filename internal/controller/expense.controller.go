package controller

import (
	"exp-go/internal/dto"
	"exp-go/internal/services"
	"exp-go/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
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
	log.Println("Service is nil:", e.service == nil)
	log.Println("Request Data:", req)
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

func (e *ExpenseController) GetExpense(c *gin.Context) {
	id := c.Param("id")
	resp, err := e.service.GetExpense(c, id)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	utils.NewSuccessResponse(c, http.StatusOK, resp)
}

func (e *ExpenseController) GetUserExpenses(c *gin.Context) {
	id := c.Param("id")
	resp, err := e.service.GetUserExpenses(c, id)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	utils.NewSuccessResponse(c, http.StatusOK, resp)
}
