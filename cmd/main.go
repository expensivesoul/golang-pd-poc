package main

import (
	"log"
	"net/http"
	
	"pkg/route"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
func main() {
	r := mux.NewRouter()
	route.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8082", r))
}

