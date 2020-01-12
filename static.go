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
