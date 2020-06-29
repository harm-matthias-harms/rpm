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

// HandleTeamsGet gives back a list of teams
func HandleTeamsGet(c echo.Context) (err error) {
	params := new(model.TeamQuery)
	if err = c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errorParseParams)
	}
	filter := map[string]interface{}{}
	if params.Title != "" {
		filter["title"] = primitive.Regex{Pattern: params.Title}
	}
	if params.Author != "" {
		filter["author.username"] = primitive.Regex{Pattern: params.Author}
	}
	teams, err := storage.GetTeams(c.Request().Context(), filter, params.Page, params.PageSize)
	count, err := storage.CountTeams(c.Request().Context(), filter)
	response := model.TeamsList{Count: count, Teams: teams}
	return c.JSON(http.StatusOK, response)
}

// HandleTeamFind returns a team
func HandleTeamFind(c echo.Context) (err error) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errorNoIDParam)
	}
	team, err := storage.FindTeam(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errorFind)
	}
	return c.JSON(http.StatusOK, team)
}

// HandleTeamCreate creates a team by a given json
func HandleTeamCreate(c echo.Context) (err error) {
	team := new(model.Team)
	if err := c.Bind(team); err != nil {
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
	team.Author.ID = objectID
	team.Author.Username = claims["username"].(string)
	team.CreatedAt = time.Now()
	if err = team.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = storage.CreateTeam(c.Request().Context(), team); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errorCreated)
	}
	return c.JSON(http.StatusCreated, jsonStatus{Success: true, Data: team})
}

// HandleTeamEdit updates an team
func HandleTeamEdit(c echo.Context) (err error) {
	team := new(model.Team)
	// c.Bind() is not working here
	defer c.Request().Body.Close()
	body, err := ioutil.ReadAll(c.Request().Body)
	err = json.Unmarshal(body, team)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errorParseRequest)
	}
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errorNoIDParam)
	}
	if team.ID != id {
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

	team.Editor.ID = objectID
	team.Editor.Username = claims["username"].(string)
	team.EditedAt = time.Now()
	if err = team.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = storage.UpdateTeam(c.Request().Context(), team); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errorUpdated)
	}

	return c.JSON(http.StatusOK, team)
}

// HandleTeamDelete deletes an team
func HandleTeamDelete(c echo.Context) (err error) {
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

	count, err := storage.DeleteTeam(c.Request().Context(), id, userID)
	if err != nil || count == int64(0) {
		return echo.NewHTTPError(http.StatusBadRequest, errorDelete)
	}

	return c.JSON(http.StatusOK, jsonStatus{Success: true})
}
