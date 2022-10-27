package controllers

import (
	"book_store/pkg/models"
	"book_store/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var (
	newBook models.Book
)

func GetAllBooks() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		newBooks := models.GetAllBooks()
		res, _ := json.Marshal(newBooks)
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func GetBookById() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)

		bookId, err := strconv.ParseInt(params["id"], 0, 0)
		if err != nil {
			fmt.Println(err)
		}
		neededBook, _ := models.GetBookById(bookId)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		res, _ := json.Marshal(neededBook)
		w.Write(res)

	}
}
func CreateBook() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		CreateBook := &models.Book{}
		utils.ParseBody(r, CreateBook)
		b := CreateBook.CreateBook()
		res, _ := json.Marshal(b)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func UpdateBook() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var updateBook = &models.Book{}
		utils.ParseBody(r, updateBook)
		params := mux.Vars(r)
		bookId := params["id"]
		ID, err := strconv.ParseInt(bookId, 0, 0)
		if err != nil {
			fmt.Println("error while parsing")
		}
		bookDetails, db := models.GetBookById(ID)
		if updateBook.Name != "" {
			bookDetails.Name = updateBook.Name
		}
		if updateBook.Author != "" {
			bookDetails.Author = updateBook.Author
		}
		if updateBook.Publication != "" {
			bookDetails.Publication = updateBook.Publication
		}
		db.Save(&bookDetails)
		res, _ := json.Marshal(bookDetails)
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}
func DeleteBook() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		bookId := params["id"]
		ID, err := strconv.ParseInt(bookId, 0, 0)
		if err != nil {
			fmt.Println("error while parsing")
		}
		book := models.DeleteBook(ID)
		if book.ID != 0 {
			w.Write([]byte("Something went wrong"))
			return
		}
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusOK)
		res, _ := json.Marshal(models.GetAllBooks())
		w.Write(res)
	}
}
