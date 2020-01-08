package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/harm-matthias-harms/rpm/backend/model"
	"github.com/harm-matthias-harms/rpm/backend/storage"
	"github.com/harm-matthias-harms/rpm/backend/utils"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// HandlePresetsGet gives back a number of presets
func HandlePresetsGet(c echo.Context) (err error) {
	params := new(model.PresetQuery)
	if err = c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "params couldn't be parsed")
	}
	filter := map[string]interface{}{}
	if params.Title != "" {
		filter["title"] = primitive.Regex{Pattern: params.Title}
	}
	if params.Author != "" {
		filter["author.username"] = primitive.Regex{Pattern: params.Author}
	}
	presets, err := storage.GetPresets(c.Request().Context(), filter, params.Page, params.PageSize)
	count, err := storage.CountPresets(c.Request().Context(), filter)
	response := model.PresetsList{Count: count, Presets: presets}
	return c.JSON(http.StatusOK, response)
}

// HandlePresetFind returns a preset
func HandlePresetFind(c echo.Context) (err error) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "no or false id provided")
	}
	preset, err := storage.FindPreset(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "couldn't find preset")
	}
	return c.JSON(http.StatusOK, preset)
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

// HandlePresetEdit updates an preset
func HandlePresetEdit(c echo.Context) (err error) {
	preset := new(model.Preset)
	defer c.Request().Body.Close()
	body, err := ioutil.ReadAll(c.Request().Body)
	err = json.Unmarshal(body, preset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "couldn't parse request")
	}
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "no or false id provided")
	}
	if preset.ID != id {
		return echo.NewHTTPError(http.StatusBadRequest, "id's do not match")
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

	preset.Editor.ID = objectID
	preset.Editor.Username = claims["username"].(string)
	preset.EditedAt = time.Now()
	if err = preset.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = storage.UpdatePreset(c.Request().Context(), preset); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "couldn't be updated")
	}

	return c.JSON(http.StatusOK, preset)
}
