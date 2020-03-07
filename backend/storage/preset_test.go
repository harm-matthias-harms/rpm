package storage

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"

	"github.com/harm-matthias-harms/rpm/backend/model"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestPreset(t *testing.T) {
	// setup
	_ = SetMongoDatabase()

	// models
	preset := &model.Preset{Title: "title"}
	preset2 := &model.Preset{Title: "get test", VitalSigns: model.VitalSigns{OoS: "symptom"}}

	// tests
	t.Run("create preset", func(t *testing.T) {
		err := CreatePreset(nil, preset)
		assert.NoError(t, err)
		err = CreatePreset(nil, preset2)
		assert.NoError(t, err)
	})

	t.Run("update preset", func(t *testing.T) {
		//additional model
		notExist := &model.Preset{Title: "title"}

		// change preset
		helper := int(150)
		preset.VitalSigns.Pulse = &helper
		err := UpdatePreset(nil, preset)
		assert.NoError(t, err)

		// don't update non existing
		err = UpdatePreset(nil, notExist)
		if assert.Error(t, err) {
			assert.Equal(t, "no document was found", err.Error())
		}
	})

	t.Run("find preset", func(t *testing.T) {
		// additional model
		notExist := &model.Preset{Title: "title"}

		presetFound, err := FindPreset(nil, preset.ID)
		if assert.NoError(t, err) {
			assert.Equal(t, "title", presetFound.Title)
			assert.Equal(t, 150, *presetFound.VitalSigns.Pulse)
		}

		_, err = FindPreset(nil, notExist.ID)
		assert.Error(t, err)
	})

	t.Run("get presets", func(t *testing.T) {
		result, err := GetPresets(nil, nil, 1, 1)
		if assert.NoError(t, err) {
			assert.Equal(t, 1, len(result))
		}

		// test filter
		filter := map[string]interface{}{"title": "title"}
		result, err = GetPresets(nil, filter, 1, 1)
		if assert.NoError(t, err) {
			assert.Equal(t, 1, len(result))
		}

		// test multiple filters
		filter["vitalSigns.pulse"] = 150
		result, err = GetPresets(nil, filter, 1, 1)
		if assert.NoError(t, err) {
			assert.Equal(t, 1, len(result))
			assert.Equal(t, preset.ID, result[0].ID)
		}

		// load all
		result, err = GetPresets(nil, nil, 1, 0)
		if assert.NoError(t, err) {
			assert.Greater(t, len(result), 1)
		}
	})

	t.Run("count presets", func(t *testing.T) {
		count, err := CountPresets(nil, nil)
		if assert.NoError(t, err) {
			assert.Greater(t, count, int64(0))
		}

		filter := map[string]interface{}{"title": "title"}
		count, err = CountPresets(nil, filter)
		if assert.NoError(t, err) {
			assert.Equal(t, int64(1), count)
		}
	})

	t.Run("deletes preset", func(t *testing.T) {
		c, err := DeletePreset(nil, preset.ID, preset.Author.ID)
		assert.NoError(t, err)
		assert.Equal(t, int64(1), c)
		c, err = DeletePreset(nil, primitive.NewObjectID(), preset.Author.ID)
		assert.NoError(t, err)
		assert.Equal(t, int64(0), c)
	})

	// cleanup
	resetPresetDatabase(preset)
	resetPresetDatabase(preset2)
}

func resetPresetDatabase(preset *model.Preset) {
	_, _ = presetCollection().DeleteMany(nil, bson.M{"title": preset.Title})
}
