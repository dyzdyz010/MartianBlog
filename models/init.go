package models

import (
	"github.com/astaxie/beego/config"
	"labix.org/v2/mgo"
)

var Session *mgo.Session

func init() {
	userconf, err := config.NewConfig("json", "user.json")
	var connectionStr string
	if err != nil {
		panic("No user configuration file found.")
	}

	temp, err := userconf.DIY("database")
	if err != nil {
		panic(err)
	}

	dbconf, ok := temp.(config.ConfigContainer)
	if !ok {
		panic("Configuration file error, can't read database settings.")
	}

	connectionStr = dbconf.String("host")
	// Check user settings
	if dbconf.String("username") != "" {
		connectionStr = dbconf.String("username") + ":" + dbconf.String("password") + connectionStr
	}
	// Check port
	if dbconf.String("port") != "" {
		connectionStr = connectionStr + ":" + dbconf.String("port")
	}
	// Check database name
	if dbconf.String("database") != "" {
		connectionStr = connectionStr + "/" + dbconf.String("database")
	}

	// Add head
	connectionStr = "mongodb://" + connectionStr

	// Session, err = mgo.Dial("mongodb://dyzdyz010:DYZcrh2582525775@ds037488.mongolab.com:37488/moonlightter")
	// Session, err = mgo.Dial("localhost")
	Session, err = mgo.Dial(connectionStr)
	if err != nil {
		panic(err)
	}

}
