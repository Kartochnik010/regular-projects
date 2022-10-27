package main

import (
	"book_store/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
