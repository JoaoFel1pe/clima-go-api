package controllers

import (
	"net/http"
)

func GetCurrentClimate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ACTUAL CLIMATE"))
}
