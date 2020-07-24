package handlers

import (
	"www-th3-z-xyz/models"
	"github.com/labstack/echo"
	"net/http"
)

func Live(c echo.Context) error {
	data := struct {
		Page models.Page
	} {
		Page: models.Page{
			SelectedTab: 6,
			Title:       "Live",
			Id:          "Live",
		},
	}

	return c.Render(http.StatusOK, "live/index", data)
}
