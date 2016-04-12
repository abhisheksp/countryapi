package main

import (
	"github.com/abhisheksp/countryapi/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

const GET = "GET"

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/capital/{country}", handlers.Capital).Methods(GET)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
