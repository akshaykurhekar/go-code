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

func GetBook ( w http.ResponseWriter,r *http.Request){
	newBooks := models.GetAllBook()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)

	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}	
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook( w http.ResponseWriter,r *http.Request){
	vars := mux.Vars(r)

	bookId := vars["bookId"]

	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("parsing error")
	}
	deleteBookDetails := models.DeleteBook(ID)
	res, _ := json.Marshal(deleteBookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook( w http.ResponseWriter, r *http.Request){
	var NewBook = &models.Book{}
	utils.ParseBody(r, NewBook)
	
	b := NewBook.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	var newBook = &models.Book{}
	utils.ParseBody(r, newBook)

	prams := mux.Vars(r)
	bookId := prams["bookId"]
	ID, err := strconv.ParseInt(bookId,0,0)
	if err != nil {
		fmt.Println("error while parsing...")
	}

	bookDetails, db := models.GetBookById(ID)
	
	if newBook.Name != ""{
		bookDetails.Name = newBook.Name
	} 
	if newBook.Author != ""{
		bookDetails.Author = newBook.Author
	}
	if newBook.Publication != ""{
		bookDetails.Publication = newBook.Publication
	}

	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	// w.Header().Set.("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}