package models

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego/config"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

var c_articles, c_users *mgo.Collection

var BlogInfo map[string]string
var admin *User

func init() {
	connectionStr := "localhost"

	session, err := mgo.Dial(connectionStr)
	if err != nil {
		panic(err)
	}

	DB := session.DB("MartianBlog")
	c_articles = DB.C("articles")
	c_users = DB.C("users")

	readConfig()
}

func readConfig() {
	jsonconf, err := config.NewConfig("json", "conf/config.json")
	if err != nil {
		panic(err)
	}

	BlogInfo = make(map[string]string)
	BlogInfo["title"] = jsonconf.String("blog::title")
	BlogInfo["subtitle"] = jsonconf.String("blog::subtitle")
	fmt.Println(BlogInfo)

	admin = new(User)
	admin.Email = jsonconf.String("admin::email")
	admin.Username = jsonconf.String("admin::username")
	admin.Password = jsonconf.String("admin::password")
	_, err = c_users.RemoveAll(bson.M{})
	if err != nil {
		panic(err)
	}
	AddUser(admin)

	h := md5.New()
	h.Write([]byte(admin.Email))
	BlogInfo["avatarhash"] = fmt.Sprintf("%x", h.Sum(nil))
}
