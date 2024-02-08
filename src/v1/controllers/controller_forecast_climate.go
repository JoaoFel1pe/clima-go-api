package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetForecastClimate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	city, days := params["city"], params["days"]

	fmt.Println("Forecast Climate: ", city, days)
}
