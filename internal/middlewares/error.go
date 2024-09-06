package middlewares

import (
	"exp-go/internal/utils"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

// ErrorHandlingMiddleware is a middleware function for handling errors and recovery
func ErrorHandlingMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic recovered: %v\nStack trace:\n%s\n", err, debug.Stack())
				utils.NewErrorResponse(c, http.StatusInternalServerError, "Internal Server Error")
			}
		}()
		c.Next()
	}
}
