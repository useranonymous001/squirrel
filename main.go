package main

import (
	"fmt"
	"squirrel/examples"
	"squirrel/server"
)

func main() {
	fmt.Println("Welcome to Squirrel")
	fmt.Println("Simple & Basic Web Backend Framework For Go")
	fmt.Println("Built on top of GO for Go")

	server := server.SpawnServer()
	server.Get("/home", examples.SquirrelApp)
	server.Listen(":9000")
}
