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

func TestTeam(t *testing.T) {
	// setup
	_ = storage.SetMongoDatabase()
	header := jsonHeader()
	jwtCookie := loginUser(t)

	// model
	team := model.Team{Title: "test", Medivac: true}
	teamInvalid := model.Team{Title: ""}
	triggerBadRequest := `"title":"title"}`

	// tests
	t.Run("create", func(t *testing.T) {

		// create team
		rec, err := testRequest(http.MethodPost, "/api/teams", strings.NewReader(structToJSONString(team)), HandleTeamCreate, header, nil, jwtCookie)
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			response := &jsonStatus{}
			parseResponse(rec, response)
			assert.Equal(t, http.StatusCreated, rec.Code)
			assert.Equal(t, team.Title, response.Data.(map[string]interface{})["title"])
			assert.Equal(t, team.Medivac, response.Data.(map[string]interface{})["medivac"])
			assert.NotNil(t, response.Data.(map[string]interface{})["author"].(map[string]interface{})["id"])
			assert.NotNil(t, response.Data.(map[string]interface{})["author"].(map[string]interface{})["username"])
			assert.NotNil(t, response.Data.(map[string]interface{})["createdAt"])

			jsonString, _ := json.Marshal(response.Data)
			json.Unmarshal(jsonString, &team)
		}

		// no payload error
		rec, err = testRequest(http.MethodPost, "/api/teams", nil, HandleTeamCreate, header, nil, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
		}

		// bad request
		rec, err = testRequest(http.MethodPost, "/api/teams", strings.NewReader(triggerBadRequest), HandleTeamCreate, header, nil, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
			assert.Equal(t, errorParseRequest, err.(*echo.HTTPError).Message)
		}

		// not create invalid
		rec, err = testRequest(http.MethodPost, "/api/teams", strings.NewReader(structToJSONString(teamInvalid)), HandleTeamCreate, header, nil, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
			assert.Equal(t, "title not set", err.(*echo.HTTPError).Message)
		}
	})

	t.Run("update", func(t *testing.T) {
		// setup
		team.Type = "EMT 2"

		invalidTeam := team
		invalidTeam.CreatedAt = time.Time{}

		id := primitive.NewObjectID()
		falseTeam := team
		falseTeam.ID = id

		// requests
		rec, err := testRequest(http.MethodPut, "/api/teams/:id", strings.NewReader(structToJSONString(team)), HandleTeamEdit, header, map[string]string{"id": team.ID.Hex()}, jwtCookie)
		if assert.NoError(t, err) {
			response := &model.Team{}
			parseResponse(rec, response)
			assert.Equal(t, team.Title, response.Title)
			assert.Equal(t, team.Type, response.Type)
			assert.NotNil(t, response.EditedAt)
			assert.NotNil(t, response.Editor.ID)
			assert.NotNil(t, response.Editor.Username)
		}

		// no team provided
		_, err = testRequest(http.MethodPut, "/api/teams/:id", nil, HandleTeamEdit, header, map[string]string{"id": team.ID.Hex()}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, errorParseRequest, err.(*echo.HTTPError).Message)
		}

		// bad request
		rec, err = testRequest(http.MethodPut, "/api/teams/:id", strings.NewReader(triggerBadRequest), HandleTeamEdit, header, map[string]string{"id": team.ID.Hex()}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
			assert.Equal(t, errorParseRequest, err.(*echo.HTTPError).Message)
		}

		// no id provided
		_, err = testRequest(http.MethodPut, "/api/teams/:id", strings.NewReader(structToJSONString(team)), HandleTeamEdit, header, map[string]string{"id": ""}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, errorNoIDParam, err.(*echo.HTTPError).Message)
		}

		// false id
		_, err = testRequest(http.MethodPut, "/api/teams/:id", strings.NewReader(structToJSONString(team)), HandleTeamEdit, header, map[string]string{"id": primitive.NewObjectID().Hex()}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, errorIDNotMatch, err.(*echo.HTTPError).Message)
		}

		// Invalid
		_, err = testRequest(http.MethodPut, "/api/teams/:id", strings.NewReader(structToJSONString(invalidTeam)), HandleTeamEdit, header, map[string]string{"id": team.ID.Hex()}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, "created at not set", err.(*echo.HTTPError).Message)
		}
		// Could not update
		_, err = testRequest(http.MethodPut, "/api/teams/:id", strings.NewReader(structToJSONString(falseTeam)), HandleTeamEdit, header, map[string]string{"id": id.Hex()}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, errorUpdated, err.(*echo.HTTPError).Message)
		}

	})

	t.Run("find", func(t *testing.T) {
		// no id provided
		_, err := testRequest(http.MethodGet, "/api/teams/:id", nil, HandleTeamFind, nil, map[string]string{"id": ""}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, errorNoIDParam, err.(*echo.HTTPError).Message)
		}
		// false id
		_, err = testRequest(http.MethodGet, "/api/teams/:id", nil, HandleTeamFind, nil, map[string]string{"id": primitive.NewObjectID().Hex()}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, errorFind, err.(*echo.HTTPError).Message)
		}
		// good id
		rec, err := testRequest(http.MethodGet, "/api/teams/:id", nil, HandleTeamFind, nil, map[string]string{"id": team.ID.Hex()}, jwtCookie)
		if assert.NoError(t, err) {
			response := &model.Team{}
			parseResponse(rec, response)
			assert.Equal(t, team.Title, response.Title)
		}
	})

	t.Run("get", func(t *testing.T) {
		q := make(url.Values)
		rec, err := testRequest(http.MethodGet, "/api/teams?"+q.Encode(), nil, HandleTeamsGet, nil, nil, jwtCookie)
		if assert.NoError(t, err) {
			response := &model.TeamsList{}
			parseResponse(rec, response)
			assert.Greater(t, response.Count, int64(0))
		}
		q.Add("title", team.Title)
		rec, err = testRequest(http.MethodGet, "/api/teams?"+q.Encode(), nil, HandleTeamsGet, nil, nil, jwtCookie)
		if assert.NoError(t, err) {
			response := &model.TeamsList{}
			parseResponse(rec, response)
			assert.Greater(t, response.Count, int64(0))
			assert.Equal(t, team.Title, response.Teams[0].Title)
		}
		q.Add("page", "1")
		q.Add("limit", "1")
		rec, err = testRequest(http.MethodGet, "/api/teams?"+q.Encode(), nil, HandleTeamsGet, nil, nil, jwtCookie)
		if assert.NoError(t, err) {
			response := &model.TeamsList{}
			parseResponse(rec, response)
			assert.Equal(t, 1, len(response.Teams))
			assert.Equal(t, team.Title, response.Teams[0].Title)
		}
		q.Del("page")
		q.Del("limit")
		q.Add("author", team.Author.Username)
		rec, err = testRequest(http.MethodGet, "/api/teams?"+q.Encode(), nil, HandleTeamsGet, nil, nil, jwtCookie)
		if assert.NoError(t, err) {
			response := &model.TeamsList{}
			parseResponse(rec, response)
			assert.Greater(t, response.Count, int64(0))
			assert.Equal(t, team.Title, response.Teams[0].Title)
			assert.Equal(t, team.Author.Username, response.Teams[0].Author.Username)
		}
	})

	t.Run("delete", func(t *testing.T) {
		// no id provided
		_, err := testRequest(http.MethodDelete, "/api/teams/:id", nil, HandleTeamDelete, nil, map[string]string{"id": ""}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, errorNoIDParam, err.(*echo.HTTPError).Message)
		}
		// false id
		_, err = testRequest(http.MethodDelete, "/api/teams/:id", nil, HandleTeamDelete, nil, map[string]string{"id": primitive.NewObjectID().Hex()}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, errorDelete, err.(*echo.HTTPError).Message)
		}
		// good id
		_, err = testRequest(http.MethodDelete, "/api/teams/:id", nil, HandleTeamDelete, nil, map[string]string{"id": team.ID.Hex()}, jwtCookie)
		assert.NoError(t, err)
	})

	// cleanup
	resetTeam(&team)
}

func resetTeam(team *model.Team) {
	_, _ = storage.MongoSession.Collection("teams").DeleteMany(nil, bson.M{"title": team.Title})
}
