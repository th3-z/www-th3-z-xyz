package handlers

import (
	"net/http"
	"www-th3-z-xyz/models"

	"github.com/labstack/echo"
)

func Articles(c echo.Context) error {
	session := models.GetSession(c)
	defer session.Write(c)

	data := struct {
		Page     models.Page
		Articles []models.Article
	}{
		Page: models.Page{
			SelectedTab: 5,
			Title:       "Articles",
			Id:          "articles",
			Session:     session,
		},
		Articles: models.GetArticles("templates/articles"),
	}

	return c.Render(http.StatusOK, "articles/index", data)
}
