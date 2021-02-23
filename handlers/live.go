package handlers

import (
	"net/http"
	"www-th3-z-xyz/models"

	"github.com/labstack/echo"
)

func Live(c echo.Context) error {
	session := models.GetSession(c)
	defer session.Write(c)

	data := struct {
		Page models.Page
	}{
		Page: models.Page{
			SelectedTab: 6,
			Title:       "Live",
			Id:          "Live",
			Session:     session,
		},
	}

	return c.Render(http.StatusOK, "live/index", data)
}
