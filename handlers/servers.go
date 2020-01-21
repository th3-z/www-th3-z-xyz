package handlers

import (
	"beta-th3-z-xyz/models"
	"github.com/labstack/echo"
	"net/http"
)

func Servers(c echo.Context) error {
	data := struct {
		Page           models.Page
		Servers        []models.Server
		Infrastructure []models.Infrastructure
	} {
		Page: models.Page{
			SelectedTab: 2,
			Title:       "Servers",
			Id:          "servers",
		},
		Servers: models.GetServers(),
		Infrastructure: models.GetInfrastructure(),
	}

	return c.Render(http.StatusOK, "servers/index", data)
}
