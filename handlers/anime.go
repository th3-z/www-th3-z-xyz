package handlers

import (
	"beta-th3-z-xyz/models"
	malmodels"github.com/th3-z/mal-sqlite-migrate/models"
	"beta-th3-z-xyz/storage"
	"github.com/labstack/echo"
	"net/http"
)

func Anime(c echo.Context) error {
	data := struct {
		Page models.Page
		AnimeList []malmodels.Anime
	} {
		Page: models.Page{
			SelectedTab: 4,
			Title:       "Anime",
			Id:          "anime",
		},
        AnimeList: malmodels.GetAnimeList(storage.Db),
	}

	return c.Render(http.StatusOK, "anime/index", data)
}
