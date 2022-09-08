package utils

import (
	"github.com/corentings/uca-edt/pkg/database"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadVar() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.MongoURL = os.Getenv("MONGO_URL") // Get url from env
}
