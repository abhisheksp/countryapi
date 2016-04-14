package dependencyfactory

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var CountryCapitals map[string]string

func InitCountryCapitals() {
	capitals, e := ioutil.ReadFile("./jsons/capitals.json")
	CountryCapitals = make(map[string]string)
	json.Unmarshal(capitals, &CountryCapitals)

	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
}
