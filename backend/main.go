package main

import (
	"log"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/harm-matthias-harms/rpm/backend/handler"
)

func main() {
	e, err := handler.Server()
	if err != nil {
		log.Fatal(err)
	}
	e.Server.Addr = ":3001"
	//Starts the server
	e.Logger.Fatal(gracehttp.Serve(e.Server))
}
