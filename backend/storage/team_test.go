package storage

import (
	"testing"

	"github.com/harm-matthias-harms/rpm/backend/model"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	t.Run("update team", func(t *testing.T) {
		//additional model
		notExist := &model.Team{Title: "title"}

		// change team
		team.Title = "team 2"
		err := UpdateTeam(nil, team)
		assert.NoError(t, err)

		// don't update non existing
		err = UpdateTeam(nil, notExist)
		if assert.Error(t, err) {
			assert.Equal(t, "no document was found", err.Error())
		}
	})

	t.Run("find team", func(t *testing.T) {
		// additional model
		notExist := &model.Team{Title: "title"}

		teamFound, err := FindTeam(nil, team.ID)
		if assert.NoError(t, err) {
			assert.Equal(t, "team 2", teamFound.Title)
		}

		_, err = FindTeam(nil, notExist.ID)
		assert.Error(t, err)
	})

	t.Run("get teams", func(t *testing.T) {
		result, err := GetTeams(nil, nil, 1, 1)
		if assert.NoError(t, err) {
			assert.Equal(t, 1, len(result))
		}

		// test filter
		filter := map[string]interface{}{"title": "team 2"}
		result, err = GetTeams(nil, filter, 1, 1)
		if assert.NoError(t, err) {
			assert.Equal(t, 1, len(result))
		}

		// load all
		result, err = GetTeams(nil, nil, 1, 0)
		if assert.NoError(t, err) {
			assert.GreaterOrEqual(t, len(result), 1)
		}
	})

	t.Run("count teams", func(t *testing.T) {
		count, err := CountTeams(nil, nil)
		if assert.NoError(t, err) {
			assert.GreaterOrEqual(t, count, int64(1))
		}

		filter := map[string]interface{}{"title": "team 2"}
		count, err = CountTeams(nil, filter)
		if assert.NoError(t, err) {
			assert.Equal(t, int64(1), count)
		}
	})

	t.Run("deletes team", func(t *testing.T) {
		c, err := DeleteTeam(nil, team.ID, team.Author.ID)
		assert.NoError(t, err)
		assert.Equal(t, int64(1), c)
		c, err = DeleteTeam(nil, primitive.NewObjectID(), team.Author.ID)
		assert.NoError(t, err)
		assert.Equal(t, int64(0), c)
	})

	// cleanup
	resetTeamDatabase(team)
}

func resetTeamDatabase(team *model.Team) {
	_, _ = teamCollection().DeleteMany(nil, bson.M{"title": team.Title})
}
