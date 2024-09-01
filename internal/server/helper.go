package server

import (
	"fmt"
	"net/http"
)

type APIError struct {
	StatusCode int `json:"status_code"`
	Message    any `json:"message"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("api error: %d", e.StatusCode)
}
func NewAPIError(statusCode int, err error) *APIError {
	return &APIError{
		StatusCode: statusCode,
		Message:    err.Error(),
	}
}

func InvalidRequestData(errs map[string]string) APIError {
	return APIError{
		StatusCode: http.StatusUnprocessableEntity,
		Message:    errs,
	}
}

func InvalidJSON() APIError {
	return *NewAPIError(http.StatusBadRequest, fmt.Errorf("invalid json"))
}

//TODO: set up goose for migrations
//TODO: set up gin for routes
//TODO: set up make decorators for routes
