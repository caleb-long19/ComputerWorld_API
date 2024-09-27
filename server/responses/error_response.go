package responses

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

func ErrorResponse(context echo.Context, code int, err error) error {
	// Set the status code for the response
	if context != nil {
		context.Response().WriteHeader(code)
	}

	// Standard error structure
	errorResponse := map[string]interface{}{
		"status": code,
		"error":  err.Error(),
	}

	// If context is present, return the JSON response
	if context != nil {
		return context.JSON(code, errorResponse)
	}

	// Return the error for testing or other non-echo purposes
	return echo.NewHTTPError(code, errorResponse)
}

type HTTPError struct {
	StatusCode int
	Message    string
}

// Implement the error interface
func (e *HTTPError) Error() string {
	return fmt.Sprintf("status %d: %s", e.StatusCode, e.Message)
}

func NewHTTPError(statusCode int, message string) error {
	return &HTTPError{
		StatusCode: statusCode,
		Message:    message,
	}
}
