package handlers

import (
	"www-th3-z-xyz/models"
	"github.com/labstack/echo"
	"net/http"
)

func Software(c echo.Context) error {
	data := struct {
		Page models.Page
		Software []models.Software
	} {
		Page: models.Page{
			SelectedTab: 1,
			Title:       "Software",
			Id:          "Software",
		},
		Software: models.GetSoftware(),
	}

	return c.Render(http.StatusOK, "software/index", data)
}
