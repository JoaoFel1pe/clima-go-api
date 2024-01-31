package router

import (
	"api/src/v1/config"
	"api/src/v1/router/routes"

	"github.com/gorilla/mux"
)

func Generate() *mux.Router {
	r := mux.NewRouter()
	api_v1 := r.PathPrefix(config.API_V1_STR).Subrouter()
	return routes.Configurar(api_v1)
}
