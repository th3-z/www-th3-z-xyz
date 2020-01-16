package handlers

import (
	"github.com/labstack/echo"
	"net/http"
)

func Index(c echo.Context) error {
	data := struct {
		Page Page
	} {
		Page: Page {
			SelectedTab: 0,
			Title:       "Home",
			Id:          "home",
		},
	}

	return c.Render(http.StatusOK, "index", data)
}
