package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestPresetValidate(t *testing.T) {
	preset := Preset{}
	err := preset.Validate()
	assert.EqualError(t, err, "author not set")
	preset.Author.ID = primitive.NewObjectID()
	preset.Author.Username = "username"
	err = preset.Validate()
	assert.EqualError(t, err, "created at not set")
	preset.CreatedAt = time.Now()
	err = preset.Validate()
	assert.EqualError(t, err, "title not set")
	preset.Title = "title"
	err = preset.Validate()
	assert.EqualError(t, err, "vital signs not set")
	preset.VitalSigns.OoS = "symptoms"
	err = preset.Validate()
	assert.NoError(t, err)
}

func TestPresetToShortList(t *testing.T) {
	preset := Preset{ID: primitive.NewObjectID(), Author: LimitedUser{ID: primitive.NewObjectID(), Username: "test"}, Title: "test"}
	presets := []Preset{preset}
	result := PresetToShortList(presets)
	first := result[0]
	assert.Equal(t, preset.ID, first.ID)
	assert.Equal(t, preset.Author.ID, first.Author.ID)
	assert.Equal(t, preset.Author.Username, first.Author.Username)
	assert.Equal(t, preset.Title, first.Title)
}
