package handlers

import (
	"beta-th3-z-xyz/models"
	"github.com/labstack/echo"
	"net/http"
)

func Projects(c echo.Context) error {
	data := struct {
		Page models.Page
		Projects []models.Project
	} {
		Page: models.Page{
			SelectedTab: 1,
			Title:       "Projects",
			Id:          "projects",
		},
		Projects: models.GetProjects(),
	}

	return c.Render(http.StatusOK, "projects/index", data)
}
