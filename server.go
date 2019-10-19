package main


import (
	"net/http"
    "database/sql"

	"github.com/labstack/echo"
     _ "github.com/mattn/go-sqlite3"
)

func initDB(filepath string) *sql.DB {
    db, err := sql.Open("sqlite3", "storage.db")

    if err != nil {
        panic(err)
    }

    if db == nil {
        panic("db nil")
    }

    return db
}

func migrate(db *sql.DB) {
    sql := `
    CREATE TABLE IF NOT EXISTS tasks(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        name VARCHAR NOT NULL
    );
    `

    _, err := db.Exec(sql)

    if err != nil {
        panic(err)
    }
}

func main() {
    db := initDB("storage.db")
    migrate(db)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
    })

    e.GET("/tasks",
        func(c echo.Context) error {
            return c.JSON(200, "GET Tasks")
        })
    e.PUT("/tasks",
        func(c echo.Context) error {
            return c.JSON(200, "PUT Tasks")
        })
    e.DELETE("/tasks/:id",
        func(c echo.Context) error {
            return c.JSON(200, "DELETE Task "+c.Param("id"))
        })

	e.Logger.Fatal(e.Start(":5555"))
}


