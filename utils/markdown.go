package utils

import (
	"io/ioutil"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

func MdToHtml(filename string) *[]byte {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)
	output := markdown.ToHTML(bytes, parser, nil)

	return &output
}
