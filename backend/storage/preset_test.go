package storage

import (
	"testing"

	"github.com/harm-matthias-harms/rpm/backend/model"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestFindPreset(t *testing.T) {
	// First create him because no preset is inserted
	preset := &model.Preset{Title: "title"}
	resetPresetDatabase(preset)
	err := CreatePreset(nil, preset)
	assert.NoError(t, err)
	presetFound, err := FindPreset(nil, preset)
	assert.NoError(t, err)
	assert.Equal(t, "title", presetFound.Title)
	notExist := &model.Preset{Title: "title"}
	_, err = FindPreset(nil, notExist)
	assert.Error(t, err)
}

func TestCreatePreset(t *testing.T) {
	// Creates a User
	preset := &model.Preset{Title: "test"}
	resetPresetDatabase(preset)
	err := CreatePreset(nil, preset)
	assert.NoError(t, err)
}

func resetPresetDatabase(preset *model.Preset) {
	_ = SetMongoDatabase()
	_, _ = userCollection().DeleteOne(nil, bson.M{"title": preset.Title})
}
