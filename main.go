package main

import (
	"context"
	"fmt"
	"github.com/corentings/uca-edt/app/routes"
	"github.com/corentings/uca-edt/pkg/core"
	"github.com/corentings/uca-edt/pkg/database"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	MongoURL string
)

func LoadVar() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	MongoURL = os.Getenv("MONGO_URL") // Get url from env
}

func main() {

	// Load var from .env file
	LoadVar()

	err := database.Connect(MongoURL)
	if err != nil {
		log.Panic(err)
	}

	defer func() {
		fmt.Println("Disconnect")
		err := database.Mg.Client.Disconnect(context.TODO())
		if err != nil {
			return
		}
	}()

	core.ComputeStudentEDT()

	// Create the app
	app := routes.New()
	log.Panic(app.Listen(":3000"))

}
