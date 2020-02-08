package models

import (
	"www-th3-z-xyz/storage"
)

type Server struct {
	Name    string
	Service    string
	Address string
	WebUrl string
	Locked  int
	MaxPlayers  int
}


func GetServers() []Server {
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

