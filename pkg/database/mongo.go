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
	Client *mongo.Client
	DB     *mongo.Database
}

var (
	MongoURL string
	Mg       MongoInstance
)

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

func GetCollection(name string) *mongo.Collection {
	return Mg.DB.Collection(name)
}

func StoreEdt(edt models.StudentEDT) {
	collection := GetCollection("edt")

	fmt.Printf("Storing %d edt\n", len(edt))

	for index, studentEDT := range edt {
		_, err := collection.UpdateOne(context.Background(), bson.M{"_id": index}, bson.D{{"$set", bson.M{"_id": index, "edt": studentEDT}}}, options.Update().SetUpsert(true))
		if err != nil {
			log.Printf("Error while storing student edt: %v", err)
			continue
		}
	}
}

func GetEdt(uuid string) bson.M {
	collection := GetCollection("edt")

	var result bson.M

	err := collection.FindOne(context.Background(), bson.M{"_id": uuid}).Decode(&result)
	if err != nil {
		log.Printf("Error while getting student edt: %v", err)
		return nil
	}

	return result
}
