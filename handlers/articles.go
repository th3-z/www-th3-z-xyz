package handlers

import (
	"net/http"
	"www-th3-z-xyz/models"

	"strings"

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

	path := strings.Split(c.Request().RequestURI, "/")
	articleName := path[len(path)-1]

	article := models.GetArticle(articleName)
	/*if err != nil { FIXME
		return c.NoContent(http.StatusBadRequest)
	}*/

	data := struct {
		Page    models.Page
		Article *models.Article
	}{
		Page: models.Page{
			SelectedTab: 5,
			Title:       "Articles",
			Id:          "articles",
			Session:     session,
		},
		Article: article,
	}

	return c.Render(http.StatusOK, "articles/base", data)
}
