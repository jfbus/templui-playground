package handlers

import (
	views "github.com/jfbus/templui-playground/templplayground/views"
	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return views.Base().Render(c.Request().Context(), c.Response().Writer)
}
