package handler

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/harm-matthias-harms/rpm/backend/model"
	"github.com/harm-matthias-harms/rpm/backend/storage"
	"github.com/harm-matthias-harms/rpm/backend/utils"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// HandlePresetGet gives back a number of presets
func HandlePresetGet(c echo.Context) (err error) {
	return nil
}

// HandlePresetCreate creates a preset from a json
func HandlePresetCreate(c echo.Context) (err error) {
	preset := new(model.Preset)
	if err := c.Bind(preset); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "couldn't parse request")
	}
	cookie, _ := c.Cookie(echo.HeaderAuthorization)
	token, _ := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(utils.GetEnv("JWT_SECRET", "secret")), nil
	})
	claims := token.Claims.(jwt.MapClaims)
	objectID, err := primitive.ObjectIDFromHex(claims["id"].(string))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "authorization couldn't be parsed")
	}
	preset.Author.ID = objectID
	preset.Author.Username = claims["username"].(string)
	preset.CreatedAt = time.Now()
	if err = preset.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = storage.CreatePreset(c.Request().Context(), preset); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "couldn't be created")
	}
	return c.JSON(http.StatusCreated, jsonStatus{Success: true, Data: preset})
}
