package handlers

import (
	"github.com/labstack/echo"
	"net/http"
)

func Articles(c echo.Context) error {
	data := struct {
		Page Page
	} {
		Page: Page {
			SelectedTab: 3,
			Title:       "Articles",
			Id:          "articles",
		},
	}

	return c.Render(http.StatusOK, "articles/index", data)
}
