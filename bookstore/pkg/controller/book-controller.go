package controller

import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/akshay/bookstore/pkg/models"
	"github.com/akshay/bookstore/pkg/utils"
)

var NewBook models.Book

func getBook ( w http.ResponseWriter,r *http.Request){
	newBooks := models.getAllBook()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func getBookById (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)

	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}	
	bookDetails, _ := models.getBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func deleteBook( w http.ResponseWriter,r *http.Request){
	prams := mux.Vars(r)

	bookId := params["bookId"]

	ID, err := strconv.ParseInt(bokId, 0, 0)
	if err != nil {
		fmt.Println("parsing error")
	}
	deleteBookDetails := models.deleteBook(ID)
	res, _ := json.Marshal(deleteBookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func createBook( w http.ResponseWriter, r *http.Request){
	prams := mux.Vars(r)

	NewBook = params.body

	newBook := models.createBook(NewBook)
	res := json.Marshal(newBook)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}