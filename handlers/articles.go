package handlers

import (
	"www-th3-z-xyz/models"
	"github.com/labstack/echo"
	"net/http"
)

func Articles(c echo.Context) error {
	data := struct {
		Page models.Page
		Articles []models.Article
	} {
		Page: models.Page{
			SelectedTab: 5,
			Title:       "Articles",
			Id:          "articles",
		},
		Articles: models.GetArticles("templates/articles"),
	}

	return c.Render(http.StatusOK, "articles/index", data)
}