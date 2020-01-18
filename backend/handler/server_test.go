package handler

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/harm-matthias-harms/rpm/backend/model"
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
	if len(pathParams) == 0 {
		req = httptest.NewRequest(method, url, body)
	}
	if header != nil {
		req.Header = header
	}
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if len(pathParams) > 0 {
		c.SetPath(url)
		paramNames := []string{}
		paramValues := []string{}
		for key, value := range pathParams {
			paramNames = append(paramNames, key)
			paramValues = append(paramValues, value)
		}

		c.SetParamNames(paramNames...)
		c.SetParamValues(paramValues...)

	}
	err := handler(c)
	return rec, err
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

func jsonHeader() http.Header {
	header := http.Header{}
	header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	return header
}

func structToJSONString(v interface{}) string {
	s, _ := json.Marshal(v)
	return string(s)
}

func parseResponse(rec *httptest.ResponseRecorder, v interface{}) {
	defer rec.Result().Body.Close()
	body, _ := ioutil.ReadAll(rec.Result().Body)
	_ = json.Unmarshal(body, v)
}
