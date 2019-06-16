package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	e := server()

	//Check routes are existing
	req := httptest.NewRequest(http.MethodGet, "/api/healthcheck", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, handlerHealthCheck(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "Application is running.", rec.Body.String())
	}

}
