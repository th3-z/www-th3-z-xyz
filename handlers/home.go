package handlers

import (
	"beta-th3-z-xyz/models"
	"github.com/labstack/echo"
	"net/http"
)

func Index(c echo.Context) error {
	data := struct {
		Page models.Page
	} {
		Page: models.Page{
			SelectedTab: 0,
			Title:       "Home",
			Id:          "home",
		},
	}

	return c.Render(http.StatusOK, "index", data)
}
