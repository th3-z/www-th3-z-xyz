package handlers

import (
	"www-th3-z-xyz/models"
	"github.com/labstack/echo"
	"net/http"
)

func Articles(c echo.Context) error {
	data := struct {
		Page models.Page
	} {
		Page: models.Page{
			SelectedTab: 3,
			Title:       "Articles",
			Id:          "articles",
		},
	}

	return c.Render(http.StatusOK, "articles/index", data)
}
