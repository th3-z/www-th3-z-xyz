package handlers

import (
	"beta-th3-z-xyz/models"
	"github.com/labstack/echo"
	"net/http"
)

func Pastes(c echo.Context) error {
	data := struct {
		Page models.Page
	} {
		Page: models.Page{
			SelectedTab: 5,
			Title:       "Pastes",
			Id:          "pastes",
		},
	}

	return c.Render(http.StatusOK, "pastes/index", data)
}
