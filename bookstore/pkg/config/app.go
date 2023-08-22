package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// _ "gorm.io/gorm" 
	// "gorm.io/driver/mysql" 
)

var (
	db * gorm.DB
)

func Connect(){
	d, err := gorm.Open("mysql", "ram:Ram@123456@tcp(127.0.0.1:3306)/go_test?charset=utf8&parseTime=True&loc=Local")
	// d, err := gorm.Open(mysql.Open("root:admin@tcp(127.0.0.1:3306)/yourdatabase?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}