package models

import (
	"bufio"
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
	IconFilename   string
	SourceFilename string
	OutputFilename string
}

type meta struct {
	Date        string
	Title       string
	Description string
	Content     string
	Icon        string
}

const srcPath string = "templates/articles/_src"
const outPath string = "templates/articles"
const baseTemplatePath string = "templates/articles/base.html"

func readMeta(filename string) *meta {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var meta meta
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
		}
	}

	return &meta
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
		article.IconFilename = "/images/articles/" + f.Name() + "/" + meta.Icon
		article.SourceFilename = srcPath + "/" + article.Name + "/" + meta.Content
		article.OutputFilename = outPath + "/" + article.Name + "/" + strings.Split(meta.Content, ".")[0] + ".html"

		articles = append(articles, article)

		article.Bake()
	}

	if err != nil {
		panic(err)
	}

	return articles
}

func (article Article) Content() *string {
	content, err := ioutil.ReadFile(article.OutputFilename)
	if err != nil {
		panic(err)
	}

	contentStr := string(content)

	return &contentStr
}

func (article Article) Bake() {
	aSrcPath := srcPath + "/" + article.Name
	aOutPath := outPath + "/" + article.Name
	os.Mkdir(aOutPath, 0755)

	srcFiles, err := ioutil.ReadDir(aSrcPath)
	if err != nil {
		panic(err)
	}

	for _, f := range srcFiles {
		if f.IsDir() {
			continue
		}
		utils.Copy(aSrcPath+"/"+f.Name(), aOutPath+"/"+f.Name())
	}

	html := utils.MdToHtml(article.SourceFilename)

	outputFile, err := os.Create(article.OutputFilename)
	defer outputFile.Close()

	outputFile.Write(*html)
	writer := bufio.NewWriter(outputFile)
	writer.Write(*html)
	writer.Flush()
}
