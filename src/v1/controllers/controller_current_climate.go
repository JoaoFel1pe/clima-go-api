package controllers

import (
	"api/src/v1/helpers"
	"fmt"
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

	geolocation, err2 := helpers.GetGeolocation(city)
	if err2 != nil {
		helpers.ResponseError(w, http.StatusBadRequest, err2)
		return
	}

	query := helpers.BuildQueryCurrentCondition(fmt.Sprintf("%f", geolocation["lat"]), fmt.Sprintf("%f", geolocation["lng"]))
	fmt.Println(query)
	w.Write([]byte(fmt.Sprintf(query)))

}
