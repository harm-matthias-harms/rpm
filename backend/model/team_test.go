package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestTeamValidate(t *testing.T) {
	team := Team{}

	err := team.Validate()
	assert.EqualError(t, err, "author not set")

	team.Author.ID = primitive.NewObjectID()
	team.Author.Username = "username"
	err = team.Validate()
	assert.EqualError(t, err, "created at not set")

	team.CreatedAt = time.Now()
	err = team.Validate()
	assert.EqualError(t, err, "title not set")

	team.Title = "title"
	err = team.Validate()
	assert.NoError(t, err)
}
