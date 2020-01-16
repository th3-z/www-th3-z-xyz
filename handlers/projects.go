package handlers

import (
	"github.com/labstack/echo"
	"net/http"
)

func Projects(c echo.Context) error {
	data := struct {
		Page Page
	} {
		Page: Page {
			SelectedTab: 1,
			Title:       "Projects",
			Id:          "projects",
		},
	}

	return c.Render(http.StatusOK, "projects/index", data)
}
