package storage

import (
	"context"

	"github.com/harm-matthias-harms/rpm/backend/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

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
