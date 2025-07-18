package main

import (
	"fmt"
	"squirrel/core"
	"squirrel/examples"
	"squirrel/middlewares"
	"squirrel/server"
)

func main() {
	fmt.Println("Welcome to Squirrel")
	fmt.Println("Simple & Basic Web Backend Framework For Go")
	fmt.Println("Built on top of GO for Go")

	// server.DisableAutoRecover()
	server := server.SpawnServer()

	middlewares.SetGlobalMiddleware(func(a any, r1 *core.Request, r2 *core.Response) {
		r2.SetStatus(500)
		r2.Write("kei errror aayo")
	})
	server.Get("/home", examples.SquirrelApp)
	server.Listen(":9000")

}
