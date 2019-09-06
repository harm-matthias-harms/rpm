package handler

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	rec, err := testRequest(http.MethodGet, "/api/healthcheck", nil, handlerHealthCheck, nil, nil)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "Application is running.", rec.Body.String())
	}

}

func testRequest(method string, url string, body io.Reader, handler func(c echo.Context) error, header http.Header, pathParams map[string]string, cookies ...*http.Cookie) (*httptest.ResponseRecorder, error) {
	e, _ := Server()
	req := httptest.NewRequest(method, "/", body)
	req.Header = header
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(url)
	paramNames := []string{}
	paramValues := []string{}
	for key, value := range pathParams {
		paramNames = append(paramNames, key)
		paramValues = append(paramValues, value)
	}
	c.SetParamNames(paramNames...)
	c.SetParamValues(paramValues...)
	err := handler(c)
	return rec, err
}
