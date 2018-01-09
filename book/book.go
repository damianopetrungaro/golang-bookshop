package book

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func LoadRoutes(prefix string, router *mux.Router) {
	router.HandleFunc(prefix, index)
	router.HandleFunc(prefix+ "/{id}", show)
}
func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Book index")
}

func show(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Book show")
}
