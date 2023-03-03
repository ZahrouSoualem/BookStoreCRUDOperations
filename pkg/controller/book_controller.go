package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	model "github.com/ZahreddineSouale/pkg/model"
	"github.com/ZahreddineSouale/pkg/utils"
	"github.com/gorilla/mux"
)

var Newbook model.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	Newbook := model.GetAllBooks()
	res, _ := json.Marshal(Newbook)
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API</h1>"))
}
func GetBoookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		panic(err)
	}

	bookDetails, _ := model.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "Aplication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CreateBooks(w http.ResponseWriter, r *http.Request) {
	CreateBook := &model.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "Aplication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		panic(err)
	}

	book := model.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "Aplication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
