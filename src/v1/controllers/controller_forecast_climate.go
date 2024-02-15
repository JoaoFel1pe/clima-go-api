package controllers

import (
	"api/src/v1/helpers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetForecastClimate(w http.ResponseWriter, r *http.Request) {
	days := mux.Vars(r)["days"]
	geolocation, err := helpers.GetLocation(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	query := helpers.BuildQueryForecast(fmt.Sprintf("%f", geolocation["lat"]), fmt.Sprintf("%f", geolocation["lng"]), days)

	result := helpers.ExecuteQuery(query)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(result))
}
