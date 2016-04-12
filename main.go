package main

import (
	"github.com/abhisheksp/countryapi/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const GET = "GET"

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/capital/{country}", handlers.Capital).Methods(GET)
	log.Fatal(http.ListenAndServe(":8080", router))
}
