package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/harm-matthias-harms/rpm/backend/model"
	"github.com/harm-matthias-harms/rpm/backend/storage"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestPresetGet(t *testing.T) {
	preset := model.Preset{Title: "test get", Author: model.LimitedUser{Username: "test"}}
	preset2 := model.Preset{Title: "test get", Author: model.LimitedUser{Username: "second"}}
	resetPreset(&preset)
	resetPreset(&preset2)
	_ = storage.CreatePreset(nil, &preset)
	_ = storage.CreatePreset(nil, &preset2)
	jwtCookie := loginUser(t)
	q := make(url.Values)
	rec, err := testRequest(http.MethodGet, "/api/presets?"+q.Encode(), nil, HandlePresetsGet, nil, nil, jwtCookie)
	if assert.NoError(t, err) {
		response := model.PresetsList{}
		defer rec.Result().Body.Close()
		body, _ := ioutil.ReadAll(rec.Result().Body)
		_ = json.Unmarshal(body, &response)
		assert.Greater(t, response.Count, int64(0))
	}
	q.Add("title", preset.Title)
	rec, err = testRequest(http.MethodGet, "/api/presets?"+q.Encode(), nil, HandlePresetsGet, nil, nil, jwtCookie)
	if assert.NoError(t, err) {
		response := model.PresetsList{}
		defer rec.Result().Body.Close()
		body, _ := ioutil.ReadAll(rec.Result().Body)
		_ = json.Unmarshal(body, &response)
		assert.Greater(t, response.Count, int64(0))
		assert.Equal(t, preset.Title, response.Presets[0].Title)
	}
	q.Add("page", "1")
	q.Add("limit", "1")
	rec, err = testRequest(http.MethodGet, "/api/presets?"+q.Encode(), nil, HandlePresetsGet, nil, nil, jwtCookie)
	if assert.NoError(t, err) {
		response := model.PresetsList{}
		defer rec.Result().Body.Close()
		body, _ := ioutil.ReadAll(rec.Result().Body)
		_ = json.Unmarshal(body, &response)
		assert.Equal(t, 1, len(response.Presets))
		assert.Equal(t, preset.Title, response.Presets[0].Title)
	}
	q.Del("page")
	q.Del("limit")
	q.Add("author", preset.Author.Username)
	rec, err = testRequest(http.MethodGet, "/api/presets?"+q.Encode(), nil, HandlePresetsGet, nil, nil, jwtCookie)
	if assert.NoError(t, err) {
		response := model.PresetsList{}
		defer rec.Result().Body.Close()
		body, _ := ioutil.ReadAll(rec.Result().Body)
		_ = json.Unmarshal(body, &response)
		assert.Greater(t, response.Count, int64(0))
		assert.Equal(t, preset.Title, response.Presets[0].Title)
		assert.Equal(t, preset.Author.Username, response.Presets[0].Author.Username)
	}
}

