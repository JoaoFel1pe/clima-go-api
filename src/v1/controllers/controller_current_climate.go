package controllers

import (
	"api/src/v1/helpers"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetCurrentClimate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	city, exists := params["city"]
	if !exists {
		helpers.ResponseError(w, http.StatusBadRequest, fmt.Errorf("city not found"))
		return
	}

	geolocation, err := helpers.GetGeolocation(city)
	if err != nil {
		helpers.ResponseError(w, http.StatusNotFound, fmt.Errorf("City not found"))
		return
	}

	query := helpers.BuildQueryCurrentCondition(fmt.Sprintf("%f", geolocation["lat"]), fmt.Sprintf("%f", geolocation["lng"]))

	result := executeQuery(query)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(result))
}

func executeQuery(query string) string {
	resp, err := http.Get(query)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}
