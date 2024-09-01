package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorResponse represents the structure for error responses
type ErrorResponse struct {
	Code    int `json:"code"`
	Message any `json:"message"`
}

// SuccessResponse represents the structure for successful responses
type SuccessResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

// NewErrorResponse sends a JSON response with the specified error message and HTTP status code
func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, ErrorResponse{
		Code:    statusCode,
		Message: message,
	})
	c.AbortWithStatus(statusCode)
}

// NewSuccessResponse sends a JSON response with the specified data and HTTP status code
func NewSuccessResponse(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, SuccessResponse{
		Code: statusCode,
		Data: data,
	})
}

func InvalidJSON(c *gin.Context) {
	c.JSON(http.StatusBadRequest, ErrorResponse{
		Code:    http.StatusBadRequest,
		Message: "Invalid JSON",
	})
}

func InvalidRequestData(c *gin.Context, errs map[string]string) {

	c.JSON(http.StatusBadRequest, ErrorResponse{
		Code:    http.StatusBadRequest,
		Message: errs,
	})
}
