package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInjectValidate(t *testing.T) {
	inject := &Inject{}
	team := Team{Title: "Team 1", Author: LimitedUser{ID: primitive.NewObjectID(), Username: "teamauthor"}, CreatedAt: time.Now()}
	
	err := inject.Validate()
	assert.EqualError(t, err, "author not set")

	inject.Author.ID = primitive.NewObjectID()
	inject.Author.Username = "username"
	err = inject.Validate()
	assert.EqualError(t, err, "created at not set")

	inject.CreatedAt = time.Now()
	err = inject.Validate()
	assert.EqualError(t, err, "exercise id is not set")

	inject.ExerciseID = primitive.NewObjectID()
	err = inject.Validate()
	assert.EqualError(t, err, "team is not set")

	inject.Team = team
	err = inject.Validate()
	assert.EqualError(t, err, "medical case is not set")

	inject.MedicalCase = MedicalCase{Title: "Medical Case"}
	err = inject.Validate()
	assert.EqualError(t, err, "role player is not set")

	inject.Roleplayer = Roleplayer{Fullname: "Don Joe"}
	err = inject.Validate()
	assert.NoError(t, err)
}
