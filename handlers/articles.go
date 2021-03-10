package handlers

import (
	"net/http"
	"www-th3-z-xyz/models"

	"io/ioutil"
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

func ArticleSrc(c echo.Context) error {
	session := models.GetSession(c)
	defer session.Write(c)

	path := strings.Split(c.Request().RequestURI, "/")
	articleName := path[len(path)-2]

	article := models.GetArticle(articleName)
	/*if err != nil { FIXME
		return c.NoContent(http.StatusBadRequest)
	}*/

	src, err := ioutil.ReadFile(article.SourceFilename)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, string(src))
}

func ArticleVisibility(c echo.Context) error {
	session := models.GetSession(c)
	defer session.Write(c)

	path := strings.Split(c.Request().RequestURI, "/")
	articleName := path[len(path)-2]
	visible := c.FormValue("visible")

	article := models.GetArticle(articleName)
	/*if err != nil { FIXME
		return c.NoContent(http.StatusBadRequest)
	}*/

	if visible == "true" {
		article.Meta.Visible = "1"
	} else {
		article.Meta.Visible = "0"
	}

	article.Meta.Write()

	return c.NoContent(http.StatusOK)
}
