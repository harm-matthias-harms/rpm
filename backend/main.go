package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/harm-matthias-harms/rpm/backend/utils"
	"github.com/harm-matthias-harms/rpm/backend/handler"
)

func main() {
	e := server()
	//Starts the server
	e.Logger.Fatal(e.Start(":3000"))
}

func server() *echo.Echo {
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
	a.POST("/register", handler.HandleRegister)

	// ADD THE ENDPOINTS HERE
	// r.GET("/endpoint", handler)

	// Gives a healthcheck entrypoint
	e.GET("/api/healthcheck", handlerHealthCheck)

	return e
}

func handlerHealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "Application is running.")
}
