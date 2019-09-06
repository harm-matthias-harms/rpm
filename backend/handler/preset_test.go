package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/harm-matthias-harms/rpm/backend/model"
	"github.com/harm-matthias-harms/rpm/backend/storage"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestPresetFind(t *testing.T) {
	preset := model.Preset{Title: "test", VitalSigns: model.VitalSigns{OoS: "symptom"}}
	resetPreset(&preset)
	_ = storage.CreatePreset(nil, &preset)
	header := http.Header{}
	header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	jwtCookie := loginUser(t)
	// no id provided
	_, err := testRequest(http.MethodGet, "/api/presets/:id", nil, HandlePresetFind, header, map[string]string{"id": ""}, jwtCookie)
	if assert.Error(t, err) {
		assert.Equal(t, "no or false id provided", err.(*echo.HTTPError).Message)
	}
	// false id
	_, err = testRequest(http.MethodGet, "/api/presets/:id", nil, HandlePresetFind, header, map[string]string{"id": primitive.NewObjectID().Hex()}, jwtCookie)
	if assert.Error(t, err) {
		assert.Equal(t, "couldn't find preset", err.(*echo.HTTPError).Message)
	}
	// good id
	rec, err := testRequest(http.MethodGet, "/api/presets/:id", nil, HandlePresetFind, header, map[string]string{"id": preset.ID.Hex()}, jwtCookie)
	if assert.NoError(t, err) {
		response := &model.Preset{}
		defer rec.Result().Body.Close()
		body, _ := ioutil.ReadAll(rec.Result().Body)
		_ = json.Unmarshal(body, response)
		assert.Equal(t, preset.Title, response.Title)
		assert.Equal(t, preset.VitalSigns.OoS, response.VitalSigns.OoS)
	}
}

func TestPresetCreate(t *testing.T) {
	preset := model.Preset{Title: "test", VitalSigns: model.VitalSigns{OoS: "symptom"}}
	resetPreset(&preset)
	presetString, _ := json.Marshal(preset)
	header := http.Header{}
	header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	jwtCookie := loginUser(t)

	// create preset
	rec, err := testRequest(http.MethodPost, "/api/presets", strings.NewReader(string(presetString)), HandlePresetCreate, header, nil, jwtCookie)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		response := &jsonStatus{}
		defer rec.Result().Body.Close()
		body, _ := ioutil.ReadAll(rec.Result().Body)
		_ = json.Unmarshal(body, response)
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, preset.Title, response.Data.(map[string]interface{})["title"])
		fmt.Println(response.Data)
		assert.Equal(t, preset.VitalSigns.OoS, response.Data.(map[string]interface{})["vital_signs"].(map[string]interface{})["oos"])
		assert.NotNil(t, response.Data.(map[string]interface{})["author"].(map[string]interface{})["id"])
		assert.NotNil(t, response.Data.(map[string]interface{})["author"].(map[string]interface{})["username"])
		assert.NotNil(t, response.Data.(map[string]interface{})["created_at"])
	}
	// no payload error
	rec, err = testRequest(http.MethodPost, "/api/presets", nil, HandlePresetCreate, header, nil, jwtCookie)
	if assert.Error(t, err) {
		err, ok := err.(*echo.HTTPError)
		if ok {
			assert.Equal(t, http.StatusBadRequest, err.Code)
			assert.Equal(t, "couldn't parse request", err.Message)
		}
	}
	// not create invalid
	presetInvalid := model.Preset{Title: "test"}
	resetPreset(&presetInvalid)
	presetInvalidString, _ := json.Marshal(presetInvalid)
	rec, err = testRequest(http.MethodPost, "/api/presets", strings.NewReader(string(presetInvalidString)), HandlePresetCreate, header, nil, jwtCookie)
	if assert.Error(t, err) {
		err, ok := err.(*echo.HTTPError)
		if ok {
			assert.Equal(t, http.StatusBadRequest, err.Code)
			assert.Equal(t, "vital signs not set", err.Message)
		}
	}
}

func resetPreset(preset *model.Preset) {
	_ = storage.SetMongoDatabase()
	_, _ = storage.MongoSession.Collection("preset").DeleteOne(nil, bson.M{"title": preset.Title})
}
