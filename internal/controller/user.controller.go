package controller

import (
	"exp-go/internal/services"
	"exp-go/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type UserController struct {
	service   services.UserService
	Validator *validator.Validate
}

func (u *UserController) UserRegistration(c *gin.Context) {
	var req utils.RegisterUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.InvalidRequestData(c, utils.ParseValidationErrors(err))
		return
	}
	if err := u.Validator.Struct(req); err != nil {
		utils.InvalidRequestData(c, utils.ParseValidationErrors(err))
		return
	}
	resp, err := u.service.RegisterUser(c, &req)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.NewSuccessResponse(c, http.StatusCreated, resp)
}

func NewUserController(service services.UserService, Validator *validator.Validate) UserController {
	return UserController{service: service, Validator: Validator}
}
