package models

import (
	"www-th3-z-xyz/storage"
)

type Software struct {
	Name    string
	ProjectUrl    string
	RepoUrl string
	Description string
	Status string
}


func GetSoftware() []Software {
	var softwares []Software

	rows := storage.PreparedQuery(
		storage.Db,
		"SELECT name, project_url, repo_url, description, status FROM software",
	)
	defer rows.Close()

	for rows.Next() {
		var software Software
		err := rows.Scan(
			&software.Name, &software.ProjectUrl, &software.RepoUrl, &software.Description, &software.Status,
		)
		if err != nil {
			panic(err)
		}

		softwares = append(softwares, software)
	}

	return softwares
}

