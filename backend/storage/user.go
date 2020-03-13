package storage

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"

	"github.com/harm-matthias-harms/rpm/backend/model"
	"go.mongodb.org/mongo-driver/mongo"
)

// FindUser will find a user on a subset of a complete user model
func FindUser(ctx context.Context, user *model.User) (result *model.User, err error) {
	c := userCollection()
	filter := bson.D{{Key: "code", Value: user.Code}}
	if user.Code == "" {
		filter = bson.D{{Key: "$or", Value: bson.A{
			bson.D{{Key: "_id", Value: user.ID}},
			bson.D{{Key: "username", Value: user.Username}},
			bson.D{{Key: "email", Value: user.Username}},
		}}}
	}
	result = new(model.User)
	err = c.FindOne(ctx, filter).Decode(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateUser will create a user
func CreateUser(ctx context.Context, user *model.User) (err error) {
	c := userCollection()
	_, err = c.InsertOne(ctx, user)
	return
}

func userCollection() *mongo.Collection {
	return MongoSession.Collection("users")
}

func createUserIndexes() {
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
			{
				Keys:    bsonx.Doc{{Key: "code", Value: bsonx.Int32(1)}},
				Options: options.Index().SetUnique(true),
			},
		})
}
