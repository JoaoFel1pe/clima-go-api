package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Rota representa todas as routes da API
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

func Configurar(r *mux.Router) *mux.Router {
	routes := routesClimaAtual
	for _, rota := range routes {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}
	return r
}