func TestPresetFind(t *testing.T) {
	preset := model.Preset{Title: "test", VitalSigns: model.VitalSigns{OoS: "symptom"}}
	resetPreset(&preset)
	_ = storage.CreatePreset(nil, &preset)

	jwtCookie := loginUser(t)
	// no id provided
	_, err := testRequest(http.MethodGet, "/api/presets/:id", nil, HandlePresetFind, nil, map[string]string{"id": ""}, jwtCookie)
	if assert.Error(t, err) {
		assert.Equal(t, errorNoIDParam, err.(*echo.HTTPError).Message)
	}
	// false id
	_, err = testRequest(http.MethodGet, "/api/presets/:id", nil, HandlePresetFind, nil, map[string]string{"id": primitive.NewObjectID().Hex()}, jwtCookie)
	if assert.Error(t, err) {
		assert.Equal(t, errorFind, err.(*echo.HTTPError).Message)
	}
	// good id
	rec, err := testRequest(http.MethodGet, "/api/presets/:id", nil, HandlePresetFind, nil, map[string]string{"id": preset.ID.Hex()}, jwtCookie)
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
		assert.Equal(t, preset.VitalSigns.OoS, response.Data.(map[string]interface{})["vitalSigns"].(map[string]interface{})["oos"])
		assert.NotNil(t, response.Data.(map[string]interface{})["author"].(map[string]interface{})["id"])
		assert.NotNil(t, response.Data.(map[string]interface{})["author"].(map[string]interface{})["username"])
		assert.NotNil(t, response.Data.(map[string]interface{})["createdAt"])
	}
	// no payload error
	rec, err = testRequest(http.MethodPost, "/api/presets", nil, HandlePresetCreate, header, nil, jwtCookie)
	if assert.Error(t, err) {
		err, ok := err.(*echo.HTTPError)
		if ok {
			assert.Equal(t, http.StatusBadRequest, err.Code)
			assert.Error(t, err)
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

func TestPresetEdit(t *testing.T) {
	preset := model.Preset{Title: "test", Author: model.LimitedUser{ID: primitive.NewObjectID(), Username: "username"}, CreatedAt: time.Now(), VitalSigns: model.VitalSigns{OoS: "symptom"}}
	resetPreset(&preset)
	_ = storage.CreatePreset(nil, &preset)
	helper := int(190)
	preset.VitalSigns.Height = &helper
	presetString, _ := json.Marshal(preset)
	header := http.Header{}
	header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	jwtCookie := loginUser(t)
	// no preset provided
	_, err := testRequest(http.MethodPut, "/api/presets/:id", nil, HandlePresetEdit, header, map[string]string{"id": preset.ID.Hex()}, jwtCookie)
	if assert.Error(t, err) {
		assert.Equal(t, errorParseRequest, err.(*echo.HTTPError).Message)
	}
	// no id provided
	_, err = testRequest(http.MethodPut, "/api/presets/:id", strings.NewReader(string(presetString)), HandlePresetEdit, header, map[string]string{"id": ""}, jwtCookie)
	if assert.Error(t, err) {
		assert.Equal(t, errorNoIDParam, err.(*echo.HTTPError).Message)
	}
	// false id
	_, err = testRequest(http.MethodPut, "/api/presets/:id", strings.NewReader(string(presetString)), HandlePresetEdit, header, map[string]string{"id": primitive.NewObjectID().Hex()}, jwtCookie)
	if assert.Error(t, err) {
		assert.Equal(t, errorIDNotMatch, err.(*echo.HTTPError).Message)
	}
	// Invalid
	invalidPreset := preset
	invalidPreset.CreatedAt = time.Time{}
	invalidPresetString, _ := json.Marshal(invalidPreset)
	_, err = testRequest(http.MethodPut, "/api/presets/:id", strings.NewReader(string(invalidPresetString)), HandlePresetEdit, header, map[string]string{"id": preset.ID.Hex()}, jwtCookie)
	if assert.Error(t, err) {
		assert.Equal(t, "created at not set", err.(*echo.HTTPError).Message)
	}
	// Could not update
	id := primitive.NewObjectID()
	falsePreset := preset
	falsePreset.ID = id
	falsePresetString, _ := json.Marshal(falsePreset)
	_, err = testRequest(http.MethodPut, "/api/presets/:id", strings.NewReader(string(falsePresetString)), HandlePresetEdit, header, map[string]string{"id": id.Hex()}, jwtCookie)
	if assert.Error(t, err) {
		assert.Equal(t, errorUpdated, err.(*echo.HTTPError).Message)
	}

	// good id
	rec, err := testRequest(http.MethodPut, "/api/presets/:id", strings.NewReader(string(presetString)), HandlePresetEdit, header, map[string]string{"id": preset.ID.Hex()}, jwtCookie)
	if assert.NoError(t, err) {
		response := &model.Preset{}
		defer rec.Result().Body.Close()
		body, _ := ioutil.ReadAll(rec.Result().Body)
		_ = json.Unmarshal(body, response)
		assert.Equal(t, preset.Title, response.Title)
		assert.Equal(t, preset.VitalSigns.OoS, response.VitalSigns.OoS)
		assert.Equal(t, preset.VitalSigns.Height, response.VitalSigns.Height)
		assert.NotNil(t, response.EditedAt)
		assert.NotNil(t, response.Editor.ID)
		assert.NotNil(t, response.Editor.Username)
	}
}

func resetPreset(preset *model.Preset) {
	_ = storage.SetMongoDatabase()
	_, _ = storage.MongoSession.Collection("presets").DeleteOne(nil, bson.M{"title": preset.Title})
}
