package helpers

import (
	"io"
	"log"
	"net/http"
)

func ExecuteQuery(query string) string {
	resp, err := http.Get(query)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}
