package helper

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GoDotEnvVariable(_connectionString string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error in loading .env file")
	}

	return os.Getenv(_connectionString)
}
