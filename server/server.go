package main

import (
    "database/sql"
    "html/template"
    "io"
    "net/http"

    "github.com/labstack/echo"
    _ "github.com/mattn/go-sqlite3"
)


type Template struct {
    templates *template.Template
}



func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

func Hello(c echo.Context) error {
    return c.Render(http.StatusOK, "hello", "World")
}

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

    t := &Template{
        templates: template.Must(template.ParseGlob("public/views/*.html")),
    }


    e := echo.New()
    e.Renderer = t

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

    e.GET("/hello", Hello)

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
