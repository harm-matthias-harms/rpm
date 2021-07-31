package handler

import (
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/harm-matthias-harms/rpm/backend/model"
	"github.com/harm-matthias-harms/rpm/backend/storage"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUser(t *testing.T) {
	// setup
	_ = storage.SetMongoDatabase()
	header := jsonHeader()
	jwtCookie := &http.Cookie{}

	// models
	userString := `{"username":"username","email":"user@mail.com","password":"abc123","code":"code123"}`
	triggerParseError := `{"user1name":"username","email1":"user@mail.com","password1":"abc123" }`
	triggerBadRequest := `"username":"username","password":"abc123" }`
	wrongPassword := `{"username":"username","email":"user@mail.com","password":"abc1234" }`
	codeSignIn := `{"code":"code123"}`
	codeSignInFail := `{"code":"code1234"}`

	// tests
	t.Run("register handler", func(t *testing.T) {
		//Throw error if no data is provided
		rec, err := testRequest(http.MethodPost, "/auth/register", nil, HandleRegister, header, nil)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
			assert.Error(t, err)
		}

		rec, err = testRequest(http.MethodPost, "/auth/register", strings.NewReader(triggerBadRequest), HandleRegister, header, nil)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
			assert.Equal(t, errorParseRequest, err.(*echo.HTTPError).Message)
		}

		// Register user
		rec, err = testRequest(http.MethodPost, "/auth/register", strings.NewReader(userString), HandleRegister, header, nil)
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			assert.Equal(t, `{"success":true}`, strings.TrimSpace(rec.Body.String()))
		}

		// Doesn't create two times
		rec, err = testRequest(http.MethodPost, "/auth/register", strings.NewReader(userString), HandleRegister, header, nil)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
			assert.Equal(t, "user already exists", err.(*echo.HTTPError).Message)
		}

		// Not register wrong user
		//Throw error if no data is provided
		rec, err = testRequest(http.MethodPost, "/auth/register", strings.NewReader(triggerParseError), HandleRegister, header, nil)
		if assert.Error(t, err) {
			err, ok := err.(*echo.HTTPError)
			if ok {
				assert.Equal(t, http.StatusBadRequest, err.Code)
				assert.Equal(t, "no username provided", err.Message)
			}
		}
	})

	t.Run("authenticate handler", func(t *testing.T) {

		// no payload
		rec, err := testRequest(http.MethodPost, "/auth/authenticate", nil, HandleAuthenticate, header, nil)
		if assert.Error(t, err) {
			err, ok := err.(*echo.HTTPError)
			if ok {
				assert.Equal(t, http.StatusUnauthorized, err.Code)
				assert.Error(t, err)
			}
		}

		// parse error
		rec, err = testRequest(http.MethodPost, "/auth/authenticate", strings.NewReader(triggerBadRequest), HandleAuthenticate, header, nil)
		if assert.Error(t, err) {
			err, ok := err.(*echo.HTTPError)
			if ok {
				assert.Equal(t, http.StatusBadRequest, err.Code)
				assert.Equal(t, errorParseRequest, err.Message)
			}
		}

		// Authenticate user
		rec, err = testRequest(http.MethodPost, "/auth/authenticate", strings.NewReader(userString), HandleAuthenticate, header, nil)
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, `{"success":true}`, strings.TrimSpace(rec.Body.String()))
			assert.NotEqual(t, []*http.Cookie{}, rec.Result().Cookies())
		}
		cookies := rec.Result().Cookies()
		for _, cookie := range cookies {
			if cookie.Name == echo.HeaderAuthorization {
				jwtCookie = cookie
			}
		}

		// Don't authenticate wrong password
		rec, err = testRequest(http.MethodPost, "/auth/authenticate", strings.NewReader(wrongPassword), HandleAuthenticate, header, nil)
		if assert.Error(t, err) {
			err, ok := err.(*echo.HTTPError)
			if ok {
				assert.Equal(t, echo.ErrUnauthorized.Code, err.Code)
				assert.Equal(t, echo.ErrUnauthorized.Message, err.Message)
				assert.Equal(t, []*http.Cookie{}, rec.Result().Cookies())
			}
		}

		// Authenticate code
		rec, err = testRequest(http.MethodPost, "/auth/authenticate", strings.NewReader(codeSignIn), HandleAuthenticate, header, nil)
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, `{"success":true}`, strings.TrimSpace(rec.Body.String()))
			assert.NotEqual(t, []*http.Cookie{}, rec.Result().Cookies())
		}

		// Wrong code
		rec, err = testRequest(http.MethodPost, "/auth/authenticate", strings.NewReader(codeSignInFail), HandleAuthenticate, header, nil)
		if assert.Error(t, err) {
			err, ok := err.(*echo.HTTPError)
			if ok {
				assert.Equal(t, echo.ErrUnauthorized.Code, err.Code)
				assert.Equal(t, echo.ErrUnauthorized.Message, err.Message)
				assert.Equal(t, []*http.Cookie{}, rec.Result().Cookies())
			}
		}
	})

	t.Run("find", func(t *testing.T) {
		user, _ := storage.FindUser(nil, &model.User{Username: "username"})
		// no id provided
		_, err := testRequest(http.MethodGet, "/api/user/:id", nil, HandleUserFind, nil, map[string]string{"id": ""}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, errorNoIDParam, err.(*echo.HTTPError).Message)
		}
		// false id
		_, err = testRequest(http.MethodGet, "/api/user/:id", nil, HandleUserFind, nil, map[string]string{"id": primitive.NewObjectID().Hex()}, jwtCookie)
		if assert.Error(t, err) {
			assert.Equal(t, errorNotAuthorized, err.(*echo.HTTPError).Message)
		}
		// good id
		rec, err := testRequest(http.MethodGet, "/api/user/:id", nil, HandleUserFind, nil, map[string]string{"id": user.ID.Hex()}, jwtCookie)
		if assert.NoError(t, err) {
			response := &model.User{}
			parseResponse(rec, response)
			assert.Equal(t, user.ID, response.ID)
			assert.Equal(t, "", response.Password)
		}
	})

	t.Run("get", func(t *testing.T) {
		user, _ := storage.FindUser(nil, &model.User{Username: "username"})

		q := make(url.Values)
		rec, err := testRequest(http.MethodGet, "/api/user?"+q.Encode(), nil, HandleUserGet, nil, nil, jwtCookie)
		if assert.NoError(t, err) {
			var response []model.LimitedUser
			parseResponse(rec, &response)
			assert.Greater(t, len(response), 0)
		}
		q.Add("username", user.Username)
		rec, err = testRequest(http.MethodGet, "/api/user?"+q.Encode(), nil, HandleUserGet, nil, nil, jwtCookie)
		if assert.NoError(t, err) {
			var response []model.LimitedUser
			parseResponse(rec, &response)
			assert.Greater(t, len(response), 0)
			assert.Equal(t, user.Username, response[0].Username)
		}
		q.Add("page", "1")
		q.Add("limit", "1")
		rec, err = testRequest(http.MethodGet, "/api/user?"+q.Encode(), nil, HandleUserGet, nil, nil, jwtCookie)
		if assert.NoError(t, err) {
			var response []model.LimitedUser
			parseResponse(rec, &response)
			assert.Equal(t, 1, len(response))
			assert.Equal(t, user.Username, response[0].Username)
		}
		q.Del("page")
		q.Del("limit")
		q.Add("email", user.Email)
		rec, err = testRequest(http.MethodGet, "/api/user?"+q.Encode(), nil, HandleUserGet, nil, nil, jwtCookie)
		if assert.NoError(t, err) {
			var response []model.LimitedUser
			parseResponse(rec, &response)
			assert.Greater(t, len(response), 0)
			assert.Equal(t, user.Username, response[0].Username)
			assert.Equal(t, user.Email, response[0].Email)
		}
	})

	// cleanup
	resetUserDatabase(&model.User{Username: "username"})
}

func resetUserDatabase(user *model.User) {
	_, _ = storage.MongoSession.Collection("users").DeleteMany(nil, bson.M{"username": user.Username})
}
