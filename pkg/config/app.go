package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// the point of that file is to provide a variable called DB

var (
	db *gorm.DB
)

func Connect() {
	// to open the connection with the database
	d, err := gorm.Open(mysql.Open("root:Zahrou.001@tcp(127.0.0.1:3306)/store?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	// simplerest?charset=utf86parseTime=True&loc=Local
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
