package routes

import (
	"github.com/ZahreddineSouale/pkg/controller"
	"github.com/gorilla/mux"
)

var RegisterBookStore = func(router *mux.Router) {
	router.HandleFunc("/", controller.ServeHome).Methods("GET")
	router.HandleFunc("/book", controller.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controller.GetBoookById).Methods("POST")
	router.HandleFunc("/book", controller.CreateBooks).Methods("POST")
	router.HandleFunc("/book/{bookId}", controller.DeleteBook).Methods("DELETE")
}
