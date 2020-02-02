package handlers

import (
	"beta-th3-z-xyz/models"
	"github.com/labstack/echo"
	"net/http"
)

func Pastes(c echo.Context) error {
	data := struct {
		Page models.Page
		Pastes []models.Paste
	} {
		Page: models.Page{
			SelectedTab: 5,
			Title:       "Pastes",
			Id:          "pastes",
		},
		Pastes: models.GetPastes(),
	}

	return c.Render(http.StatusOK, "pastes/index", data)
}
