package handler

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/harm-matthias-harms/rpm/backend/model"
	"github.com/harm-matthias-harms/rpm/backend/storage"
	"github.com/harm-matthias-harms/rpm/backend/utils"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
)

type injectCreateResponse struct {
	Success bool           `json:"status"`
	Data    []model.Inject `json:"data"`
}

// HandleInjectsGet gives back a list of injects
func HandleInjectsGet(c echo.Context) (err error) {
	params := new(model.InjectQuery)
	if err = c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errorParseParams)
	}
	filter := map[string]interface{}{}
	if !params.ExerciseID.IsZero() {
		filter["exerciseId"] = params.ExerciseID
	}
	if !params.Team.ID.IsZero() {
		filter["team._id"] = params.Team.ID
	}
	if params.MakeupCenterTitle != "" {
		filter["makeupCenterTitle"] = params.MakeupCenterTitle
	}
	if params.Author != "" {
		filter["author.username"] = primitive.Regex{Pattern: params.Author}
	}
	injects, err := storage.GetInjects(c.Request().Context(), filter, params.Page, params.PageSize)
	count, err := storage.CountInjects(c.Request().Context(), filter)
	response := model.InjectList{Count: count, Injects: injects}
	return c.JSON(http.StatusOK, response)
}

// HandleInjectFind returns an inject
func HandleInjectFind(c echo.Context) (err error) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errorNoIDParam)
	}
	inject, err := storage.FindInject(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errorFind)
	}
	return c.JSON(http.StatusOK, inject)
}

// HandleInjectsCreate creates injects from a list of injects!
func HandleInjectsCreate(c echo.Context) (err error) {
	var injects []model.Inject
	if err := c.Bind(&injects); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errorParseRequest)
	}
	cookie, _ := c.Cookie(echo.HeaderAuthorization)
	token, _ := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		return []byte(utils.GetEnv("JWT_SECRET", "secret")), nil
	})
	claims := token.Claims.(jwt.MapClaims)
	authorID, err := primitive.ObjectIDFromHex(claims["id"].(string))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errorAuthParse)
	}
	// do not throw errors if a single inject could not be created.
	for i := 0; i < len(injects); i++ {
		injects[i].Author.ID = authorID
		injects[i].Author.Username = claims["username"].(string)
		injects[i].CreatedAt = time.Now()
		if err = injects[i].Validate(); err != nil {
			injects = append(injects[:i], injects[i+1:]...)
			i--
			continue
		}
		if err = storage.CreateInject(c.Request().Context(), &injects[i]); err != nil {
			injects = append(injects[:i], injects[i+1:]...)
			i--
			continue
		}
	}
	return c.JSON(http.StatusCreated, injectCreateResponse{Success: true, Data: injects})
}

// HandleInjectEdit updates an inject
func HandleInjectEdit(c echo.Context) (err error) {
	inject := new(model.Inject)
	// c.Bind() is not working here
	defer c.Request().Body.Close()
	body, err := ioutil.ReadAll(c.Request().Body)
	err = json.Unmarshal(body, inject)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errorParseRequest)
	}
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errorNoIDParam)
	}
	if inject.ID != id {
		return echo.NewHTTPError(http.StatusBadRequest, errorIDNotMatch)
	}

	cookie, _ := c.Cookie(echo.HeaderAuthorization)
	token, _ := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		return []byte(utils.GetEnv("JWT_SECRET", "secret")), nil
	})
	claims := token.Claims.(jwt.MapClaims)
	editorID, err := primitive.ObjectIDFromHex(claims["id"].(string))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errorAuthParse)
	}

	inject.Editor.ID = editorID
	inject.Editor.Username = claims["username"].(string)
	inject.EditedAt = time.Now()
	if err = inject.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = storage.UpdateInject(c.Request().Context(), inject); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errorUpdated)
	}

	return c.JSON(http.StatusOK, inject)
}

// HandleInjectDelete deletes an inject
func HandleInjectDelete(c echo.Context) (err error) {
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

	count, err := storage.DeleteInject(c.Request().Context(), id, userID)
	if err != nil || count == int64(0) {
		return echo.NewHTTPError(http.StatusBadRequest, errorDelete)
	}

	return c.JSON(http.StatusOK, jsonStatus{Success: true})
}

func userToAccessRights(ctx context.Context, user model.User, excerciseID primitive.ObjectID) (right accessRight, err error) {
	excercise, err := storage.FindExercise(ctx, excerciseID)
	if err != nil {
		return accessRight{}, errors.New(errorFind)
	}
	if excercise.Author.ID == user.ID {
		return accessRight{Type: "exerciseID", ID: excerciseID}, nil
	}
	for _, rpm := range excercise.RoleplayManager {
		if rpm.ID == user.ID {
			return accessRight{Type: "exerciseID", ID: excerciseID}, nil
		}
	}
	for _, makeupCenter := range excercise.MakeupCenter {
		if makeupCenter.Account.ID == user.ID {
			return accessRight{Type: "makeupCenterTitle", ID: makeupCenter.Title}, nil
		}
	}
	for _, team := range excercise.Teams {
		if team.Trainer.ID == user.ID {
			return accessRight{Type: "teamID", ID: team.Team.ID}, nil
		}
	}
	return accessRight{}, errors.New(errorFind)
}

func accessRightToParam(right accessRight, params *model.InjectQuery) (err error) {
	switch rightType := right.Type; rightType {
	case "exerciseID":
		params.ExerciseID = right.ID.(primitive.ObjectID)
		return nil
	case "makeupCenterTitle":
		params.MakeupCenterTitle = right.ID.(string)
		return nil
	case "teamID":
		params.Team.ID = right.ID.(primitive.ObjectID)
		return nil
	}
	return errors.New("unknown access right")
}

func hasAccessToInject(ctx context.Context, user model.User, inject model.Inject) (err error) {
	right, err := userToAccessRights(ctx, user, inject.ExerciseID)
	if err != nil {
		return
	}
	switch rightType := right.Type; rightType {
	case "exerciseID":
		if inject.ExerciseID == right.ID.(primitive.ObjectID) {
			return nil
		}
	case "makeupCenterTitle":
		if inject.MakeupCenterTitle == right.ID.(string) {
			return nil
		}
	case "teamID":
		if inject.Team.ID == right.ID.(primitive.ObjectID) {
			return nil
		}
	}
	return errors.New("no access")
}

type accessRight struct {
	Type string
	ID     interface{}
}
