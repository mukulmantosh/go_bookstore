package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go_bookstore/pkg/models"
	"go_bookstore/pkg/utils"
	"net/http"
	"strconv"
)

var NewBook models.Book

func GetBook(writer http.ResponseWriter, request *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)

}

func GetBookById(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Parse Error: ", err.Error())
	}

	bookDetails, _ := models.GetBookById(id)
	res, _ := json.Marshal(bookDetails)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
}

func CreateBook(writer http.ResponseWriter, request *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(request, CreateBook)
	book := CreateBook.CreateBook()
	res, _ := json.Marshal(book)
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
}

func DeleteBook(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Parse Error: ", err.Error())
	}

	book := models.DeleteBook(id)
	result, _ := json.Marshal(book)

	writer.WriteHeader(http.StatusNoContent)
	writer.Write(result)

}

func UpdateBook(writer http.ResponseWriter, request *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(request, updateBook)
	vars := mux.Vars(request)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Parse Error: ", err.Error())
	}

	bookDetails, db := models.GetBookById(id)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}

	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}

	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

}
