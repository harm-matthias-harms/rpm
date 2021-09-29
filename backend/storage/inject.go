package storage

import (
	"context"
	"errors"

	"github.com/harm-matthias-harms/rpm/backend/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CountInjects returns the number of total injects in the database
func CountInjects(ctx context.Context, params map[string]interface{}) (int64, error) {
	if len(params) == 0 {
		return injectCollection().CountDocuments(ctx, bson.M{})
	}
	return injectCollection().CountDocuments(ctx, params)
}

// GetInjects returns an array of injects
func GetInjects(ctx context.Context, params map[string]interface{}, page int, pageSize int) (result []model.InjectShort, err error) {
	// setup & prevent null array
	c := injectCollection()
	result = []model.InjectShort{}

	options := options.Find()
	options.SetSort(bson.D{{Key: "_id", Value: -1}})
	options.SetSkip(int64((page - 1) * pageSize))
	options.SetLimit(int64(pageSize))

	cursor, err := c.Find(ctx, params, options)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var inject model.InjectShort
		cursor.Decode(&inject)
		result = append(result, inject)
	}

	return result, nil
}

// FindInject finds an inject by id
func FindInject(ctx context.Context, id primitive.ObjectID) (result *model.Inject, err error) {
	c := injectCollection()
	result = new(model.Inject)
	err = c.FindOne(ctx, bson.D{{Key: "_id", Value: id}}).Decode(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateInject will insert an inject
func CreateInject(ctx context.Context, inject *model.Inject) (err error) {
	c := injectCollection()
	res, err := c.InsertOne(ctx, inject)
	inject.ID = res.InsertedID.(primitive.ObjectID)
	return
}

// UpdateInject updates an inject
func UpdateInject(ctx context.Context, inject *model.Inject) (err error) {
	c := injectCollection()
	res, err := c.ReplaceOne(ctx, bson.M{"_id": inject.ID}, inject)
	if res.MatchedCount == 0 {
		return errors.New("no document was found")
	}
	return
}

// DeleteInject deletes an inject by a given id
func DeleteInject(ctx context.Context, id primitive.ObjectID, userID primitive.ObjectID) (count int64, err error) {
	c := injectCollection()
	result, err := c.DeleteOne(ctx, bson.D{{Key: "_id", Value: id}})
	return result.DeletedCount, err
}

func injectCollection() *mongo.Collection {
	return MongoSession.Collection("injects")
}
