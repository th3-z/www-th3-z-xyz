package storage

import (
	"database/sql"
	_"github.com/mattn/go-sqlite3"
)

func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("Couldn't init database - db nil")
	}

	return db
}

func Migrate(db *sql.DB) {
	query := `
        CREATE TABLE IF NOT EXISTS servers (
            id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
            name VARCHAR NOT NULL,
            address VARCHAR NOT NULL,
            locked INT NOT NULL DEFAULT 0
        );
    `

	_, err := db.Exec(query)

	if err != nil {
		panic(err)
	}
}
