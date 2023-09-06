package utils

import (
	"log"
	"os"

	"github.com/corentings/uca-edt/app/middleware"
	"github.com/corentings/uca-edt/pkg/database"
	"github.com/joho/godotenv"
)

// LoadVar loads the var from the .env file
func LoadVar() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.MongoURL = os.Getenv("MONGO_URL") // Get url from env
	middleware.SecurityKey = os.Getenv("SECURITY_KEY")
}
