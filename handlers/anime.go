package handlers

import (
	"www-th3-z-xyz/models"
	malmodels"github.com/th3-z/malgo/models"
	"www-th3-z-xyz/storage"
	"github.com/labstack/echo"
	"net/http"
)

func Anime(c echo.Context) error {
	data := struct {
		Page models.Page
		MalUser *malmodels.User
	} {
		Page: models.Page {
			SelectedTab: 4,
			Title:       "Anime",
			Id:          "anime",
		},
        MalUser: malmodels.SearchUser(storage.Db, "th3-z"),
	}

	return c.Render(http.StatusOK, "anime/index", data)
}
