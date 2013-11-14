package models

import (
	"fmt"
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
	err := c_articles.Find(bson.M{"id": 226}).All(&articles)
	if err != nil {
		panic(err)
	}
	fmt.Println(articles)
	// c.Insert(bson.M{"id": 226, "title": "Hello2"})

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
