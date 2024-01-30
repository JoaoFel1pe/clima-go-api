package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	API_V1_STR = ""

	AZURE_API_KEY = ""

	APPLICATION_PORT = 0

	AZURE_BASE_URL = ""
)

func LoadEnvs() {

	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	APPLICATION_PORT, erro = strconv.Atoi(os.Getenv("APPLICATION_PORT"))
	if erro != nil {
		APPLICATION_PORT = 9000
	}

	API_V1_STR = os.Getenv("API_V1_STR")
	AZURE_API_KEY = os.Getenv("AZURE_API_KEY")
	AZURE_BASE_URL = os.Getenv("AZURE_BASE_URL")

}
