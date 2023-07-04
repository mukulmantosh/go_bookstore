package main

import (
	"github.com/gorilla/mux"
	"go_bookstore/pkg/routes"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8000", router))
}
