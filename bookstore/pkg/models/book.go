package models

import (
	"github.com/jinzhu/gorm"
	"github.com/akshay/bookstore/pkg/config"
)

var db *gorm.DB

// syntax to define struct in data base with this required field
type Book struct {
	gorm.Model
	Name string `gorm:""json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

// init Database in mysql
func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

//create book function in db
func (b *Book) CreateBook() *Book {
	// this will create in entry in db using gorm
	db.NewRecord(b) 
	db.Create(&b)
	return b
}

func GetAllBook() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?",Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(Id int64) Book{
	var book Book
	db.Where("ID=?",Id).Delete(book)
	return book
}