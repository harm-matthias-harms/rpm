package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/harm-matthias-harms/rpm/backend/model"
	"github.com/harm-matthias-harms/rpm/backend/storage"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestUserRegisterHandler(t *testing.T) {
	userString := `{"username":"username","email":"user@mail.com","password":"abc123" }`
	wrongUserString := `{"user1name":"username","email1":"user@mail.com","password1":"abc123" }`
	resetUserDatabase()
	e, _ := Server()

	//Throw error if no data is provided
	req := httptest.NewRequest(http.MethodPost, "/auth/register", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := HandleRegister(c)
	if assert.Error(t, err) {
		err, ok := err.(*echo.HTTPError)
		if ok {
			assert.Equal(t, http.StatusBadRequest, err.Code)
			assert.Equal(t, "couldn't parse request", err.Message)
		}
	}

	// Register user
	req = httptest.NewRequest(http.MethodPost, "/auth/register", strings.NewReader(userString))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	err = HandleRegister(c)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, `{"success":true}`, strings.TrimSpace(rec.Body.String()))
	}

	// Not register wrong user
	//Throw error if no data is provided
	req = httptest.NewRequest(http.MethodPost, "/auth/register", strings.NewReader(wrongUserString))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	err = HandleRegister(c)
	if assert.Error(t, err) {
		err, ok := err.(*echo.HTTPError)
		if ok {
			assert.Equal(t, http.StatusBadRequest, err.Code)
			assert.Equal(t, "no username provided", err.Message)
		}
	}
}

func TestUserAuthenticateHandler(t *testing.T) {
	userString := `{"username":"username", "password":"abc123" }`
	wrongUserString := `{"username":"username","password1":"abc1234" }`
	resetUserDatabase()
	_ = register(&model.User{Username: "username", Email: "test@mail.com", Password: "abc123"})
	e, _ := Server()

	//Throw error if no data is provided
	req := httptest.NewRequest(http.MethodPost, "/auth/authenticate", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := HandleAuthenticate(c)
	if assert.Error(t, err) {
		err, ok := err.(*echo.HTTPError)
		if ok {
			assert.Equal(t, http.StatusBadRequest, err.Code)
			assert.Equal(t, "couldn't parse request", err.Message)
		}
	}

	// Authenticate user
	req = httptest.NewRequest(http.MethodPost, "/auth/authenticate", strings.NewReader(userString))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	err = HandleAuthenticate(c)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, `{"success":true}`, strings.TrimSpace(rec.Body.String()))
		assert.NotEqual(t, []*http.Cookie{}, rec.Result().Cookies())
	}

	// Don't authenticate wrong password
	req = httptest.NewRequest(http.MethodPost, "/auth/authenticate", strings.NewReader(wrongUserString))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	err = HandleAuthenticate(c)
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
	resetUserDatabase()
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
	for _, tt := range tests {
		err := register(tt.User)
		if tt.Err {
			assert.Error(t, err)
		}
	}
}

func TestUserAuthenticate(t *testing.T) {
	resetUserDatabase()
	user := &model.User{Username: "testperson", Email: "test@mail.com", Password: "123456"}
	wrongUsername := &model.User{Username: "testperso", Password: "123456"}
	wrongPassword := &model.User{Username: "testperson", Password: "1234567"}
	_ = register(&model.User{Username: "testperson", Email: "test@mail.com", Password: "123456"})
	// auth good user
	_, err := authenticate(user)
	assert.NoError(t, err)
	// not auth none present user
	_, err = authenticate(wrongUsername)
	assert.Error(t, err)
	// not auth wrong password
	_, err = authenticate(wrongPassword)
	assert.Error(t, err)
}

func resetUserDatabase() {
	_ = storage.SetMongoDatabase()
	_ = storage.MongoSession.Collection("user").Drop(context.Background())
	storage.CreateUserIndexes()
}
