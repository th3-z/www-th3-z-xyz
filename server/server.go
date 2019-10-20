package main

import (
	"database/sql"
	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
)

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("db nil")
	}

	return db
}

func migrate(db *sql.DB) {
	query := `
        CREATE TABLE IF NOT EXISTS tasks(
            id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
            name VARCHAR NOT NULL
        );
    `

	_, err := db.Exec(query)

	if err != nil {
		panic(err)
	}
}

func main() {
	db := initDB("storage.db")
	migrate(db)

    e := echo.New()

	e.Static("/", "../ui/dist")
	e.File("/", "../ui/dist/index.html")

	e.Logger.Fatal(e.Start(":5555"))
}
