package main

import (
	"ComputerWorld_API/server/routes"
)

func NewServer() {
	e := routes.Echo()

	// Execute server
	e.Logger.Fatal(e.Start(":5000"))
}

func main() {
	NewServer()
}
