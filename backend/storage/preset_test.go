package storage

import (
	"testing"

	"github.com/harm-matthias-harms/rpm/backend/model"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestCountPresets(t *testing.T) {
	preset := &model.Preset{Title: "count test"}
	resetPresetDatabase(preset)
	_ = CreatePreset(nil, preset)
	count, err := CountPresets(nil, nil)
	if assert.NoError(t, err) {
		assert.Greater(t, count, int64(0))
	}
	filter := map[string]interface{}{"title": "count test"}
	count, err = CountPresets(nil, filter)
	if assert.NoError(t, err) {
		assert.Equal(t, int64(1), count)
	}
}

func TestGetPresets(t *testing.T) {
	preset := &model.Preset{Title: "get test", VitalSigns: model.VitalSigns{OoS: "symptom"}}
	resetPresetDatabase(preset)
	_ = CreatePreset(nil, preset)
	// test no filter
	result, err := GetPresets(nil, nil, 1, 1)
	if assert.NoError(t, err) {
		assert.Equal(t, 1, len(result))
	}
	// test filter
	filter := map[string]interface{}{"title": "get test"}
	result, err = GetPresets(nil, filter, 1, 1)
	if assert.NoError(t, err) {
		assert.Equal(t, 1, len(result))
	}
	// test multiple filters
	filter["vitalSigns.oos"] = "symptom"
	result, err = GetPresets(nil, filter, 1, 1)
	if assert.NoError(t, err) {
		assert.Equal(t, 1, len(result))
		assert.Equal(t, preset.ID, result[0].ID)
	}
	// test load all
	preset2 := &model.Preset{Title: "get test", VitalSigns: model.VitalSigns{OoS: "symptom"}}
	_ = CreatePreset(nil, preset2)
	result, err = GetPresets(nil, nil, 1, 0)
	if assert.NoError(t, err) {
		assert.Greater(t, len(result), 1)
	}
}

func TestFindPreset(t *testing.T) {
	// First create it because no preset is inserted
	preset := &model.Preset{Title: "title"}
	resetPresetDatabase(preset)
	err := CreatePreset(nil, preset)
	assert.NoError(t, err)
	presetFound, err := FindPreset(nil, preset.ID)
	if assert.NoError(t, err) {
		assert.Equal(t, "title", presetFound.Title)
	}
	notExist := &model.Preset{Title: "title"}
	_, err = FindPreset(nil, notExist.ID)
	assert.Error(t, err)
}

func TestCreatePreset(t *testing.T) {
	// Creates a Preset
	preset := &model.Preset{Title: "test"}
	resetPresetDatabase(preset)
	err := CreatePreset(nil, preset)
	assert.NoError(t, err)
}

func TestUpdatePreset(t *testing.T) {
	preset := &model.Preset{Title: "update preset"}
	resetPresetDatabase(preset)
	err := CreatePreset(nil, preset)
	assert.NoError(t, err)
	preset.VitalSigns.Height = 190
	err = UpdatePreset(nil, preset)
	assert.NoError(t, err)
	presetFound, err := FindPreset(nil, preset.ID)
	if assert.NoError(t, err) {
		assert.Equal(t, 190, presetFound.VitalSigns.Height)
	}
	notExist := &model.Preset{Title: "title"}
	err = UpdatePreset(nil, notExist)
	if assert.Error(t, err) {
		assert.Equal(t, "no document was found", err.Error())
	}
}

func resetPresetDatabase(preset *model.Preset) {
	_ = SetMongoDatabase()
	_, _ = presetCollection().DeleteMany(nil, bson.M{"title": preset.Title})
}
