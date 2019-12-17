package utils

import (
    "fmt"
    "os"
	"io/ioutil"
    "github.com/gomarkdown/markdown"
    "github.com/gomarkdown/markdown/parser"
)

func main() {
	if (len(os.Args) != 2) {
		fmt.Println("Missing argument: filename")
		return
	}

	filename := os.Args[1]

	bytes, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Println(err)
		return
   }

	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)
	html := markdown.ToHTML(bytes, parser, nil)

    fmt.Println(string(html))
}

