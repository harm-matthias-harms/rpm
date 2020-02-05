package storage

import (
	"testing"

	"github.com/harm-matthias-harms/rpm/backend/model"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestTeam(t *testing.T) {
	// setup
	_ = SetMongoDatabase()

	// models
	team := &model.Team{Title: "team 1"}

	// tests
	t.Run("create team", func(t *testing.T) {
		err := CreateTeam(nil, team)
		assert.NoError(t, err)
	})

	// cleanup
	resetTeamDatabase(team)
}

func resetTeamDatabase(team *model.Team) {
	_, _ = presetCollection().DeleteMany(nil, bson.M{"title": team.Title})
}
