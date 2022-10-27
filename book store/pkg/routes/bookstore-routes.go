package routes

import (
	"book_store/pkg/controllers"

	"github.com/gorilla/mux"
)

var (
	RegisterBookStoreRoutes = func(r *mux.Router) {
		r.HandleFunc("/book", controllers.GetAllBooks()).Methods("GET")        // get all books
		r.HandleFunc("/book/{id}", controllers.GetBookById()).Methods("GET")   // get book by id
		r.HandleFunc("/book", controllers.CreateBook()).Methods("POST")        // create a book
		r.HandleFunc("/book/{id}", controllers.UpdateBook()).Methods("PUT")    // update a book by id
		r.HandleFunc("/book/{id}", controllers.DeleteBook()).Methods("DELETE") // delete a book by id
	}
)
