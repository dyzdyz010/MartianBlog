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
	err := c_articles.Find(bson.M{}).All(&articles)
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
	article.Id = bson.NewObjectId()
	// err := c_articles.Insert(article)
	// if err != nil {
	// 	panic(err)
	// }

	return article.Id.Hex()
}

func UpdateArticle(article Article) bool {
	err := c_articles.UpdateId(article.Id, article)
	if err != nil {
		panic(err)
		return false
	}

	return true
}

func DeleteArticle(idStr string) bool {
	// id := bson.ObjectIdHex(idStr)
	// err := c_articles.RemoveId(id)
	// if err != nil {
	// 	panic(err)
	// 	return false
	// }

	return true
}
