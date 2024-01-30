package routes

import (
	"api/src/controllers"
	"net/http"
)

var routesClimaAtual = []Rota{
	{
		URI:                "/clima-atual",
		Metodo:             http.MethodGet,
		Funcao:             controllers.GetCurrentClimate,
		RequerAutenticacao: false,
	},
}
