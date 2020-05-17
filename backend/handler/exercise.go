package handler

import (
	"context"
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

	if err = addExercisePermissions(c.Request().Context(), exercise); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errorParseRequest)
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

	oldExerciseVersion, err := storage.FindExercise(c.Request().Context(), exercise.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errorFind)
	}

	if err = storage.UpdateExercise(c.Request().Context(), exercise, objectID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errorUpdated)
	}

	if err = revokeExercisePermissions(c.Request().Context(), oldExerciseVersion); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errorParseRequest)
	}

	if err = addExercisePermissions(c.Request().Context(), exercise); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errorParseRequest)
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

	// needs to load exercise to revoke permissions
	exercise, err := storage.FindExercise(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errorFind)
	}

	count, err := storage.DeleteExercise(c.Request().Context(), id, userID)
	if err != nil || count == int64(0) {
		return echo.NewHTTPError(http.StatusBadRequest, errorDelete)
	}

	err = revokeExercisePermissions(c.Request().Context(), exercise)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, jsonStatus{Success: true})
}

func addExercisePermissions(ctx context.Context, exercise *model.Exercise) (err error) {
	err = addExerciseToUser(ctx, exercise, exercise.Author, model.ExerciseAdminRole)
	if err != nil {
		return
	}
	for _, team := range exercise.Teams {
		err = addExerciseToUser(ctx, exercise, team.Trainer, model.TrainerRole)
		if err != nil {
			return
		}
	}
	for _, rpm := range exercise.RoleplayManager {
		err = addExerciseToUser(ctx, exercise, rpm, model.RolePlayManagerRole)
		if err != nil {
			return
		}
	}
	for _, mc := range exercise.MakeupCenter {
		err = addExerciseToUser(ctx, exercise, mc.Account, model.MakeUpCenterRole)
		if err != nil {
			return
		}
	}
	return
}

func revokeExercisePermissions(ctx context.Context, exercise *model.Exercise) (err error) {
	err = removeExerciseFromUser(ctx, exercise, exercise.Author, model.ExerciseAdminRole)
	if err != nil {
		return
	}
	for _, team := range exercise.Teams {
		err = removeExerciseFromUser(ctx, exercise, team.Trainer, model.TrainerRole)
		if err != nil {
			return
		}
	}
	for _, rpm := range exercise.RoleplayManager {
		err = removeExerciseFromUser(ctx, exercise, rpm, model.RolePlayManagerRole)
		if err != nil {
			return
		}
	}
	for _, mc := range exercise.MakeupCenter {
		err = removeExerciseFromUser(ctx, exercise, mc.Account, model.MakeUpCenterRole)
		if err != nil {
			return
		}
	}
	return
}

// TODO delete user if no more Role exists?
func removeExerciseFromUser(ctx context.Context, exercise *model.Exercise, user model.LimitedUser, role string) (err error) {
	dbUser := new(model.User)
	utils.Convert(user, dbUser)
	dbUser, err = storage.FindUser(ctx, dbUser)
	if err != nil {
		return
	}
	for i := 0; i < len(dbUser.Roles); i++ {
		if dbUser.Roles[i].Exercise.ID == exercise.ID && dbUser.Roles[i].Role == role {
			dbUser.Roles = append(dbUser.Roles[:i], dbUser.Roles[i+1:]...)
			i--
		}
	}
	err = storage.UpdateUser(ctx, dbUser)
	if dbUser.Code != "" && len(dbUser.Roles) == 0 {
		storage.DeleteUser(ctx, dbUser)
	}
	return
}

func addExerciseToUser(ctx context.Context, exercise *model.Exercise, user model.LimitedUser, role string) (err error) {
	fullUser := new(model.User)
	utils.Convert(user, fullUser)
	if fullUser.ID.IsZero() {
		err = storage.CreateUser(ctx, fullUser)
		if err != nil {
			return
		}
	}
	fullUser, err = storage.FindUser(ctx, fullUser)
	if err != nil {
		return
	}

	shortExercise := new(model.ExerciseShort)
	utils.Convert(exercise, shortExercise)
	userRole := model.UserRole{Role: role, Exercise: shortExercise}
	fullUser.Roles = append(fullUser.Roles, userRole)
	err = storage.UpdateUser(ctx, fullUser)
	return
}
