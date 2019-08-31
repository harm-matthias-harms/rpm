package storage

import (
	"context"

	"github.com/harm-matthias-harms/rpm/backend/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// FindPreset finds a preset
func FindPreset(ctx context.Context, preset *model.Preset) (result *model.Preset, err error) {
	c := presetCollection()
	result = new(model.Preset)
	err = c.FindOne(ctx, bson.D{{Key: "_id", Value: preset.ID}}).Decode(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreatePreset will create a preset
func CreatePreset(ctx context.Context, preset *model.Preset) (err error) {
	c := presetCollection()
	res, err := c.InsertOne(ctx, preset)
	preset.ID = res.InsertedID.(primitive.ObjectID)
	return
}

func presetCollection() *mongo.Collection {
	return MongoSession.Collection("preset")
}
