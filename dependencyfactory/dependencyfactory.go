package dependencyfactory

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

var once sync.Once
var countryCapitals map[string]string

func CountryCapitals() map[string]string {
	once.Do(func() {
		capitals, e := ioutil.ReadFile("./jsons/capitals.json")
		countryCapitals = make(map[string]string)

		json.Unmarshal(capitals, &countryCapitals)
		if e != nil {
			fmt.Printf("File error: %v\n", e)
			os.Exit(1)
		}
	})

	return countryCapitals
}
