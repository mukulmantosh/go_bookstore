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
	log.Println("Listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
