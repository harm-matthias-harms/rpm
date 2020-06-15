package storage

import (
	"testing"

	"github.com/harm-matthias-harms/rpm/backend/model"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInject(t *testing.T) {
	// setup
	_ = SetMongoDatabase()

	// models
	newexerciseID, _ := primitive.ObjectIDFromHex("0000000000000000000000ab")
	inject := &model.Inject{ExerciseID: primitive.NewObjectID()}

	// tests
	t.Run("create inject", func(t *testing.T) {
		err := CreateInject(nil, inject)
		assert.NoError(t, err)
	})

	t.Run("update inject", func(t *testing.T) {
		//additional model
		notExist := &model.Inject{ExerciseID: primitive.NewObjectID()}

		// change inject
		inject.ExerciseID = newexerciseID
		err := UpdateInject(nil, inject)
		assert.NoError(t, err)

		// don't update non existing
		err = UpdateInject(nil, notExist)
		if assert.Error(t, err) {
			assert.Equal(t, "no document was found", err.Error())
		}
	})

	t.Run("find inject", func(t *testing.T) {
		// additional model
		notExist := &model.Inject{ExerciseID: primitive.NewObjectID()}

		injectFound, err := FindInject(nil, inject.ID)
		if assert.NoError(t, err) {
			assert.Equal(t, newexerciseID, injectFound.ExerciseID)
		}

		_, err = FindInject(nil, notExist.ID)
		assert.Error(t, err)
	})

	t.Run("get injects", func(t *testing.T) {
		result, err := GetInjects(nil, nil, 1, 1)
		if assert.NoError(t, err) {
			assert.Equal(t, 1, len(result))
		}

		// test filter
		filter := map[string]interface{}{"exerciseID": newexerciseID}
		result, err = GetInjects(nil, filter, 1, 1)
		if assert.NoError(t, err) {
			assert.Equal(t, 1, len(result))
		}

		// load all
		result, err = GetInjects(nil, nil, 1, 0)
		if assert.NoError(t, err) {
			assert.GreaterOrEqual(t, len(result), 1)
		}
	})

	t.Run("count injects", func(t *testing.T) {
		count, err := CountInjects(nil, nil)
		if assert.NoError(t, err) {
			assert.GreaterOrEqual(t, count, int64(1))
		}

		filter := map[string]interface{}{"exerciseID": newexerciseID}
		count, err = CountInjects(nil, filter)
		if assert.NoError(t, err) {
			assert.Equal(t, int64(1), count)
		}
	})

	t.Run("deletes inject", func(t *testing.T) {
		c, err := DeleteInject(nil, inject.ID, inject.Author.ID)
		assert.NoError(t, err)
		assert.Equal(t, int64(1), c)
		c, err = DeleteInject(nil, primitive.NewObjectID(), inject.Author.ID)
		assert.NoError(t, err)
		assert.Equal(t, int64(0), c)
	})

	// cleanup
	resetInjectDatabase(inject)
}

func resetInjectDatabase(inject *model.Inject) {
	_, _ = injectCollection().DeleteMany(nil, bson.M{"_id": inject.ID})
}
