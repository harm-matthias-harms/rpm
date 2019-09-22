package storage

import (
	"context"
	"time"

	"github.com/harm-matthias-harms/rpm/backend/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoSession Handels the mongo Database needs to be public for integration testing.
var MongoSession *mongo.Database

// SetMongoDatabase sets the database when ever a server is created
func SetMongoDatabase() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(utils.GetEnv("MONGO_URL", "mongodb://localhost:27017")))
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return err
	}
	MongoSession = client.Database(utils.GetEnv("MONGO_DATABASE", "test"))

	//create indexes
	createUserIndexes()
	return nil
}
