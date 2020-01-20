package handlers

import (
	"github.com/labstack/echo"
	"net/http"
	"beta-th3-z-xyz/storage"
)

type Server struct {
	Name    string
	Service    string
	Address string
	WebUrl string
	Locked  int
	MaxPlayers  int
}

type Infrastructure struct {
	Hostname    string
	Address    string
	Os string
}

func getServers() []Server {
	var servers []Server

	rows := storage.PreparedQuery(
		storage.Db,
		"SELECT address, service, name, web_url, locked, max_players FROM server",
	)
	defer rows.Close()

	for rows.Next() {
		var server Server
		err := rows.Scan(
			&server.Address, &server.Service, &server.Name, &server.WebUrl,
			&server.Locked, &server.MaxPlayers,
		)
		if err != nil {
			panic(err)
		}

		servers = append(servers, server)
	}

	return servers
}

func getInfrastructure() []Infrastructure {
	var infrastructures []Infrastructure

	rows := storage.PreparedQuery(
		storage.Db,
		"SELECT hostname, address, os FROM infrastructure",
	)
	defer rows.Close()

	for rows.Next() {
		var infrastructure Infrastructure
		err := rows.Scan(
			&infrastructure.Hostname, &infrastructure.Address, &infrastructure.Os,
		)
		if err != nil {
			panic(err)
		}

		infrastructures = append(infrastructures, infrastructure)
	}

	return infrastructures
}

func Servers(c echo.Context) error {
	data := struct {
		Page Page
		Servers []Server
		Infrastructure []Infrastructure
	} {
		Page: Page {
			SelectedTab: 2,
			Title:       "Servers",
			Id:          "servers",
		},
		Servers: getServers(),
		Infrastructure: getInfrastructure(),
	}

	return c.Render(http.StatusOK, "servers/index", data)
}
