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

// HandleExercisesGet no need to list exercises
func HandleExercisesGet(c echo.Context) (err error) {
	return echo.NewHTTPError(http.StatusNotImplemented)
}

// HandleExerciseFind returns an exercise
func HandleExerciseFind(c echo.Context) (err error) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errorNoIDParam)
	}
	exercise, err := storage.FindExercise(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errorFind)
	}
	return c.JSON(http.StatusOK, exercise)
}

// HandleExerciseCreate creates an exercise from a json
func HandleExerciseCreate(c echo.Context) (err error) {
	exercise := new(model.Exercise)
	if err := c.Bind(exercise); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errorParseRequest)
	}
	cookie, _ := c.Cookie(echo.HeaderAuthorization)
	token, _ := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		return []byte(utils.GetEnv("JWT_SECRET", "secret")), nil
	})
	claims := token.Claims.(jwt.MapClaims)
	objectID, err := primitive.ObjectIDFromHex(claims["id"].(string))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errorAuthParse)
	}
	exercise.Author.ID = objectID
	exercise.Author.Username = claims["username"].(string)
	exercise.CreatedAt = time.Now()
	if err = exercise.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = storage.CreateExercise(c.Request().Context(), exercise); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errorCreated)
	}
	return c.JSON(http.StatusCreated, jsonStatus{Success: true, Data: exercise})
}

// HandleExerciseEdit updates an exercise
func HandleExerciseEdit(c echo.Context) (err error) {
	exercise := new(model.Exercise)
	// c.Bind is not working here
	defer c.Request().Body.Close()
	body, err := ioutil.ReadAll(c.Request().Body)
	err = json.Unmarshal(body, exercise)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errorParseRequest)
	}
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errorNoIDParam)
	}
	if exercise.ID != id {
		return echo.NewHTTPError(http.StatusBadRequest, errorIDNotMatch)
	}

	cookie, _ := c.Cookie(echo.HeaderAuthorization)
	token, _ := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		return []byte(utils.GetEnv("JWT_SECRET", "secret")), nil
	})
	claims := token.Claims.(jwt.MapClaims)
	objectID, err := primitive.ObjectIDFromHex(claims["id"].(string))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errorAuthParse)
	}

	if exercise.Author.ID == objectID {
		exercise.EditedAt = time.Now()
	}
	if err = exercise.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = storage.UpdateExercise(c.Request().Context(), exercise, objectID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errorUpdated)
	}

	return c.JSON(http.StatusOK, exercise)
}

// HandleExerciseDelete updates an exercise
func HandleExerciseDelete(c echo.Context) (err error) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errorNoIDParam)
	}

	cookie, _ := c.Cookie(echo.HeaderAuthorization)
	token, _ := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		return []byte(utils.GetEnv("JWT_SECRET", "secret")), nil
	})
	claims := token.Claims.(jwt.MapClaims)
	userID, err := primitive.ObjectIDFromHex(claims["id"].(string))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errorAuthParse)
	}

	count, err := storage.DeleteExercise(c.Request().Context(), id, userID)
	if err != nil || count == int64(0) {
		return echo.NewHTTPError(http.StatusBadRequest, errorDelete)
	}

	return c.JSON(http.StatusOK, jsonStatus{Success: true})
}
