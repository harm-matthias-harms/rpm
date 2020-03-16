package storage

import (
	"context"
	"errors"

	"github.com/harm-matthias-harms/rpm/backend/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// FindExercise finds an exercise
func FindExercise(ctx context.Context, id primitive.ObjectID) (result *model.Exercise, err error) {
	c := exerciseCollection()
	result = new(model.Exercise)
	err = c.FindOne(ctx, bson.D{{Key: "_id", Value: id}}).Decode(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateExercise will insert an exercise
func CreateExercise(ctx context.Context, exercise *model.Exercise) (err error) {
	c := exerciseCollection()
	res, err := c.InsertOne(ctx, exercise)
	exercise.ID = res.InsertedID.(primitive.ObjectID)
	return
}

// UpdateExercise updates an exercise
func UpdateExercise(ctx context.Context, exercise *model.Exercise, authorID primitive.ObjectID) (err error) {
	c := exerciseCollection()
	// take care entity is not manipulated and only the owner could update the exercise
	res, err := c.ReplaceOne(ctx, bson.D{{Key: "_id", Value: exercise.ID}, {Key: "author._id", Value: authorID}}, exercise)
	if res.MatchedCount == 0 {
		return errors.New("no document was found")
	}
	return
}

// DeleteExercise deletes an exercise by a given id
func DeleteExercise(ctx context.Context, id primitive.ObjectID, userID primitive.ObjectID) (count int64, err error) {
	c := exerciseCollection()
	result, err := c.DeleteOne(ctx, bson.D{{Key: "_id", Value: id}, {Key: "author._id", Value: userID}})
	return result.DeletedCount, err
}

func exerciseCollection() *mongo.Collection {
	return MongoSession.Collection("exercises")
}
