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

func TestInject(t *testing.T) {
	// setup
	_ = storage.SetMongoDatabase()
	header := jsonHeader()
	jwtCookie := loginUser(t)

	// init other required objects for the inject
	user, _ := storage.FindUser(nil, &model.User{Username: "testuser"})
	limitedUser := model.LimitedUser{ID: user.ID, Username: user.Username}
	team := model.Team{ID: primitive.NewObjectID(), Title: "test team", Author: limitedUser, CreatedAt: time.Now()}
	exerciseID := primitive.NewObjectID()
	medicalCase := model.MedicalCase{Title: "test medical case"}
	makeUpCenter := model.MakeupCenter{Title: "make up Center", Account: limitedUser}
	roleplayer := model.Roleplayer{Fullname: "role player"}

	teams := []model.ExerciseTeam{{Team: team, Trainer: limitedUser}}
	exercise := model.Exercise{ID: exerciseID, Author: limitedUser, CreatedAt: time.Now(), Teams: teams, RoleplayManager: []model.LimitedUser{limitedUser}, MakeupCenter: []model.MakeupCenter{makeUpCenter}}
	_ = storage.CreateExercise(nil, &exercise)
	user.Roles = []model.UserRole{{Exercise: &model.ExerciseShort{ID: exerciseID}}}
	_ = storage.UpdateUser(nil, user)
	// model
	inject := model.Inject{Author: limitedUser, CreatedAt: time.Now(), ExerciseID: exerciseID, MedicalCase: medicalCase, Roleplayer: roleplayer, Team: team, MakeupCenterTitle: makeUpCenter.Title}
	invalidInject := model.Inject{}
	triggerBadRequest := `"title":"title"}`
	injectsToCreate := []*model.Inject{&inject}
	invalidInjectToCreate := []*model.Inject{&invalidInject}

	// tests
	t.Run("create", func(t *testing.T) {
		// create team
		rec, err := testRequest(http.MethodPost, "/api/exercises/:exercise_id/injects", strings.NewReader(structToJSONString(injectsToCreate)), HandleInjectsCreate, header, map[string]string{"exercise_id": exerciseID.Hex()}, jwtCookie)
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			response := &injectCreateResponse{}
			parseResponse(rec, response)
			assert.Equal(t, http.StatusCreated, rec.Code)
			assert.Equal(t, inject.ExerciseID, response.Data[0].ExerciseID)
			assert.NotNil(t, response.Data[0].Author.ID)
			assert.NotNil(t, response.Data[0].Author.Username)
			assert.NotNil(t, response.Data[0].CreatedAt)
			assert.False(t, response.Data[0].ID.IsZero())

			jsonString, _ := json.Marshal(response.Data)
			json.Unmarshal(jsonString, &injectsToCreate)
		}

		// no payload error
		rec, err = testRequest(http.MethodPost, "/api/exercises/:exercise_id/injects", nil, HandleInjectsCreate, header, map[string]string{"exercise_id": exerciseID.Hex()}, jwtCookie)
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			response := &injectCreateResponse{}
			parseResponse(rec, response)
			assert.Equal(t, 0, len(response.Data))
		}

		// bad request
		rec, err = testRequest(http.MethodPost, "/api/exercises/:exercise_id/injects", strings.NewReader(triggerBadRequest), HandleInjectsCreate, header, map[string]string{"exercise_id": exerciseID.Hex()}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
			assert.Equal(t, errorParseRequest, err.(*echo.HTTPError).Message)
		}

		// no excercise ID
		rec, err = testRequest(http.MethodPost, "/api/exercises/:exercise_id/injects", strings.NewReader(structToJSONString(injectsToCreate)), HandleInjectsCreate, header, nil, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
			assert.Equal(t, errorNoIDParam, err.(*echo.HTTPError).Message)
		}

		// wrong excercise ID
		rec, err = testRequest(http.MethodPost, "/api/exercises/:exercise_id/injects", strings.NewReader(structToJSONString(injectsToCreate)), HandleInjectsCreate, header, map[string]string{"exercise_id": primitive.NewObjectID().Hex()}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusUnauthorized, err.(*echo.HTTPError).Code)
		}

		//not create invalid
		rec, err = testRequest(http.MethodPost, "/api/exercises/:exercise_id/injects", strings.NewReader(structToJSONString(invalidInjectToCreate)), HandleInjectsCreate, header, map[string]string{"exercise_id": exerciseID.Hex()}, jwtCookie)
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			response := &injectCreateResponse{}
			parseResponse(rec, response)
			assert.Equal(t, 0, len(response.Data))
		}
	})

	t.Run("update", func(t *testing.T) {
		// setup
		inject.ExerciseID = primitive.NewObjectID()

		invalidInject := inject
		invalidInject.CreatedAt = time.Time{}

		id := primitive.NewObjectID()
		falseInject := inject
		falseInject.ID = id

		// requests
		rec, err := testRequest(http.MethodPut, "/api/exercises/:exercise_id/injects/:id", strings.NewReader(structToJSONString(inject)), HandleInjectEdit, header, map[string]string{"id": inject.ID.Hex(), "exercise_id": exerciseID.Hex()}, jwtCookie)
		if assert.NoError(t, err) {
			response := &model.Inject{}
			parseResponse(rec, response)
			assert.Equal(t, inject.ExerciseID, response.ExerciseID)
			assert.NotNil(t, response.EditedAt)
			assert.NotNil(t, response.Editor.ID)
			assert.NotNil(t, response.Editor.Username)
		}

		// no team provided
		_, err = testRequest(http.MethodPut, "/api/exercises/:exercise_id/injects/:id", nil, HandleInjectEdit, header, map[string]string{"id": inject.ID.Hex(), "exercise_id": exerciseID.Hex()}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, errorParseRequest, err.(*echo.HTTPError).Message)
		}

		// bad request
		rec, err = testRequest(http.MethodPut, "/api/exercises/:exercise_id/injects/:id", strings.NewReader(triggerBadRequest), HandleInjectEdit, header, map[string]string{"id": inject.ID.Hex(), "exercise_id": exerciseID.Hex()}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
			assert.Equal(t, errorParseRequest, err.(*echo.HTTPError).Message)
		}

		// no id provided
		_, err = testRequest(http.MethodPut, "/api/exercises/:exercise_id/injects/:id", strings.NewReader(structToJSONString(inject)), HandleInjectEdit, header, map[string]string{"id": "", "exercise_id": exerciseID.Hex()}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, errorNoIDParam, err.(*echo.HTTPError).Message)
		}

		// false id
		_, err = testRequest(http.MethodPut, "/api/exercises/:exercise_id/injects/:id", strings.NewReader(structToJSONString(inject)), HandleInjectEdit, header, map[string]string{"id": primitive.NewObjectID().Hex(), "exercise_id": exerciseID.Hex()}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, errorIDNotMatch, err.(*echo.HTTPError).Message)
		}

		// no excercise id
		rec, err = testRequest(http.MethodPut, "/api/exercises/:exercise_id/injects/:id", strings.NewReader(structToJSONString(inject)), HandleInjectEdit, header, map[string]string{"id": inject.ID.Hex(), "exercise_id": ""}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
			assert.Equal(t, errorNoIDParam, err.(*echo.HTTPError).Message)
		}

		// false exercise id
		rec, err = testRequest(http.MethodPut, "/api/exercises/:exercise_id/injects/:id", strings.NewReader(structToJSONString(inject)), HandleInjectEdit, header, map[string]string{"id": inject.ID.Hex(), "exercise_id": primitive.NewObjectID().Hex()}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusUnauthorized, err.(*echo.HTTPError).Code)
		}

		// Invalid
		_, err = testRequest(http.MethodPut, "/api/exercises/:exercise_id/injects/:id", strings.NewReader(structToJSONString(invalidInject)), HandleInjectEdit, header, map[string]string{"id": inject.ID.Hex(), "exercise_id": exerciseID.Hex()}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, "created at not set", err.(*echo.HTTPError).Message)
		}
		// Could not update
		_, err = testRequest(http.MethodPut, "/api/exercises/:exercise_id/injects/:id", strings.NewReader(structToJSONString(falseInject)), HandleInjectEdit, header, map[string]string{"id": id.Hex(), "exercise_id": exerciseID.Hex()}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, errorUpdated, err.(*echo.HTTPError).Message)
		}

	})

	t.Run("find", func(t *testing.T) {
		// no id provided
		_, err := testRequest(http.MethodGet, "/api/exercises/:exercise_id/injects/:id", nil, HandleInjectFind, nil, map[string]string{"id": "", "exercise_id": exerciseID.Hex()}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, errorNoIDParam, err.(*echo.HTTPError).Message)
		}
		// false id
		_, err = testRequest(http.MethodGet, "/api/exercises/:exercise_id/injects/:id", nil, HandleInjectFind, nil, map[string]string{"id": primitive.NewObjectID().Hex(), "exercise_id": exerciseID.Hex()}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, errorFind, err.(*echo.HTTPError).Message)
		}
		// no exercise id
		_, err = testRequest(http.MethodGet, "/api/exercises/:exercise_id/injects/:id", nil, HandleInjectFind, nil, map[string]string{"id": inject.ID.Hex(), "exercise_id": ""}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
			assert.Equal(t, errorNoIDParam, err.(*echo.HTTPError).Message)
		}

		// false exercise id
		_, err = testRequest(http.MethodGet, "/api/exercises/:exercise_id/injects/:id", nil, HandleInjectFind, nil, map[string]string{"id": inject.ID.Hex(), "exercise_id": primitive.NewObjectID().Hex()}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusUnauthorized, err.(*echo.HTTPError).Code)
		}

		// good id
		rec, err := testRequest(http.MethodGet, "/api/exercises/:exercise_id/injects/:id", nil, HandleInjectFind, nil, map[string]string{"id": inject.ID.Hex(), "exercise_id": exerciseID.Hex()}, jwtCookie)
		if assert.NoError(t, err) {
			response := &model.Inject{}
			parseResponse(rec, response)
			assert.Equal(t, inject.ExerciseID, response.ExerciseID)
		}
	})

	t.Run("get", func(t *testing.T) {
		q := make(url.Values)
		rec, err := testRequest(http.MethodGet, "/api/exercises/:exercise_id/injects?"+q.Encode(), nil, HandleInjectsGet, nil, map[string]string{"exercise_id": exerciseID.Hex()}, jwtCookie)
		if assert.NoError(t, err) {
			response := &model.TeamsList{}
			parseResponse(rec, response)
			assert.Greater(t, response.Count, int64(0))
		}
		q.Add("exercise_id", inject.ExerciseID.Hex())
		rec, err = testRequest(http.MethodGet, "/api/exercises/:exercise_id/injects?"+q.Encode(), nil, HandleInjectsGet, nil, map[string]string{"exercise_id": exerciseID.Hex()}, jwtCookie)
		if assert.NoError(t, err) {
			response := &model.InjectList{}
			parseResponse(rec, response)
			assert.Greater(t, response.Count, int64(0))
			assert.Equal(t, inject.ExerciseID, response.Injects[0].ExerciseID)
		}
		q.Add("page", "1")
		q.Add("limit", "1")
		rec, err = testRequest(http.MethodGet, "/api/exercises/:exercise_id/injects?"+q.Encode(), nil, HandleInjectsGet, nil, map[string]string{"exercise_id": exerciseID.Hex()}, jwtCookie)
		if assert.NoError(t, err) {
			response := &model.InjectList{}
			parseResponse(rec, response)
			assert.Equal(t, 1, len(response.Injects))
			assert.Equal(t, inject.ExerciseID, response.Injects[0].ExerciseID)
		}
		q.Del("page")
		q.Del("limit")
		q.Add("author", inject.Author.Username)
		rec, err = testRequest(http.MethodGet, "/api/exercises/:exercise_id/injects?"+q.Encode(), nil, HandleInjectsGet, nil, map[string]string{"exercise_id": exerciseID.Hex()}, jwtCookie)
		if assert.NoError(t, err) {
			response := &model.InjectList{}
			parseResponse(rec, response)
			assert.Greater(t, response.Count, int64(0))
			assert.Equal(t, inject.ExerciseID, response.Injects[0].ExerciseID)
			assert.Equal(t, inject.Author.Username, response.Injects[0].Author.Username)
		}

		// no exercise id
		rec, err = testRequest(http.MethodGet, "/api/exercises/:exercise_id/injects?"+q.Encode(), nil, HandleInjectsGet, nil, map[string]string{"exercise_id": ""}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
			assert.Equal(t, errorNoIDParam, err.(*echo.HTTPError).Message)
		}

		// false exercise id
		rec, err = testRequest(http.MethodGet, "/api/exercises/:exercise_id/injects?"+q.Encode(), nil, HandleInjectsGet, nil, map[string]string{"exercise_id": primitive.NewObjectID().Hex()}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusUnauthorized, err.(*echo.HTTPError).Code)
		}
	})

	t.Run("delete", func(t *testing.T) {
		// no id provided
		_, err := testRequest(http.MethodDelete, "/api/exercises/:exercise_id/injects/:id", nil, HandleInjectDelete, nil, map[string]string{"id": "", "exercise_id": exerciseID.Hex()}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, errorNoIDParam, err.(*echo.HTTPError).Message)
		}
		// false id
		_, err = testRequest(http.MethodDelete, "/api/exercises/:exercise_id/injects/:id", nil, HandleInjectDelete, nil, map[string]string{"id": primitive.NewObjectID().Hex(), "exercise_id": exerciseID.Hex()}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, errorDelete, err.(*echo.HTTPError).Message)
		}
		// no exercise id
		_, err = testRequest(http.MethodDelete, "/api/exercises/:exercise_id/injects/:id", nil, HandleInjectDelete, nil, map[string]string{"id": inject.ID.Hex(), "exercise_id": ""}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
			assert.Equal(t, errorNoIDParam, err.(*echo.HTTPError).Message)
		}
		// false exercise id
		_, err = testRequest(http.MethodDelete, "/api/exercises/:exercise_id/injects/:id", nil, HandleInjectDelete, nil, map[string]string{"id": inject.ID.Hex(), "exercise_id": primitive.NewObjectID().Hex()}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusUnauthorized, err.(*echo.HTTPError).Code)
		}
		// good id
		_, err = testRequest(http.MethodDelete, "/api/exercises/:exercise_id/injects/:id", nil, HandleInjectDelete, nil, map[string]string{"id": inject.ID.Hex(), "exercise_id": exerciseID.Hex()}, jwtCookie)
		assert.NoError(t, err)
	})

	// cleanup
	resetInject(&inject)
	resetExercise(&exercise)
}

func resetInject(inject *model.Inject) {
	_, _ = storage.MongoSession.Collection("injects").DeleteMany(nil, bson.M{"_id": inject.ID})
}
