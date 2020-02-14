package storage

import (
	"context"

	"github.com/harm-matthias-harms/rpm/backend/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CountTeams returns the number of total teams inserted
func CountTeams(ctx context.Context, params map[string]interface{}) (int64, error) {
	if len(params) == 0 {
		return teamCollection().CountDocuments(ctx, bson.M{})
	}
	return teamCollection().CountDocuments(ctx, params)
}

// GetTeams returns an array of the teams
func GetTeams(ctx context.Context, params map[string]interface{}, page int, pageSize int) (result []model.Team, err error) {
	// setup & prevent null array
	result = []model.Team{}
	c := teamCollection()

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
		var team model.Team
		cursor.Decode(&team)
		result = append(result, team)
	}

	return result, nil
}

// FindTeam finds a specific team
func FindTeam(ctx context.Context, id primitive.ObjectID) (result *model.Team, err error) {
	c := teamCollection()
	result = new(model.Team)
	err = c.FindOne(ctx, bson.D{{Key: "_id", Value: id}}).Decode(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateTeam will insert a team
func CreateTeam(ctx context.Context, team *model.Team) (err error) {
	c := teamCollection()
	res, err := c.InsertOne(ctx, team)
	team.ID = res.InsertedID.(primitive.ObjectID)
	return
}

func teamCollection() *mongo.Collection {
	return MongoSession.Collection("teams")
}
