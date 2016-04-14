package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/abhisheksp/countryapi/dependencyfactory"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

type CountryInfo struct {
	Country map[string]string
}

type CountryResponse struct {
	Country string `json:"country"`
	Capital string `json:"capital"`
}

func Capital(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	country := strings.ToLower(vars["country"])
	capital, found := dependencyfactory.CountryCapitals[country]
	if !found {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Invalid Country Name")
		return
	}
	countryResponse := CountryResponse{
		Country: country,
		Capital: capital,
	}
	response, err := json.Marshal(countryResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
	}
	fmt.Fprintf(w, string(response))
}
