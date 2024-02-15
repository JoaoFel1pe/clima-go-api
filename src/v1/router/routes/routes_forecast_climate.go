package routes

import (
	"api/src/v1/controllers"
	"net/http"
)

var routeForecastClimate = []Rota{
	{
		URI:                "/forecast-climate/{city}/{days}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.GetForecastClimate,
		RequerAutenticacao: false,
	},
}
