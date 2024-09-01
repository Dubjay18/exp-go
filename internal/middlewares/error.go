package middlewares

import (
	"exp-go/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ErrorHandlingMiddleware is a middleware function for handling errors and recovery
func ErrorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				utils.NewErrorResponse(c, http.StatusInternalServerError, "Internal Server Error")
			}
		}()
		c.Next()
	}
}
