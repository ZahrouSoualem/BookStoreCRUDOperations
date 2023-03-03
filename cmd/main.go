package main

import (
	"fmt"
	"log"
	"net/http"

	/* "gorm.io/driver/mysql"
	"gorm.io/gorm" */
	"github.com/ZahreddineSouale/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Soualem Zahreddine")

	r := mux.NewRouter()
	routes.RegisterBookStore(r)
	//http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":4000", r))

}
