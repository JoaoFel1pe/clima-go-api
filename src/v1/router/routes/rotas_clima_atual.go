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

var routeForecastClimate = []Rota{
	{
		URI:                "/forecast-climate/{city}/{days}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.GetForecastClimate,
		RequerAutenticacao: false,
	},
}
