package main

import (
	"github.com/abhisheksp/countryapi/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handlers.Captial)
	log.Fatal(http.ListenAndServe(":8080", router))
}
