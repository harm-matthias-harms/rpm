package handler

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/harm-matthias-harms/rpm/backend/model"
	"github.com/harm-matthias-harms/rpm/backend/storage"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
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

		// create preset
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

	// cleanup
	resetTeam(&team)
}

func resetTeam(team *model.Team) {
	_, _ = storage.MongoSession.Collection("teams").DeleteMany(nil, bson.M{"title": team.Title})
}
