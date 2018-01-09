package main

import (
	"github.com/damianopetrungaro/golang-bookshop/book"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	book.LoadRoutes("/books", router)
	http.ListenAndServe(":80", router)
}
