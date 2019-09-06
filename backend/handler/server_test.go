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
	rec, err := testRequest(http.MethodGet, "/api/healthcheck", nil, handlerHealthCheck, nil)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "Application is running.", rec.Body.String())
	}

}

func testRequest(method string, url string, body io.Reader, handler func(c echo.Context) error, header http.Header, cookies ...*http.Cookie) (*httptest.ResponseRecorder, error) {
	e, _ := Server()
	req := httptest.NewRequest(method, url, body)
	req.Header = header
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := handler(c)
	return rec, err
}
