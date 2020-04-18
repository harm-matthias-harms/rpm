package handler

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/harm-matthias-harms/rpm/backend/model"
	"github.com/harm-matthias-harms/rpm/backend/storage"
	"github.com/harm-matthias-harms/rpm/backend/utils"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// HandleUserGet gives back a list of users
func HandleUserGet(c echo.Context) (err error) {
	params := new(model.UserQuery)
	if err = c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errorParseParams)
	}
	filter := map[string]interface{}{}
	if params.Username != "" {
		filter["username"] = primitive.Regex{Pattern: params.Username}
	}
	if params.Email != "" {
		filter["email"] = primitive.Regex{Pattern: params.Email}
	}
	users, err := storage.GetUser(c.Request().Context(), filter, params.Page, params.PageSize)
	return c.JSON(http.StatusOK, users)
}

// HandleRegister provides the registration endpoint for the api
func HandleRegister(c echo.Context) (err error) {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errorParseRequest)
	}
	if err = register(c.Request().Context(), user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, jsonStatus{Success: true})
}

// HandleAuthenticate describes the endpoint to sign in
func HandleAuthenticate(c echo.Context) (err error) {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errorParseRequest)
	}
	result, err := authenticate(c.Request().Context(), user)
	if err != nil {
		return echo.ErrUnauthorized
	}

	// create jwt tokem
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = result.Username
	claims["id"] = result.ID
	claims["exp"] = time.Now().Add(time.Hour * 24 * 10).Unix()
	claims["code"] = result.Code

	t, err := token.SignedString([]byte(utils.GetEnv("JWT_SECRET", "secret")))
	if err != nil {
		return err
	}

	// write in Cookie
	cookie := new(http.Cookie)
	cookie.Name = echo.HeaderAuthorization
	cookie.Value = t
	cookie.Expires = time.Now().Add(time.Hour * 24 * 10)
	cookie.Path = "/"
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, jsonStatus{Success: true})
}

func register(ctx context.Context, user *model.User) (err error) {
	if err = user.Validate(); err != nil {
		return
	}
	if err = user.HashPassword(); err != nil {
		return
	}
	if err := storage.CreateUser(ctx, user); err != nil {
		return errors.New("user already exists")
	}
	return nil
}

func authenticate(ctx context.Context, user *model.User) (result *model.User, err error) {
	result, err = storage.FindUser(ctx, user)
	if err != nil {
		return nil, err
	}
	if user.Code == "" {
		if err := result.Authenticate(user.Password); err != nil {
			return nil, err
		}
	}
	return
}
