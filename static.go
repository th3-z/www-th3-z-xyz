package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

var mdOutputPath = "templates/"
var mdOutputExt = ".html"

func FindAndRenderGraphs(rootDir string) {
	/*
	 * Render DOT files from /static/graphs to /static/images/graphs
	 */
}

func FindAndRenderMarkdown(rootDir string) {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)

	cleanRoot := filepath.Clean(rootDir)
	pfx := len(cleanRoot) + 1

	err := filepath.Walk(cleanRoot, func(path string, info os.FileInfo, e1 error) error {
		if !info.IsDir() && strings.HasSuffix(path, ".md") {
			if e1 != nil {
				return e1
			}

			b, e2 := ioutil.ReadFile(path)
			if e2 != nil {
				return e2
			}

			name := path[pfx:]

			html := markdown.ToHTML(b, parser, nil)
			filename := mdOutputPath + strings.TrimSuffix(name, filepath.Ext(name)) + mdOutputExt
			e3 := ioutil.WriteFile(filename, html, 0644)
			if e3 != nil {
				return e3
			}
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

}

func main() {
	FindAndRenderMarkdown("static/markdown")
}