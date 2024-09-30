package responses

import (
	"fmt"
	"net/http"

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

// ErrorResponse formats and sends an error response with a custom status code using Echo
func ErrorResponse(c echo.Context, statusCode int, err error) error {
	return c.JSON(statusCode, map[string]string{
		"error": err.Error(),
	})
}

// SuccessResponse sends a success message back to the client using Echo
func SuccessResponse(c echo.Context, message string) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": message,
	})
}
