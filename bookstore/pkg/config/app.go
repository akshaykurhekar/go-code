package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db * gorm.DB
)

func connect(){
	d, err := gorm.Open("mysql", "username:password/dbName?charset=utf8&parseTime=True&loc=Local")
	err != nil {
		panic(err)
	}
	db = d
}

func getDB() *gorm.DB {
	return db
}