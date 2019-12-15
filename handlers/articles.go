package handlers

import (
	"github.com/labstack/echo"
	"net/http"
)

func Articles(c echo.Context) error {
	return c.Render(http.StatusOK, "articles/index", "")
}
