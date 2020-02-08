package models

import (
	"www-th3-z-xyz/storage"
)

type Project struct {
	Name    string
	ProjectUrl    string
	RepoUrl string
	Description string
	Status string
}


func GetProjects() []Project {
	var projects []Project

	rows := storage.PreparedQuery(
		storage.Db,
		"SELECT name, project_url, repo_url, description, status FROM project",
	)
	defer rows.Close()

	for rows.Next() {
		var project Project
		err := rows.Scan(
			&project.Name, &project.ProjectUrl, &project.RepoUrl, &project.Description, &project.Status,
		)
		if err != nil {
			panic(err)
		}

		projects = append(projects, project)
	}

	return projects
}

