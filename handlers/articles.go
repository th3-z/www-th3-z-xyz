package handlers

import (
	"html/template"
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

func Article(c echo.Context) error {
	session := models.GetSession(c)
	defer session.Write(c)

	data := struct {
		Page     models.Page
		Articles []models.Article
		Content  template.HTML
	}{
		Page: models.Page{
			SelectedTab: 5,
			Title:       "Articles",
			Id:          "articles",
			Session:     session,
		},
		Content: template.HTML(*models.GetArticles("templates/articles")[0].Content()),
	}

	return c.Render(http.StatusOK, "articles/base", data)
}
