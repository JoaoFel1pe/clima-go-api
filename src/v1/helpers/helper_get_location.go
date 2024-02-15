package helpers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetLocation(r *http.Request) (map[string]float64, error) {
	params := mux.Vars(r)
	city, exists := params["city"]
	if !exists {
		return nil, fmt.Errorf("city not found")
	}

	geolocation, err := GetGeolocation(city)
	if err != nil {
		return nil, fmt.Errorf("city not found")
	}

	return geolocation, nil
}
