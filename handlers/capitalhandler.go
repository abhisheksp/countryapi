package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/abhisheksp/countryapi/dependencyfactory"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
)

type CountryInfo struct {
	Country map[string]string
}

type CountryResponse struct {
	Country string `json:"country"`
	Capital string `json:"capital"`
	Error   bool   `json:"error"`
}

func Capital(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	country := strings.ToLower(vars["country"])
	countryResponse := CountryResponse{}
	capital, found := dependencyfactory.CountryCapitals[country]
	if !found {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("[INFO]", country, "Invalid Country Name")
		countryResponse = CountryResponse{Error: true}
	} else {
		countryResponse = CountryResponse{
			Country: country,
			Capital: capital,
		}
	}
	response, err := json.Marshal(countryResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("[FATAL]", err.Error())
		countryResponse = CountryResponse{Error: true}
	}
	fmt.Fprintf(w, string(response))
}
