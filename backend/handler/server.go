package handler

import (
	"net/http"

	"github.com/harm-matthias-harms/rpm/backend/utils"
	"github.com/harm-matthias-harms/rpm/backend/storage"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type jsonStatus struct {
	Success bool `json:"success,omitempty"`
	Message string `json:"message,omitempty"`
}

// Server provides the server for the api
func Server() (*echo.Echo, error) {
	e := echo.New()

	// Middlewares
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookieSecure:   true,
		CookieHTTPOnly: true,
	}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{utils.GetEnv("DOMAIN", "http://localhost:3000")},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
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

	// Auth - NO JWT
	a := e.Group("/auth")
	a.POST("/register", HandleRegister)

	// ADD THE ENDPOINTS HERE
	// r.GET("/endpoint", handler)

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
