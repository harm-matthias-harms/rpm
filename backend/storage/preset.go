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

// CountPresets returns the number of total documents inserted
func CountPresets(ctx context.Context, params map[string]interface{}) (int64, error) {
	if len(params) == 0 {
		return presetCollection().CountDocuments(ctx, bson.M{})
	}
	return presetCollection().CountDocuments(ctx, params)
}

// GetPresets returns an array of the presets in the short version
func GetPresets(ctx context.Context, params map[string]interface{}, page int, pageSize int) (result []model.PresetShort, err error) {
	// setup & prevent null array
	result = []model.PresetShort{}
	c := presetCollection()

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
		var preset model.PresetShort
		cursor.Decode(&preset)
		result = append(result, preset)
	}

	return result, nil
}

// FindPreset finds a preset
func FindPreset(ctx context.Context, id primitive.ObjectID) (result *model.Preset, err error) {
	c := presetCollection()
	result = new(model.Preset)
	err = c.FindOne(ctx, bson.D{{Key: "_id", Value: id}}).Decode(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreatePreset will insert a preset
func CreatePreset(ctx context.Context, preset *model.Preset) (err error) {
	c := presetCollection()
	res, err := c.InsertOne(ctx, preset)
	preset.ID = res.InsertedID.(primitive.ObjectID)
	return
}

// UpdatePreset updates a preset
func UpdatePreset(ctx context.Context, preset *model.Preset) (err error) {
	c := presetCollection()
	res, err := c.ReplaceOne(ctx, bson.M{"_id": preset.ID}, preset)
	if res.MatchedCount == 0 {
		return errors.New("no document was found")
	}
	return
}

// DeletePreset deletes a preset by a given id
func DeletePreset(ctx context.Context, id primitive.ObjectID, userID primitive.ObjectID) (count int64, err error) {
	c := presetCollection()
	result, err := c.DeleteOne(ctx, bson.D{{Key: "_id", Value: id}, {Key: "author._id", Value: userID}})
	return result.DeletedCount, err
}

func presetCollection() *mongo.Collection {
	return MongoSession.Collection("presets")
}
