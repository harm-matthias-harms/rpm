package handler

import (
	"encoding/json"
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

func TestPreset(t *testing.T) {
	// setup
	_ = storage.SetMongoDatabase()
	header := jsonHeader()
	jwtCookie := loginUser(t)

	// model
	preset := model.Preset{Title: "test", VitalSigns: model.VitalSigns{OoS: "symptom"}}
	presetInvalid := model.Preset{Title: "test"}
	triggerBadRequest := `"title":"title"}`

	// tests
	t.Run("create", func(t *testing.T) {

		// create preset
		rec, err := testRequest(http.MethodPost, "/api/presets", strings.NewReader(structToJSONString(preset)), HandlePresetCreate, header, nil, jwtCookie)
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			response := &jsonStatus{}
			parseResponse(rec, response)
			assert.Equal(t, http.StatusCreated, rec.Code)
			assert.Equal(t, preset.Title, response.Data.(map[string]interface{})["title"])
			assert.Equal(t, preset.VitalSigns.OoS, response.Data.(map[string]interface{})["vitalSigns"].(map[string]interface{})["oos"])
			assert.NotNil(t, response.Data.(map[string]interface{})["author"].(map[string]interface{})["id"])
			assert.NotNil(t, response.Data.(map[string]interface{})["author"].(map[string]interface{})["username"])
			assert.NotNil(t, response.Data.(map[string]interface{})["createdAt"])

			jsonString, _ := json.Marshal(response.Data)
			json.Unmarshal(jsonString, &preset)
		}

		// no payload error
		rec, err = testRequest(http.MethodPost, "/api/presets", nil, HandlePresetCreate, header, nil, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
		}

		// bad request
		rec, err = testRequest(http.MethodPost, "/api/presets", strings.NewReader(triggerBadRequest), HandlePresetCreate, header, nil, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
			assert.Equal(t, errorParseRequest, err.(*echo.HTTPError).Message)
		}

		// not create invalid
		rec, err = testRequest(http.MethodPost, "/api/presets", strings.NewReader(structToJSONString(presetInvalid)), HandlePresetCreate, header, nil, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
			assert.Equal(t, "vital signs not set", err.(*echo.HTTPError).Message)
		}
	})

	t.Run("update", func(t *testing.T) {
		// setup
		helper := int(190)
		preset.VitalSigns.Height = &helper

		invalidPreset := preset
		invalidPreset.CreatedAt = time.Time{}

		id := primitive.NewObjectID()
		falsePreset := preset
		falsePreset.ID = id

		// requests
		rec, err := testRequest(http.MethodPut, "/api/presets/:id", strings.NewReader(structToJSONString(preset)), HandlePresetEdit, header, map[string]string{"id": preset.ID.Hex()}, jwtCookie)
		if assert.NoError(t, err) {
			response := &model.Preset{}
			parseResponse(rec, response)
			assert.Equal(t, preset.Title, response.Title)
			assert.Equal(t, preset.VitalSigns.OoS, response.VitalSigns.OoS)
			assert.Equal(t, preset.VitalSigns.Height, response.VitalSigns.Height)
			assert.NotNil(t, response.EditedAt)
			assert.NotNil(t, response.Editor.ID)
			assert.NotNil(t, response.Editor.Username)
		}

		// no preset provided
		_, err = testRequest(http.MethodPut, "/api/presets/:id", nil, HandlePresetEdit, header, map[string]string{"id": preset.ID.Hex()}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, errorParseRequest, err.(*echo.HTTPError).Message)
		}

		// bad request
		rec, err = testRequest(http.MethodPut, "/api/presets/:id", strings.NewReader(triggerBadRequest), HandlePresetEdit, header, map[string]string{"id": preset.ID.Hex()}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
			assert.Equal(t, errorParseRequest, err.(*echo.HTTPError).Message)
		}

		// no id provided
		_, err = testRequest(http.MethodPut, "/api/presets/:id", strings.NewReader(structToJSONString(preset)), HandlePresetEdit, header, map[string]string{"id": ""}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, errorNoIDParam, err.(*echo.HTTPError).Message)
		}

		// false id
		_, err = testRequest(http.MethodPut, "/api/presets/:id", strings.NewReader(structToJSONString(preset)), HandlePresetEdit, header, map[string]string{"id": primitive.NewObjectID().Hex()}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, errorIDNotMatch, err.(*echo.HTTPError).Message)
		}

		// Invalid
		_, err = testRequest(http.MethodPut, "/api/presets/:id", strings.NewReader(structToJSONString(invalidPreset)), HandlePresetEdit, header, map[string]string{"id": preset.ID.Hex()}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, "created at not set", err.(*echo.HTTPError).Message)
		}
		// Could not update
		_, err = testRequest(http.MethodPut, "/api/presets/:id", strings.NewReader(structToJSONString(falsePreset)), HandlePresetEdit, header, map[string]string{"id": id.Hex()}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, errorUpdated, err.(*echo.HTTPError).Message)
		}

	})

	t.Run("find", func(t *testing.T) {
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
			parseResponse(rec, response)
			assert.Equal(t, preset.Title, response.Title)
			assert.Equal(t, preset.VitalSigns.OoS, response.VitalSigns.OoS)
		}
	})

	t.Run("get", func(t *testing.T) {
		q := make(url.Values)
		rec, err := testRequest(http.MethodGet, "/api/presets?"+q.Encode(), nil, HandlePresetsGet, nil, nil, jwtCookie)
		if assert.NoError(t, err) {
			response := &model.PresetsList{}
			parseResponse(rec, response)
			assert.Greater(t, response.Count, int64(0))
		}
		q.Add("title", preset.Title)
		rec, err = testRequest(http.MethodGet, "/api/presets?"+q.Encode(), nil, HandlePresetsGet, nil, nil, jwtCookie)
		if assert.NoError(t, err) {
			response := &model.PresetsList{}
			parseResponse(rec, response)
			assert.Greater(t, response.Count, int64(0))
			assert.Equal(t, preset.Title, response.Presets[0].Title)
		}
		q.Add("page", "1")
		q.Add("limit", "1")
		rec, err = testRequest(http.MethodGet, "/api/presets?"+q.Encode(), nil, HandlePresetsGet, nil, nil, jwtCookie)
		if assert.NoError(t, err) {
			response := &model.PresetsList{}
			parseResponse(rec, response)
			assert.Equal(t, 1, len(response.Presets))
			assert.Equal(t, preset.Title, response.Presets[0].Title)
		}
		q.Del("page")
		q.Del("limit")
		q.Add("author", preset.Author.Username)
		rec, err = testRequest(http.MethodGet, "/api/presets?"+q.Encode(), nil, HandlePresetsGet, nil, nil, jwtCookie)
		if assert.NoError(t, err) {
			response := &model.PresetsList{}
			parseResponse(rec, response)
			assert.Greater(t, response.Count, int64(0))
			assert.Equal(t, preset.Title, response.Presets[0].Title)
			assert.Equal(t, preset.Author.Username, response.Presets[0].Author.Username)
		}
	})

	// cleanup
	resetPreset(&preset)
}

func resetPreset(preset *model.Preset) {
	_, _ = storage.MongoSession.Collection("presets").DeleteMany(nil, bson.M{"title": preset.Title})
}
