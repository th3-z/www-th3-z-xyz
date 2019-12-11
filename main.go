package main

import (
	"github.com/labstack/echo"

	"./database"
	"./handlers"
	"html/template"
	"io"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func routes(e *echo.Echo) {
	e.GET("/hello", handlers.Hello)
}

func main() {
	db := database.InitDB("storage.db")
	database.Migrate(db)

	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}

	e := echo.New()
	e.Renderer = t
	routes(e)

	e.Logger.Fatal(e.Start(":5555"))
}
