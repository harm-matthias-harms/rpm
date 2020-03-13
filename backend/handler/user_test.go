package handler

import (
	"net/http"
	"strings"
	"testing"

	"github.com/harm-matthias-harms/rpm/backend/model"
	"github.com/harm-matthias-harms/rpm/backend/storage"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestUser(t *testing.T) {
	// setup
	_ = storage.SetMongoDatabase()
	header := jsonHeader()

	//models
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

	// cleanup
	resetUserDatabase(&model.User{Username: "username"})
}

func resetUserDatabase(user *model.User) {
	_, _ = storage.MongoSession.Collection("users").DeleteMany(nil, bson.M{"username": user.Username})
}
