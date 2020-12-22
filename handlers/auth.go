package handlers

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"net/http"
	"time"

	"www-th3-z-xyz/models"
	"www-th3-z-xyz/storage"

	"github.com/labstack/echo"
)

const maxAttempts = 5
const maxAttemptsPeriod = 60 * 60 * 24

func signInAttempt(db *sql.DB, c echo.Context) int64 {
	query := `
		INSERT INTO signin_attempt (
			ip_address, insert_date
		) VALUES (
			?, ?
		)
	`
	storage.PreparedExec(
		db, query, c.RealIP(), time.Now().Unix(),
	)

	query = `
		SELECT
			COUNT(*)
		FROM
			signin_attempt
		WHERE
			insert_date >= ?
			AND ip_address = ?
	`
	result := storage.PreparedQueryRow(
		db, query, time.Now().Unix()-maxAttemptsPeriod, c.RealIP(),
	)

	var attempts int64
	result.Scan(&attempts)

	return attempts
}

func SignIn(c echo.Context) error {
	session := models.GetSession(c)
	defer session.Write(c)

	attempts := signInAttempt(storage.Db, c)
	if attempts > maxAttempts {
		return echo.NewHTTPError(http.StatusUnauthorized, "Exceeded attempts")
	}

	username := c.FormValue("username")
	h := sha256.New()
	h.Write([]byte(c.FormValue("password")))
	passwordHex := hex.EncodeToString(h.Sum(nil))

	session.UserId = models.AuthUser(storage.Db, username, passwordHex)

	if session.UserId > 0 {
		return c.JSON(http.StatusOK, nil)
	}
	return echo.NewHTTPError(http.StatusUnauthorized, "Incorrect sign in")
}

func SignOut(c echo.Context) error {
	session := models.GetSession(c)
	defer session.Write(c)

	session.Invalidate()

	return c.JSON(http.StatusOK, nil)
}
