package database

import (
	"context"
	"fmt"
	"github.com/corentings/uca-edt/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// MongoInstance contains the Mongo client and database objects
type MongoInstance struct {
	Client *mongo.Client   // Mongo client
	DB     *mongo.Database // Mongo database
}

var (
	MongoURL string        // MongoURL is the url of the mongo database
	Mg       MongoInstance // Mg is the mongo instance
)

// Connect connects to the database
func Connect(mongoURL string) error { // Set client options
	clientOptions := options.Client().ApplyURI(mongoURL)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	// Connect to MongoDB
	client, e := mongo.Connect(ctx, clientOptions)
	if e != nil {
		return e
	}

	// Check the connection
	e = client.Ping(ctx, nil)
	if e != nil {
		return e
	}

	fmt.Println("Connected to mongoDB !")
	// get collection as ref
	db := client.Database("edt")

	Mg = MongoInstance{Client: client, DB: db}

	return nil
}

// GetCollection returns a collection from the database with the given name
func GetCollection(name string) *mongo.Collection {
	return Mg.DB.Collection(name) // get collection as ref
}

// StoreEdt stores the edt in the database
func StoreEdt(edt models.StudentEDT) {
	collection := GetCollection("edt") // get collection as ref

	for index, studentEDT := range edt {
		_, err := collection.UpdateOne(context.Background(), bson.M{"_id": index}, bson.D{{"$set", bson.M{"_id": index, "edt": studentEDT}}}, options.Update().SetUpsert(true)) // update or insert
		if err != nil {
			log.Printf("Error while storing student edt: %v", err)
			continue
		}
	}
}

// GetEdt returns the edt of a student from the database
func GetEdt(uuid string) (bson.M, error) {
	collection := GetCollection("edt") // get collection as ref

	var result bson.M // result

	err := collection.FindOne(context.Background(), bson.M{"_id": uuid}).Decode(&result) // find student edt
	if err != nil {
		log.Printf("Error while getting student edt: %v", err)
		return nil, err
	}

	return result, nil
}
