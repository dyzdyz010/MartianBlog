package models

import (
	"labix.org/v2/mgo/bson"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func FindUser(user *User) *User {
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
