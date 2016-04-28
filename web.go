package main

import (
	"fmt"
	"github.com/abhisheksp/countryapi/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

const GET = "GET"

func main() {
	bind := fmt.Sprintf("%s:%s", os.Getenv("OPENSHIFT_GO_IP"), os.Getenv("OPENSHIFT_GO_PORT"))
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/capital/{country}", handlers.Capital).Methods(GET)
	log.Fatal(http.ListenAndServe(bind, router))
}
