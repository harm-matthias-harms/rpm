package handler

import (
	"net/http"

	"github.com/harm-matthias-harms/rpm/backend/model"
	"github.com/harm-matthias-harms/rpm/backend/storage"
	"github.com/labstack/echo"
)

// HandleRegister provides the registration endpoint for the api
func HandleRegister(c echo.Context) (err error) {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "couldn't parse request")
	}
	if err = register(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	return c.JSON(http.StatusCreated, jsonStatus{Success: true})
}

func register(user *model.User) error {
	if valid, err := user.Validate(); !valid {
		return err
	}
	if err := user.HashPassword(); err != nil {
		return err
	}
	if err := storage.CreateUser(user); err != nil {
		return err
	}
	return nil
}
