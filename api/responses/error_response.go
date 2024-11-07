package responses

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

// HTTPError is a custom error type that holds a status code and an error message
type HTTPError struct {
	StatusCode int
	Message    string
}

// Error allows HTTPError to satisfy the error interface
func (e *HTTPError) Error() string {
	return fmt.Sprintf("%d: %s", e.StatusCode, e.Message)
}

// NewHTTPError creates a new instance of HTTPError
func NewHTTPError(statusCode int, message string) *HTTPError {
	return &HTTPError{
		StatusCode: statusCode,
		Message:    message,
	}
}

// ErrorResponse sends a formatted JSON error response
func ErrorResponse(c echo.Context, statusCode int, err error) error {
	if httpErr, ok := err.(*HTTPError); ok {
		// Send error with the custom HTTP status code
		return c.JSON(httpErr.StatusCode, httpErr)
	}
	return c.JSON(statusCode, map[string]string{"message": err.Error()})
}
