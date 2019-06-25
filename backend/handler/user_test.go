package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/harm-matthias-harms/rpm/backend/model"
	"github.com/harm-matthias-harms/rpm/backend/storage"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	userString      = `{"username":"username","email":"user@mail.com","password":"abc123" }`
	wrongUserString = `{"user1name":"username","email1":"user@mail.com","password1":"abc123" }`
)

func TestUserRegisterHandler(t *testing.T) {
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
			assert.Equal(t, "invalid request", err.Message)
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
			User: &model.User{Username: "testPerson", Email: "test@mail.com", Password: "123"},
			Err:  false,
		},
		{
			User: &model.User{Email: "test@mail.com", Password: "123"},
			Err:  true,
		},
		{ // Not  register same user twice.
			User: &model.User{Username: "testPerson", Email: "test@mail.com", Password: "123"},
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

func resetUserDatabase() {
	_ = storage.SetMongoDatabase()
	_ = storage.MongoSession.Collection("user").Drop(context.Background())
	storage.CreateUserIndexes()
}
