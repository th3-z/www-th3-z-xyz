package main

import (
	"github.com/labstack/echo"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"./storage"
	"./handlers"
	"html/template"
	"io"
)

type Template struct {
	templates *template.Template
}

func FindAndRenderGraphs(rootDir string, funcMap template.FuncMap) (*template.Template, error) {
    /*
     * Render DOT files from /static/graphs to /static/images/graphs
     */
    return nil
}

func FindAndRenderMarkdown(rootDir string, funcMap template.FuncMap) (*template.Template, error) {
    /*
     * Render Md files from /static/markdown to /templates
     */
	cleanRoot := filepath.Clean(rootDir)
	pfx := len(cleanRoot)+1
	root := template.New("")

	err := filepath.Walk(cleanRoot, func(path string, info os.FileInfo, e1 error) error {
		if !info.IsDir() && strings.HasSuffix(path, ".html") {
			if e1 != nil {
				return e1
			}

			b, e2 := ioutil.ReadFile(path)
			if e2 != nil {
				return e2
			}

			name := path[pfx:]
			t := root.New(name).Funcs(funcMap)
			t, e2 = t.Parse(string(b))
			if e2 != nil {
				return e2
			}
		}

		return nil
	})

	return root, err
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func routes(e *echo.Echo) {
	e.GET("/", handlers.Index)
	e.GET("/articles", handlers.Articles)

	e.Static("/styles", "static/styles")
	e.Static("/scripts", "static/scripts")
	e.Static("/vendor", "static/vendor")
	e.Static("/images", "static/images/")
}

func main() {
	db := storage.InitDB("storage.db")
	storage.Migrate(db)

	t := &Template{
		templates: template.Must(findAndParseTemplates("templates", nil)),
	}

	e := echo.New()
	e.Renderer = t
	routes(e)

	e.Logger.Fatal(e.Start(":5555"))
}
