package main

import (
	"github.com/harm-matthias-harms/rpm/backend/handler"
)

func main() {
	e := handler.Server()
	//Starts the server
	e.Logger.Fatal(e.Start(":3000"))
}
