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

func TestUserRegisterHandler(t *testing.T) {
	userString := `{"username":"username","email":"user@mail.com","password":"abc123" }`
	wrongUserString := `{"user1name":"username","email1":"user@mail.com","password1":"abc123" }`
	resetUserDatabase(&model.User{Username: "username"})
	header := http.Header{}
	header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	//Throw error if no data is provided
	rec, err := testRequest(http.MethodPost, "/auth/register", nil, HandleRegister, header, nil)
	if assert.Error(t, err) {
		err, ok := err.(*echo.HTTPError)
		if ok {
			assert.Equal(t, http.StatusBadRequest, err.Code)
			assert.Equal(t, "couldn't parse request", err.Message)
		}
	}

	// Register user
	rec, err = testRequest(http.MethodPost, "/auth/register", strings.NewReader(userString), HandleRegister, header, nil)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, `{"success":true}`, strings.TrimSpace(rec.Body.String()))
	}

	// Not register wrong user
	//Throw error if no data is provided
	rec, err = testRequest(http.MethodPost, "/auth/register", strings.NewReader(wrongUserString), HandleRegister, header, nil)
	if assert.Error(t, err) {
		err, ok := err.(*echo.HTTPError)
		if ok {
			assert.Equal(t, http.StatusBadRequest, err.Code)
			assert.Equal(t, "no username provided", err.Message)
		}
	}
}

func TestUserAuthenticateHandler(t *testing.T) {
	userString := `{"username":"username1", "password":"abc123" }`
	wrongUserString := `{"username":"username1","password1":"abc1234" }`
	user := &model.User{Username: "username1", Email: "test2@mail.com", Password: "abc123"}
	resetUserDatabase(user)
	_ = register(nil, user)
	header := http.Header{}
	header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec, err := testRequest(http.MethodPost, "/auth/authenticate", nil, HandleAuthenticate, header, nil)
	if assert.Error(t, err) {
		err, ok := err.(*echo.HTTPError)
		if ok {
			assert.Equal(t, http.StatusBadRequest, err.Code)
			assert.Equal(t, "couldn't parse request", err.Message)
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
	rec, err = testRequest(http.MethodPost, "/auth/authenticate", strings.NewReader(wrongUserString), HandleAuthenticate, header, nil)
	if assert.Error(t, err) {
		err, ok := err.(*echo.HTTPError)
		if ok {
			assert.Equal(t, echo.ErrUnauthorized.Code, err.Code)
			assert.Equal(t, echo.ErrUnauthorized.Message, err.Message)
			assert.Equal(t, []*http.Cookie{}, rec.Result().Cookies())
		}
	}
}

func TestUserRegister(t *testing.T) {
	tests := []struct {
		User *model.User
		Err  bool
	}{
		{
			User: &model.User{Username: "testperson", Email: "test@mail.com", Password: "123"},
			Err:  false,
		},
		{
			User: &model.User{Email: "test@mail.com", Password: "123"},
			Err:  true,
		},
		{ // Not  register same user twice.
			User: &model.User{Username: "testperson", Email: "test@mail.com", Password: "123"},
			Err:  true,
		},
	}
	resetUserDatabase(tests[0].User)
	for _, tt := range tests {

		err := register(nil, tt.User)
		if tt.Err {
			assert.Error(t, err)
		}
	}
}

func TestUserAuthenticate(t *testing.T) {
	user := &model.User{Username: "testperson", Email: "test@mail.com", Password: "123456"}
	wrongUsername := &model.User{Username: "testperso", Password: "123456"}
	wrongPassword := &model.User{Username: "testperson", Password: "1234567"}
	resetUserDatabase(user)
	_ = register(nil, &model.User{Username: "testperson", Email: "test@mail.com", Password: "123456"})
	// auth good user
	_, err := authenticate(nil, user)
	assert.NoError(t, err)
	// not auth none present user
	_, err = authenticate(nil, wrongUsername)
	assert.Error(t, err)
	// not auth wrong password
	_, err = authenticate(nil, wrongPassword)
	assert.Error(t, err)
}

func resetUserDatabase(user *model.User) {
	_ = storage.SetMongoDatabase()
	_, _ = storage.MongoSession.Collection("user").DeleteOne(nil, bson.M{"username": user.Username})
}

// Login a user and generates a jwt. Returns nil or the JWT-Token Cookie
func loginUser(t *testing.T) *http.Cookie {
	userString := `{"username":"testuser", "password":"abc123" }`
	user := &model.User{Username: "testuser", Email: "user@test.com", Password: "abc123"}
	_ = register(nil, user)

	header := http.Header{}
	header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec, err := testRequest(http.MethodPost, "/auth/authenticate", strings.NewReader(userString), HandleAuthenticate, header, nil)
	assert.NoError(t, err)
	cookies := rec.Result().Cookies()
	for _, cookie := range cookies {
		if cookie.Name == echo.HeaderAuthorization {
			return cookie
		}
	}
	return nil
}
