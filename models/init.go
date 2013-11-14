package models

import (
	"labix.org/v2/mgo"
)

var c_articles, c_users *mgo.Collection

func init() {
	connectionStr := "localhost"

	session, err := mgo.Dial(connectionStr)
	if err != nil {
		panic(err)
	}

	DB := session.DB("MartianBlog")
	c_articles = DB.C("articles")
	c_users = DB.C("users")
}
