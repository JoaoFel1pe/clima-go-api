package main

import (
	"api/src/v1/config"
	"api/src/v1/middlewares"
	"api/src/v1/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadEnvs()
	r := router.Generate()

	r.Use(middlewares.LogginMiddleware)
	fmt.Printf("Server running on http://localhost:%d\n", config.APPLICATION_PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", config.APPLICATION_PORT), r))
}
