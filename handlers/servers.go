package handlers

import (
	"github.com/labstack/echo"
	"net/http"
	"beta-th3-z-xyz/storage"
)

type Server struct {
	Name    string
	Address string
	Locked  int
}

func getServers() []Server {
	var servers []Server

	rows := storage.PreparedQuery(
		storage.Db,
		"SELECT address, name, locked FROM server",
	)
	defer rows.Close()

	for rows.Next() {
		var server Server
		err := rows.Scan(&server.Address, &server.Name, &server.Locked)
		if err != nil {
			panic(err)
		}

		servers = append(servers, server)
	}

	return servers
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
