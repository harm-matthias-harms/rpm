package main

import (
	"os"
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	e := server()

	//Check routes are existing
	req := httptest.NewRequest(http.MethodGet, "/healthcheck", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, handlerHealthCheck(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "Application is running.", rec.Body.String())
	}

}

func TestMongoRunning(t *testing.T) {
	mongo := setMongoDatabase()
	assert.NotEqual(t, nil, mongo)
}

func TestGetEnv(t *testing.T) {
	assert.Equal(t, "test", getEnv("SOME_RANDOM_ENV", "test"))
	os.Setenv("SOME_RANDOM_ENV", "1")
	assert.Equal(t, "1", getEnv("SOME_RANDOM_ENV", "test"))
}