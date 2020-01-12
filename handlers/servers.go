package handlers

import (
	"github.com/labstack/echo"
	"net/http"
)

func Servers(c echo.Context) error {
	return c.Render(http.StatusOK, "servers/index", "")
}
