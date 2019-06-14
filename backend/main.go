package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoSession *mongo.Database

func main() {
	e := server()
	mongoSession = setMongoDatabase()
	//Starts the server
	e.Logger.Fatal(e.Start(":3000"))
}

func setMongoDatabase() *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(getEnv("MONGO_URL", "mongodb://localhost:27017")))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return client.Database(getEnv("MONGO_DATABASE", "test"))
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
		SigningKey:  []byte(getEnv("JWT_SECRET", "secret")),
		TokenLookup: "cookie:" + echo.HeaderAuthorization,
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
