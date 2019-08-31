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
	assert.EqualError(t, err, "Author not set")
	preset.Author.ID = primitive.NewObjectID()
	preset.Author.Username = "username"
	err = preset.Validate()
	assert.EqualError(t, err, "Created at not set")
	preset.CreatedAt = time.Now()
	err = preset.Validate()
	assert.EqualError(t, err, "Title not set")
	preset.Title = "title"
	err = preset.Validate()
	assert.EqualError(t, err, "Vital signs not set")
	preset.VitalSigns.OoS = "symptoms"
	err = preset.Validate()
	assert.NoError(t, err)
}
