package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"www-th3-z-xyz/models"
	"www-th3-z-xyz/storage"

	"github.com/labstack/echo"
)

func Pastes(c echo.Context) error {
	data := struct {
		Page   models.Page
		Pastes []models.Paste
	}{
		Page: models.Page{
			SelectedTab: 7,
			Title:       "Pastes",
			Id:          "pastes",
		},
		Pastes: models.GetPastes(),
	}

	c.Response().Header().Set("Cache-Control", "no-store, must-revalidate")
	c.Response().Header().Set("Expires", "0")

	return c.Render(http.StatusOK, "pastes/index", data)
}

func NewPaste(c echo.Context) error {
	content := []byte(c.FormValue("content"))

	h := sha256.New()
	h.Write([]byte(c.RealIP()))
	uploaderId := hex.EncodeToString(h.Sum(nil))

	paste, err := models.NewPaste(storage.Db, content, uploaderId)
	if err != nil {
		return err
	}

	return c.Redirect(302, "files/"+paste.Filename)
}
