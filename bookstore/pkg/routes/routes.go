package routes

import (
	"github.com/gorilla/mux"
	"github.com/akshay/bookstore/pkg/controller"
)

var registerBookStoreRoute = func(routes *mux.Router){
	router.HandleFunc("/book/",controller.createBook).Method("POST")
	router.HandleFunc("/book/",controller.getBook).Method("GET")
	router.HandleFunc("/book/{bookId}",controller.getBookById).Method("GET")
	router.HandleFunc("/book/{bookId}",controller.updateBook).Method("PUT")
	router.HandleFunc("/book/{bookId}",controller.deleteBook).Method("DELETE")
} 