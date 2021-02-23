package handlers

import (
	"net/http"
	"www-th3-z-xyz/models"
	"www-th3-z-xyz/storage"

	"github.com/labstack/echo"
	malmodels "github.com/th3-z/malgo/models"
)

func Anime(c echo.Context) error {
	session := models.GetSession(c)
	defer session.Write(c)

	data := struct {
		Page    models.Page
		MalUser *malmodels.User
	}{
		Page: models.Page{
			SelectedTab: 4,
			Title:       "Anime",
			Id:          "anime",
			Session:     session,
		},
		MalUser: malmodels.SearchUser(storage.Db, "th3-z"),
	}

	return c.Render(http.StatusOK, "anime/index", data)
}
