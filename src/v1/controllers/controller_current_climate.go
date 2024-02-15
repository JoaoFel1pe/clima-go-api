package controllers

import (
	"api/src/v1/helpers"
	"fmt"
	"net/http"
)

func GetCurrentClimate(w http.ResponseWriter, r *http.Request) {
	geolocation, err := helpers.GetLocation(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	query := helpers.BuildQueryCurrentCondition(fmt.Sprintf("%f", geolocation["lat"]), fmt.Sprintf("%f", geolocation["lng"]))

	result := helpers.ExecuteQuery(query)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(result))
}
