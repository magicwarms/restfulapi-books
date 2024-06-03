package config

import (
	"encoding/json"
	"log"

	"github.com/joho/godotenv"
)

// PrettyPrint is make easier to print data result to console after querying data to db
func PrettyPrint(i interface{}) string {
	results, _ := json.MarshalIndent(i, "", "\t")
	return string(results)
}

func LoadEnvVariable() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("error loading .env file")
	}
}
