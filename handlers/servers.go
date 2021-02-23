package handlers

import (
	"net/http"
	"www-th3-z-xyz/models"

	"github.com/labstack/echo"
)

func Servers(c echo.Context) error {
	session := models.GetSession(c)
	defer session.Write(c)

	data := struct {
		Page           models.Page
		Servers        []models.Server
		Infrastructure []models.Infrastructure
	}{
		Page: models.Page{
			SelectedTab: 2,
			Title:       "Servers",
			Id:          "servers",
			Session:     session,
		},
		Servers:        models.GetServers(),
		Infrastructure: models.GetInfrastructure(),
	}

	return c.Render(http.StatusOK, "servers/index", data)
}
