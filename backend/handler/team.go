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

// HandleTeamCreate creates a team by a given json
func HandleTeamCreate(c echo.Context) (err error) {
	team := new(model.Team)
	if err := c.Bind(team); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errorParseRequest)
	}
	cookie, _ := c.Cookie(echo.HeaderAuthorization)
	token, _ := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
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
