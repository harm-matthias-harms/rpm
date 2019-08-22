package storage

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"

	"github.com/harm-matthias-harms/rpm/backend/model"
	"go.mongodb.org/mongo-driver/mongo"
)

// FindUser will find a user on a subset of a complete user model
func FindUser(user *model.User) (result *model.User, err error) {
	c := userCollection()
	filter := bson.D{{Key: "$or", Value: bson.A{
		bson.D{{Key: "_id", Value: user.ID}},
		bson.D{{Key: "username", Value: user.Username}},
		bson.D{{Key: "email", Value: user.Username}},
	}}}
	result = new(model.User)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = c.FindOne(ctx, filter).Decode(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateUser will create a user with provided informations
func CreateUser(user *model.User) (err error) {
	c := userCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = c.InsertOne(ctx, user)
	return
}

func userCollection() *mongo.Collection {
	return MongoSession.Collection("user")
}

// CreateUserIndexes must be public for db reset for integration testing
func CreateUserIndexes() {
	c := userCollection()
	_, _ = c.Indexes().CreateMany(context.Background(),
		[]mongo.IndexModel{{
			Keys:    bsonx.Doc{{Key: "username", Value: bsonx.Int32(1)}},
			Options: options.Index().SetUnique(true),
		},
			{
				Keys:    bsonx.Doc{{Key: "email", Value: bsonx.Int32(1)}},
				Options: options.Index().SetUnique(true),
			},
		})
}
