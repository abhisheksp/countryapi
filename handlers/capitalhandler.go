package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"os"
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
	capitals, e := ioutil.ReadFile("./jsons/capitals.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	countryInfo := make(map[string]string)
	json.Unmarshal(capitals, &countryInfo)
	vars := mux.Vars(r)
	country := strings.ToLower(vars["country"])
	capital, found := countryInfo[country]
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
