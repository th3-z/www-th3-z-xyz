package handlers

import (
	"github.com/labstack/echo"
	"net/http"
)

func Projects(c echo.Context) error {
	return c.Render(http.StatusOK, "projects/index", "")
}
