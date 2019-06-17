package handler

import (
	"net/http"

	"github.com/harm-matthias-harms/rpm/backend/model"
	"github.com/labstack/echo"
)

// HandleRegister provides the registration endpoint for the api
func HandleRegister(c echo.Context) (err error) {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "couldn't parse request")
	}
	if err = user.Register(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	return c.JSON(http.StatusCreated, jsonStatus{Success: true})
}
