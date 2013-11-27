package models

import (
	"fmt"
	"labix.org/v2/mgo/bson"
)

type Article struct {
	Id      bson.ObjectId `json:"id" bson:"_id"`
	Title   string        `json:"title"`
	Author  string        `json:"author"`
	Date    string        `json:"date"`
	Tags    string        `json:"tags"`
	Content string        `json:"content"`
	Status  string        `json:"status"`
}

func AllArticles() []Article {
	articles := []Article{}
	err := c_articles.Find(bson.M{"_id": bson.ObjectIdHex("5281b83afbb7f35cb62d0834")}).All(&articles)
	if err != nil {
		panic(err)
	}
	// fmt.Println(articles)
	return articles
}

func ArticlesByStatus(status string) []Article {
	articles := []Article{}
	err := c_articles.Find(bson.M{"status": status}).All(&articles)
	if err != nil {
		panic(err)
	}
	fmt.Println("Hex: ", articles[0].Id.Hex())

	return articles
}

func ArticleById(id string) *Article {
	article := &Article{}
	err := c_articles.FindId(bson.ObjectIdHex(id)).One(article)
	if err != nil {
		panic(err)
	}

	return article
}

func ArticleByTitle(title string) *Article {
	article := &Article{}
	err := c_articles.Find(bson.M{"title": title}).One(article)
	if err != nil {
		panic(err)
	}

	return article
}

func AddArticle(article Article) string {
	err := c_articles.Insert(article)
	if err != nil {
		panic(err)
	}

	a := ArticleByTitle(article.Title)

	return a.Id.Hex()
}
