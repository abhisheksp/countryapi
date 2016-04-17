package countryapi

import (
	"github.com/abhisheksp/countryapi/dependencyfactory"
	"github.com/abhisheksp/countryapi/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

const GET = "GET"

func Run(bindAddress string) {
	dependencyfactory.InitCountryCapitals()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/capital/{country}", handlers.Capital).Methods(GET)
	log.Fatal(http.ListenAndServe(bindAddress, router))
}
