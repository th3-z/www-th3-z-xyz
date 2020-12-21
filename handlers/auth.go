package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"

	"www-th3-z-xyz/models"
	"www-th3-z-xyz/storage"

	"github.com/labstack/echo"
)

func Login(c echo.Context) error {
	session := models.GetSession(c)

	username := c.FormValue("username")

	h := sha256.New()
	h.Write([]byte(c.FormValue("password")))
	passwordHex := hex.EncodeToString(h.Sum(nil))

	session.UserId = models.AuthUser(storage.Db, username, passwordHex)

	models.WriteSession(c, session)

	if session.UserId > 0 {
		return c.JSON(http.StatusOK, nil)
	}
	return echo.NewHTTPError(http.StatusUnauthorized, "Incorrect sign in")
}
