package models

import (
	"labix.org/v2/mgo/bson"
)

type Article struct {
	Id      int64
	Title   string
	Author  string
	Date    string
	Tags    string
	Content string
	Status  string
}

func AllArticles() []Article {
	articles := []Article{}
	err := c_articles.Find(bson.M{}).All(&articles)
	if err != nil {
		panic(err)
	}

	return articles
}

func ArticlesByStatus(status string) []Article {
	articles := []Article{}
	err := c_articles.Find(bson.M{"status": status}).All(&articles)
	if err != nil {
		panic(err)
	}

	return articles
}

func ArticleById(id int64) Article {
	article := Article{}
	err := c_articles.Find(bson.M{"id": id}).One(&article)
	if err != nil {
		panic(err)
	}

	return article
}

func ArticleByTitle(title string) Article {
	article := Article{}
	err := c_articles.Find(bson.M{"title": title}).One(&article)
	if err != nil {
		panic(err)
	}

	return article
}
