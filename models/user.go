package models

import (
	"crypto/md5"
	"fmt"
	"labix.org/v2/mgo/bson"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func AddUser(user *User) {
	fmt.Println(user)
	h := md5.New()
	h.Write([]byte(user.Password))
	user.Password = fmt.Sprintf("%x", h.Sum(nil))

	err := c_users.Insert(bson.M{"email": user.Email, "username": user.Username, "password": user.Password})
	if err != nil {
		panic(err)
	}
}

func FindUser(user *User) *User {
	fmt.Println(user)
	h := md5.New()
	h.Write([]byte(user.Password))
	user.Password = fmt.Sprintf("%x", h.Sum(nil))

	u := &User{}
	err := c_users.Find(bson.M{"email": user.Email, "password": user.Password}).One(u)
	if err != nil {
		panic(err)
	}

	return u
}

func GetUserByEmail(email string) *User {
	u := &User{}
	err := c_users.Find(bson.M{"email": email}).One(u)
	if err != nil {
		panic(err)
	}

	return u
}
