package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type (
	ConnectionOptions struct {
		DatabaseName   string
		User           string
		Password       string
		Port           string
		Host           string
		DatabaseParams string
	}
)

func GetConnection(options ConnectionOptions) *gorm.DB {
	connString := options.User + ":" + options.Password + "@tcp(" + options.Host + ":" + options.Port + ")/" + options.DatabaseName + "?" + options.DatabaseParams
	db, err := gorm.Open("mysql", connString)

	if err != nil {
		log.Println("An error occurred during opening db connection", err)
		return nil
	}

	return db
}
