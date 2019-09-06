package main

import (
	"github.com/harm-matthias-harms/rpm/backend/handler"
	"log"
)

func main() {
	e, err := handler.Server()
	if err != nil {
		log.Fatal(err)
	}
	//Starts the server
	e.Logger.Fatal(e.Start(":3001"))
}
