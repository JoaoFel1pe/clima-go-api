package main

import (
	"api/src/v1/config"
	"api/src/v1/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadEnvs()
	r := router.Generate()

	fmt.Printf("Escutando na porta %d\n", config.APPLICATION_PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.APPLICATION_PORT), r))
}
