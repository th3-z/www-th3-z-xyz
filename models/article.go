package models

type Article struct{
	Filename string
}

var articlesPath = "/articles"

func GetArticles(path string) []Article {
	var articles []Article
	return articles
}

