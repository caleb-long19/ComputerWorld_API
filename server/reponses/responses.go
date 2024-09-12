package reponses

import "github.com/labstack/echo/v4"

func ErrorResponse(context echo.Context, code int, err error) error {
	context.Response().Status = code
	return echo.NewHTTPError(code, err.Error())
}
