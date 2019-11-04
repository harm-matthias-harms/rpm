package storage

import (
	"testing"

	"github.com/harm-matthias-harms/rpm/backend/model"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestCountMedicalCases(t *testing.T) {
	mc := &model.MedicalCase{Title: "count test"}
	resetMedicalCaseDatabase(mc)
	_ = CreateMedicalCase(nil, mc)
	count, err := CountMedicalCases(nil, nil)
	if assert.NoError(t, err) {
		assert.Greater(t, count, int64(0))
	}
	filter := map[string]interface{}{"title": "count test"}
	count, err = CountMedicalCases(nil, filter)
	if assert.NoError(t, err) {
		assert.Equal(t, int64(1), count)
	}
}

func TestGetMedicalCases(t *testing.T) {
	mc := &model.MedicalCase{Title: "get medical cases test"}
	mc.GeneralInformation.Surgical = true
	resetMedicalCaseDatabase(mc)
	_ = CreateMedicalCase(nil, mc)
	// test no filter
	result, err := GetMedicalCases(nil, nil, 1, 1)
	if assert.NoError(t, err) {
		assert.Equal(t, 1, len(result))
	}
	// test filter
	filter := map[string]interface{}{"title": "get medical cases test"}
	result, err = GetMedicalCases(nil, filter, 1, 1)
	if assert.NoError(t, err) {
		assert.Equal(t, 1, len(result))
	}
	// test multiple filters
	filter["general_information.surgical"] = true
	result, err = GetMedicalCases(nil, filter, 1, 1)
	if assert.NoError(t, err) {
		assert.Equal(t, 1, len(result))
		assert.Equal(t, mc.ID, result[0].ID)
	}
	// test load all
	mc2 := &model.MedicalCase{Title: "get test"}
	resetMedicalCaseDatabase(mc2)
	_ = CreateMedicalCase(nil, mc2)
	result, err = GetMedicalCases(nil, nil, 1, 0)
	if assert.NoError(t, err) {
		assert.Greater(t, len(result), 1)
	}
}

func TestFindMedicalCase(t *testing.T) {
	// First create it because no medical case is inserted
	mc := &model.MedicalCase{Title: "test find medical case"}
	resetMedicalCaseDatabase(mc)
	err := CreateMedicalCase(nil, mc)
	assert.NoError(t, err)
	mcFound, err := FindMedicalCase(nil, mc.ID)
	if assert.NoError(t, err) {
		assert.Equal(t, "test find medical case", mcFound.Title)
	}
	notExist := &model.MedicalCase{Title: "title"}
	_, err = FindMedicalCase(nil, notExist.ID)
	assert.Error(t, err)
}

func TestCreateMedicalCase(t *testing.T) {
	mc := &model.MedicalCase{Title: "test create medical case"}
	resetMedicalCaseDatabase(mc)
	err := CreateMedicalCase(nil, mc)
	assert.NoError(t, err)
}

func resetMedicalCaseDatabase(preset *model.MedicalCase) {
	_ = SetMongoDatabase()
	_, _ = mcCollection().DeleteMany(nil, bson.M{"title": preset.Title})
}
