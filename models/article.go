package models

import (
	"bufio"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
	"time"
	"www-th3-z-xyz/utils"
)

type Article struct {
	Name           string
	Date           time.Time
	Title          string
	Description    string
	IconUrl        string
	SourceFilename string
	OutputFilename string
	Visible        bool
	Meta           *meta
}

type meta struct {
	Filename    string
	Date        string
	Title       string
	Description string
	Content     string
	Icon        string
	Visible     string
}

const srcPath string = "templates/articles/_src"
const outPath string = "templates/articles/_out"
const baseTemplatePath string = "templates/articles/base.html"

func readMeta(filename string) *meta {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var meta meta
	meta.Filename = filename
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ":")
		if len(parts) < 2 {
			continue
		}

		switch parts[0] {
		case "date":
			meta.Date = strings.TrimSpace(parts[1])
		case "title":
			meta.Title = strings.TrimSpace(parts[1])
		case "description":
			meta.Description = strings.TrimSpace(parts[1])
		case "content":
			meta.Content = strings.TrimSpace(parts[1])
		case "icon":
			meta.Icon = strings.TrimSpace(parts[1])
		case "visible":
			meta.Visible = strings.TrimSpace(parts[1])
		}
	}

	return &meta
}

func (meta meta) Write() {
	file, err := os.OpenFile(meta.Filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.WriteString("date: " + meta.Date + "\n")
	file.WriteString("title: " + meta.Title + "\n")
	file.WriteString("description: " + meta.Description + "\n")
	file.WriteString("content: " + meta.Content + "\n")
	file.WriteString("icon: " + meta.Icon + "\n")
	file.WriteString("visible: " + meta.Visible + "\n")
}

func GetArticles(path string) []Article {
	var articles []Article

	files, err := ioutil.ReadDir(srcPath)
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		meta := readMeta(srcPath + "/" + f.Name() + "/meta")

		var article Article
		article.Name = f.Name()
		date, err := time.Parse("2006-01-02", meta.Date)
		if err != nil {
			panic(err)
		}
		article.Date = date
		article.Title = meta.Title
		article.Description = meta.Description
		article.IconUrl = "/articles/res/" + f.Name() + "/" + meta.Icon
		article.SourceFilename = srcPath + "/" + article.Name + "/" + meta.Content
		article.OutputFilename = outPath + "/" + article.Name + ".html"
		article.Visible = meta.Visible != "0"
		article.Meta = meta

		articles = append(articles, article)

		article.Bake()
	}

	if err != nil {
		panic(err)
	}

	return articles
}

func GetArticle(name string) *Article {
	var article Article

	meta := readMeta(srcPath + "/" + name + "/meta")

	date, err := time.Parse("2006-01-02", meta.Date)
	if err != nil {
		panic(err)
	}

	article.Name = name
	article.Date = date
	article.Title = meta.Title
	article.Description = meta.Description
	article.IconUrl = "/articles/res/" + article.Name + "/" + meta.Icon
	article.SourceFilename = srcPath + "/" + article.Name + "/" + meta.Content
	article.OutputFilename = outPath + "/" + article.Name + ".html"
	article.Visible = meta.Visible != "0"
	article.Meta = meta

	return &article
}

func (article Article) Content() *template.HTML {
	content, err := ioutil.ReadFile(article.OutputFilename)
	if err != nil {
		panic(err)
	}

	contentStr := template.HTML(string(content))

	return &contentStr
}

func (article Article) Bake() {
	aSrcPath := srcPath + "/" + article.Name
	aResPath := "static/articles/" + article.Name
	os.Mkdir(aResPath, 0755)

	srcFiles, err := ioutil.ReadDir(aSrcPath)
	if err != nil {
		panic(err)
	}

	// Copy static resources
	for _, f := range srcFiles {
		if f.IsDir() {
			continue
		}

		if !(strings.HasSuffix(f.Name(), ".md") || f.Name() == "meta") {
			utils.Copy(aSrcPath+"/"+f.Name(), aResPath+"/"+f.Name())
		}
	}

	html := utils.MdToHtml(article.SourceFilename)
	outputFile, err := os.Create(article.OutputFilename)
	defer outputFile.Close()

	outputFile.Write(*html)
}
