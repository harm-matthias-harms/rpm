package handler

import (
	"encoding/json"
	"net/http"
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

func TestExercise(t *testing.T) {
	// setup
	_ = storage.SetMongoDatabase()
	header := jsonHeader()
	jwtCookie := loginUser(t)

	// model
	exercise := model.Exercise{Title: "test", StartTime: time.Now(), EndTime: time.Now(), Teams: []model.ExerciseTeam{{Team: model.Team{}, Trainer: model.LimitedUser{Username: "test trainer", Code: "ACBDFE"}}}, RoleplayManager: []model.LimitedUser{{Username: "test role play manager", Code: "FJKEFJ"}}, MakeupCenter: []model.MakeupCenter{{Account: model.LimitedUser{Username: "test make-up centere", Code: "HJFTGE"}}}}
	exerciseInvalid := model.Exercise{Title: "test"}
	triggerBadRequest := `"title":"title"}`

	// tests
	t.Run("create", func(t *testing.T) {

		// create exercise
		rec, err := testRequest(http.MethodPost, "/api/exercises", strings.NewReader(structToJSONString(exercise)), HandleExerciseCreate, header, nil, jwtCookie)
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			response := &jsonStatus{}
			parseResponse(rec, response)
			assert.Equal(t, http.StatusCreated, rec.Code)
			assert.Equal(t, exercise.Title, response.Data.(map[string]interface{})["title"])
			assert.NotNil(t, response.Data.(map[string]interface{})["author"].(map[string]interface{})["id"])
			assert.NotNil(t, response.Data.(map[string]interface{})["author"].(map[string]interface{})["username"])
			assert.NotNil(t, response.Data.(map[string]interface{})["createdAt"])

			jsonString, _ := json.Marshal(response.Data)
			json.Unmarshal(jsonString, &exercise)
		}

		// no payload error
		rec, err = testRequest(http.MethodPost, "/api/exercises", nil, HandleExerciseCreate, header, nil, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
		}

		// bad request
		rec, err = testRequest(http.MethodPost, "/api/exercises", strings.NewReader(triggerBadRequest), HandleExerciseCreate, header, nil, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
			assert.Equal(t, errorParseRequest, err.(*echo.HTTPError).Message)
		}

		// not create invalid
		rec, err = testRequest(http.MethodPost, "/api/exercises", strings.NewReader(structToJSONString(exerciseInvalid)), HandleExerciseCreate, header, nil, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
			assert.Equal(t, "no start or end time provided", err.(*echo.HTTPError).Message)
		}
	})

	t.Run("update", func(t *testing.T) {
		// setup
		exercise.EditedAt.AddDate(0, 0, 1)

		invalidExercise := exercise
		invalidExercise.CreatedAt = time.Time{}

		id := primitive.NewObjectID()
		falseExercise := exercise
		falseExercise.ID = id

		// requests
		rec, err := testRequest(http.MethodPut, "/api/exercises/:id", strings.NewReader(structToJSONString(exercise)), HandleExerciseEdit, header, map[string]string{"id": exercise.ID.Hex()}, jwtCookie)
		if assert.NoError(t, err) {
			response := &model.Exercise{}
			parseResponse(rec, response)
			assert.Equal(t, exercise.Title, response.Title)
			assert.NotNil(t, response.EditedAt)
		}

		// no exercise provided
		_, err = testRequest(http.MethodPut, "/api/exercises/:id", nil, HandleExerciseEdit, header, map[string]string{"id": exercise.ID.Hex()}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, errorParseRequest, err.(*echo.HTTPError).Message)
		}

		// bad request
		rec, err = testRequest(http.MethodPut, "/api/exercises/:id", strings.NewReader(triggerBadRequest), HandleExerciseEdit, header, map[string]string{"id": exercise.ID.Hex()}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
			assert.Equal(t, errorParseRequest, err.(*echo.HTTPError).Message)
		}

		// no id provided
		_, err = testRequest(http.MethodPut, "/api/exercises/:id", strings.NewReader(structToJSONString(exercise)), HandleExerciseEdit, header, map[string]string{"id": ""}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, errorNoIDParam, err.(*echo.HTTPError).Message)
		}

		// false id
		_, err = testRequest(http.MethodPut, "/api/exercises/:id", strings.NewReader(structToJSONString(exercise)), HandleExerciseEdit, header, map[string]string{"id": primitive.NewObjectID().Hex()}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, errorIDNotMatch, err.(*echo.HTTPError).Message)
		}

		// Invalid
		_, err = testRequest(http.MethodPut, "/api/exercises/:id", strings.NewReader(structToJSONString(invalidExercise)), HandleExerciseEdit, header, map[string]string{"id": exercise.ID.Hex()}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, "created at not set", err.(*echo.HTTPError).Message)
		}
		// Could not update
		_, err = testRequest(http.MethodPut, "/api/exercises/:id", strings.NewReader(structToJSONString(falseExercise)), HandleExerciseEdit, header, map[string]string{"id": id.Hex()}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, errorFind, err.(*echo.HTTPError).Message)
		}

	})

	t.Run("find", func(t *testing.T) {
		// no id provided
		_, err := testRequest(http.MethodGet, "/api/exercises/:id", nil, HandleExerciseFind, nil, map[string]string{"id": ""}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, errorNoIDParam, err.(*echo.HTTPError).Message)
		}
		// false id
		_, err = testRequest(http.MethodGet, "/api/exercises/:id", nil, HandleExerciseFind, nil, map[string]string{"id": primitive.NewObjectID().Hex()}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, errorFind, err.(*echo.HTTPError).Message)
		}
		// good id
		rec, err := testRequest(http.MethodGet, "/api/exercises/:id", nil, HandleExerciseFind, nil, map[string]string{"id": exercise.ID.Hex()}, jwtCookie)
		if assert.NoError(t, err) {
			response := &model.Exercise{}
			parseResponse(rec, response)
			assert.Equal(t, exercise.Title, response.Title)
		}
	})

	t.Run("find - removes user codes if exercise not owned", func(t *testing.T) {
		notOwnedExercise := &model.Exercise{Title: "not owned", Teams: []model.ExerciseTeam{{Team: model.Team{}, Trainer: model.LimitedUser{Username: "test trainer", Code: "ACBDFE"}}}, RoleplayManager: []model.LimitedUser{{Username: "test role play manager", Code: "FJKEFJ"}}, MakeupCenter: []model.MakeupCenter{{Account: model.LimitedUser{Username: "test make-up centere", Code: "HJFTGE"}}}}
		_ = storage.CreateExercise(nil, notOwnedExercise)
		// good id
		rec, err := testRequest(http.MethodGet, "/api/exercises/:id", nil, HandleExerciseFind, nil, map[string]string{"id": notOwnedExercise.ID.Hex()}, jwtCookie)
		if assert.NoError(t, err) {
			response := &model.Exercise{}
			parseResponse(rec, response)
			assert.Equal(t, notOwnedExercise.Title, response.Title)
			assert.Equal(t, "", response.Teams[0].Trainer.Code)
			assert.Equal(t, "", response.RoleplayManager[0].Code)
			assert.Equal(t, "", response.MakeupCenter[0].Account.Code)
		}
		resetExercise(notOwnedExercise)
	})

	t.Run("delete", func(t *testing.T) {
		// no id provided
		_, err := testRequest(http.MethodDelete, "/api/exercises/:id", nil, HandleExerciseDelete, nil, map[string]string{"id": ""}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, errorNoIDParam, err.(*echo.HTTPError).Message)
		}
		// false id
		_, err = testRequest(http.MethodDelete, "/api/exercises/:id", nil, HandleExerciseDelete, nil, map[string]string{"id": primitive.NewObjectID().Hex()}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, errorFind, err.(*echo.HTTPError).Message)
		}
		// good id
		_, err = testRequest(http.MethodDelete, "/api/exercises/:id", nil, HandleExerciseDelete, nil, map[string]string{"id": exercise.ID.Hex()}, jwtCookie)
		assert.NoError(t, err)
	})

	// cleanup
	resetExercise(&exercise)
}

func resetExercise(exercise *model.Exercise) {
	_, _ = storage.MongoSession.Collection("exercises").DeleteMany(nil, bson.M{"title": exercise.Title})
}
