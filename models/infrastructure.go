package models

import (
	"www-th3-z-xyz/storage"
)

type Infrastructure struct {
	Hostname    string
	Address    string
	Os string
}

func GetInfrastructure() []Infrastructure {
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
