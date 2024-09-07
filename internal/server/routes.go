package server

import (
	"exp-go/internal/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(middlewares.ErrorHandlingMiddleware())

	r.GET("/", s.HelloWorldHandler)
	apiVersion := r.Group("api/v1/")
	user := apiVersion.Group("user")
	{
		user.POST("/register", s.UserController.UserRegistration)
		user.POST("/login", s.UserController.UserLogin)
	}

	expense := apiVersion.Group("expense")
	expense.Use(middlewares.AuthMiddleware())
	{
		expense.POST("/add", s.ExpenseController.AddExpense)
		expense.GET("/:id", s.ExpenseController.GetExpense)
	}

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"
	c.JSON(http.StatusOK, resp)
}
