package utils

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var tagPrefixMap = map[string]string{
	"required": "Required",
	"email":    "InvalidEmail",
	"min":      "ShouldMin",
	"max":      "ShouldMax",
	"len":      "ShouldLen",
	"eq":       "ShouldEq",
	"gt":       "ShouldGt",
	"gte":      "ShouldGte",
	"lt":       "ShouldLt",
	"lte":      "ShouldLte",
}

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

// i18n is a translation function
func i18n(msgID string, params ...interface{}) string {
	// implement the translation with msgID
	return msgID
}

// composeMsgID is a helper function to compose error message ID
func composeMsgID(e validator.FieldError) string {
	if prefix, ok := tagPrefixMap[e.Tag()]; ok {
		return fmt.Sprintf("%s%s", prefix, e.Field())
	}
	return ""
}

// translateError is a helper function to translate error from validator
func TranslateError(err error) (errTxt string) {
	validationErrors := err.(validator.ValidationErrors)
	for _, e := range validationErrors {
		errTxt = i18n(composeMsgID(e), e.Param())
		break
	}
	return
}

func ParseValidationErrors(err error) map[string]string {
	errors := make(map[string]string)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, vErr := range validationErrors {
			field := strings.ToLower(vErr.Field())
			tag := vErr.Tag()

			// Customizing the error message based on the tag
			switch tag {
			case "required":
				errors[field] = "This field is required"
			case "email":
				errors[field] = "Invalid email address"
			case "min":
				errors[field] = "Too short"
			case "max":
				errors[field] = "Too long"
			case "gte":
				errors[field] = "Value is too small"
			case "lte":
				errors[field] = "Value is too large"
			default:
				errors[field] = "Validation error"
			}
		}
	}

	return errors
}
