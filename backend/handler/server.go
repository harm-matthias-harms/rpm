package handler

import (
	"net/http"

	"github.com/harm-matthias-harms/rpm/backend/storage"
	"github.com/harm-matthias-harms/rpm/backend/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var errorParseRequest = "couldn't parse request"
var errorParseParams = "params couldn't be parsed"
var errorNoIDParam = "no or false id provided"
var errorIDNotMatch = "ids do not match"
var errorFind = "couldn't find entinity"
var errorCreated = "couldn't be created"
var errorUpdated = "couldn't be updated"
var errorDelete = "couldn't be deleted"
var errorAuthParse = "authorization couldn't be parsed"
var errorNotAuthorized = "not authorized"

type jsonStatus struct {
	Data    interface{} `json:"data,omitempty"`
	Success bool        `json:"success,omitempty"`
	Message string      `json:"message,omitempty"`
}

// Server provides the server for the api
func Server() (*echo.Echo, error) {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	// Middlewares
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{utils.GetEnv("DOMAIN", "http://localhost:3000")},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	}))
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	e.Use(middleware.Secure())

	// API - JWT Group
	r := e.Group("/api")
	r.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte(utils.GetEnv("JWT_SECRET", "secret")),
		TokenLookup: "cookie:" + echo.HeaderAuthorization,
	}))
	//user
	r.GET("/user", HandleUserGet)
	r.GET("/user/:id", HandleUserFind)
	// presets
	r.POST("/presets", HandlePresetCreate)
	r.GET("/presets", HandlePresetsGet)
	r.GET("/presets/:id", HandlePresetFind)
	r.PUT("/presets/:id", HandlePresetEdit)
	r.DELETE("/presets/:id", HandlePresetDelete)
	// medical cases
	r.POST("/medical_cases", HandleMedicalCaseCreate)
	r.GET("/medical_cases", HandleMedicalCaseGet)
	r.GET("/medical_cases/:id", HandleMedicalCaseFind)
	r.PUT("/medical_cases/:id", HandleMedicalCaseEdit)
	r.DELETE("/medical_cases/:id", HandleMedicalCaseDelete)
	r.GET("/medical_cases/:mc_id/documents/:id", HandleMedicalCaseFileGet)
	r.DELETE("/medical_cases/:mc_id/documents/:id", HandleMedicalCaseFileDelete)
	// teams
	r.POST("/teams", HandleTeamCreate)
	r.GET("/teams", HandleTeamsGet)
	r.GET("/teams/:id", HandleTeamFind)
	r.PUT("/teams/:id", HandleTeamEdit)
	r.DELETE("/teams/:id", HandleTeamDelete)
	// exercises
	r.POST("/exercises", HandleTeamCreate)
	r.GET("/exercises", HandleTeamsGet)
	r.GET("/exercises/:id", HandleTeamFind)
	r.PUT("/exercises/:id", HandleTeamEdit)
	r.DELETE("/exercises/:id", HandleTeamDelete)

	// Auth - NO JWT
	a := e.Group("/auth")
	a.POST("/register", HandleRegister)
	a.POST("/authenticate", HandleAuthenticate)

	// Gives a healthcheck entrypoint
	e.GET("/api/healthcheck", handlerHealthCheck)

	err := storage.SetMongoDatabase()
	if err != nil {
		return nil, err
	}

	return e, nil
}

func handlerHealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "Application is running.")
}
