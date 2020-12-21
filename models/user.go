package models

import (
	"database/sql"
	"www-th3-z-xyz/storage"
)

func AuthUser(db *sql.DB, username string, password string) int64 {
	query := `
		SELECT
			id
		FROM
			person
		WHERE
			username = ?
			AND password = ?
	`
	row := storage.PreparedQueryRow(db, query, username, password)

	var UserId int64
	row.Scan(&UserId)

	return UserId
}
