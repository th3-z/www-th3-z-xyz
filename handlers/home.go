package handlers

import (
	"net/http"
	"www-th3-z-xyz/models"

	"github.com/labstack/echo"
)

func Index(c echo.Context) error {
	session := models.GetSession(c)
	defer models.WriteSession(c, session)

	data := struct {
		Page models.Page
	}{
		Page: models.Page{
			SelectedTab: 0,
			Title:       "Home",
			Id:          "home",
			Session:     session,
		},
	}

	return c.Render(http.StatusOK, "index", data)
}
