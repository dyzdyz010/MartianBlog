package models

import (
	"labix.org/v2/mgo/bson"
)

type User struct {
	username string
	email    string
	password string
}

func GetUserByEmail(email string) *User {
	u := &User{}
	err := c_users.Find(bson.M{"email": email}).One(u)
	if err != nil {
		panic(err)
	}

	return u
}
