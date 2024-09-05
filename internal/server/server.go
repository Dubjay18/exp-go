package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv/autoload"

	"exp-go/internal/controller"
	"exp-go/internal/database"
	"exp-go/internal/repositories"
	"exp-go/internal/services"
)

type Server struct {
	port           int
	db             database.Service
	UserController controller.UserController
}

func NewServer() *http.Server {
	validator := validator.New()
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	dbInstance := database.New()
	userRepo := repositories.NewUserRepository(dbInstance)
	userService := services.NewUserService(userRepo)
	userController := controller.NewUserController(userService, validator)
	NewServer := &Server{
		port:           port,
		UserController: userController,
		db:             dbInstance,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
