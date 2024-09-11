package response

import "github.com/labstack/echo/v4"

func ErrorResponse(ctx echo.Context, code int, err error) error {
	ctx.Response().Status = code
	return echo.NewHTTPError(code, err.Error())
}
