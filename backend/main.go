package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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
		AllowOrigins: []string{getEnv("DOMAIN", "http://localhost:3000")},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	e.Use(middleware.Secure())

	// JWT Group
	r := e.Group("/api")
	r.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(getEnv("JWT_SECRET", "secret")),
		TokenLookup:   "cookie:" + echo.HeaderAuthorization,
	}))

	// ADD THE ENDPOINTS HERE
	// r.GET("/endpoint", handler)

	// Gives a healthcheck entrypoint
	e.GET("/healthcheck", handlerHealthCheck)

	return e
}

func handlerHealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "Application is running.")
}

// Gets an ENV-variable or returns a default value
func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}