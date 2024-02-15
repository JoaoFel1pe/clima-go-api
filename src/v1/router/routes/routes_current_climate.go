package routes

import (
	"api/src/v1/controllers"
	"net/http"
)

var routeCurrentClimate = []Rota{
	{
		URI:                "/current-climate/{city}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.GetCurrentClimate,
		RequerAutenticacao: false,
	},
}
