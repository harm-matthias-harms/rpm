package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestExerciseValidate(t *testing.T) {
	e := &Exercise{}
	err := e.Validate()
	assert.EqualError(t, err, "author not set")

	e.Author.ID = primitive.NewObjectID()
	e.Author.Username = "username"
	err = e.Validate()
	assert.EqualError(t, err, "created at not set")

	e.CreatedAt = time.Now()
	err = e.Validate()
	assert.EqualError(t, err, "no title provided")

	e.Title = "test"
	err = e.Validate()
	assert.EqualError(t, err, "no start or end time provided")

	e.StartTime = time.Now()
	e.EndTime = time.Now()
	err = e.Validate()
	assert.EqualError(t, err, "no teams provided")

	e.Teams = append(e.Teams, ExerciseTeam{Team: Team{}, Trainer: LimitedUser{}})
	err = e.Validate()
	assert.EqualError(t, err, "no role-play manager provided")

	e.RoleplayManager = append(e.RoleplayManager, LimitedUser{})
	err = e.Validate()
	assert.EqualError(t, err, "no make-up center provided")

	e.MakeupCenter = append(e.MakeupCenter, MakeupCenter{Title: "make-up center 1", Account: LimitedUser{}})
	err = e.Validate()
	assert.NoError(t, err)
}
