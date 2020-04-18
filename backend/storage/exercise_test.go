package storage

import (
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/harm-matthias-harms/rpm/backend/model"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestExercise(t *testing.T) {
	// setup
	_ = SetMongoDatabase()

	// models
	currentTime := time.Now().UTC()
	exercise := &model.Exercise{Title: "title"}
	exercise2 := &model.Exercise{Title: "get test", StartTime: currentTime}
	notExist := &model.Exercise{Title: "title"}

	// tests
	t.Run("create exercise", func(t *testing.T) {
		err := CreateExercise(nil, exercise)
		assert.NoError(t, err)
		err = CreateExercise(nil, exercise2)
		assert.NoError(t, err)
	})

	t.Run("update exercise", func(t *testing.T) {
		// change exercise
		exercise.EndTime = currentTime.AddDate(0, 0, 1)
		err := UpdateExercise(nil, exercise, primitive.ObjectID{})
		assert.NoError(t, err)

		// don't update non existing
		err = UpdateExercise(nil, notExist, primitive.ObjectID{})
		if assert.Error(t, err) {
			assert.Equal(t, "no document was found", err.Error())
		}
	})

	t.Run("find exercise", func(t *testing.T) {
		exerciseFound, err := FindExercise(nil, exercise.ID)
		if assert.NoError(t, err) {
			assert.Equal(t, "title", exerciseFound.Title)
			assert.False(t, exerciseFound.EndTime.IsZero())
		}

		_, err = FindExercise(nil, notExist.ID)
		assert.Error(t, err)
	})

	t.Run("deletes exercise", func(t *testing.T) {
		c, err := DeleteExercise(nil, exercise.ID, exercise.Author.ID)
		assert.NoError(t, err)
		assert.Equal(t, int64(1), c)
		c, err = DeleteExercise(nil, primitive.NewObjectID(), exercise.Author.ID)
		assert.NoError(t, err)
		assert.Equal(t, int64(0), c)
	})

	// cleanup
	resetExerciseDatabase(exercise)
	resetExerciseDatabase(exercise2)
}

func resetExerciseDatabase(e *model.Exercise) {
	_, _ = exerciseCollection().DeleteMany(nil, bson.M{"title": e.Title})
}
