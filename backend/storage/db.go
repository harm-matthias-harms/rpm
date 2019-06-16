package storage

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/harm-matthias-harms/rpm/backend/utils"
)

var mongoSession *mongo.Database

func setMongoDatabase() {
	client, err := mongo.NewClient(options.Client().ApplyURI(utils.GetEnv("MONGO_URL", "mongodb://localhost:27017")))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	mongoSession = client.Database(utils.GetEnv("MONGO_DATABASE", "test"))
}
