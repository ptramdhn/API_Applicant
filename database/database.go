package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/hackfest"))
	if err != nil {
		panic(err)
	}

	return database
}
