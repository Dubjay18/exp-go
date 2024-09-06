package middlewares

import (
	"exp-go/internal/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Check if the user is authorized
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			utils.NewErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
			return
		}

		// Continue down the chain to handler etc
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			utils.NewErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
