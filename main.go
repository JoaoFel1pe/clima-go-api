package main

import (
	"api/src/config"
	"api/src/router"
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
