package handlers

import (
	"github.com/labstack/echo"
	"net/http"
)

type Server struct {
	Name    string
	Address string
}

func getServers() []Server {
	return []Server{{
			Name:    "vanilluxe",
			Address: "vanilluxe.th3-z.xyz",
		}, {
			Name:    "KF2",
			Address: "kf2.th3-z.xyz",
		},
	}
}

func Servers(c echo.Context) error {
	data := struct {
		Page Page
		Servers []Server
	} {
		Page: Page {
			SelectedTab: 2,
			Title:       "Servers",
			Id:          "servers",
		},
		Servers: getServers(),
	}

	return c.Render(http.StatusOK, "servers/index", data)
}
