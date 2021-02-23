package handlers

import (
	"net/http"
	"www-th3-z-xyz/models"

	"github.com/labstack/echo"
)

func Software(c echo.Context) error {
	session := models.GetSession(c)
	defer session.Write(c)

	data := struct {
		Page     models.Page
		Software []models.Software
	}{
		Page: models.Page{
			SelectedTab: 1,
			Title:       "Software",
			Id:          "Software",
			Session:     session,
		},
		Software: models.GetSoftware(),
	}

	return c.Render(http.StatusOK, "software/index", data)
}
