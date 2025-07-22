// package main

// import (
// 	"fmt"
// 	"server/core"
// 	"server/middlewares"
// 	"server/server"
// )

// func main() {
// 	fmt.Println("Welcome to Squirrel")
// 	fmt.Println("Simple & Basic Web Backend Framework For Go")
// 	fmt.Println("Built on top of GO for Go")

// 	// server.DisableAutoRecover()
// 	server := server.SpawnServer()

// 	/*
// 		prefix => /assets
// 		dirpath => core
// 		req.Path => /assets/js/file.go

// 		after trimming the prefix
// 		relPath => /js/file.go
// 		relPath => path.Clean(relPath) + dirpath
// 		relPath => core/js/file.go

// 	*/

// 	server.ServeStatic("/assets", "examples")

// 	server.Use(middlewares.Logger)
// 	middlewares.SetGlobalMiddleware(func(a any, r1 *core.Request, r2 *core.Response) {
// 		r2.SetStatus(500)
// 		r2.Write("kei errror aayo")
// 	})
// 	// server.Get("/home", examples.SquirrelApp)
// 	// server.Get("/magic", examples.RenderHtml)
// 	server.Listen(":9000")
// }

package main

import (
	"fmt"
	"log"
	"squirrel/core"
	"squirrel/server"
)

func main() {
	// Create a new Squirrel mux
	mux := server.SpawnServer()

	// Add logging middleware
	// Add recovery middleware to handle panics

	// Define routes
	// mux.HandleFunc("GET /", homeHandler)
	// mux.HandleFunc("GET /hello", helloHandler)
	// mux.HandleFunc("GET /hello/{name}", helloNameHandler)
	// mux.HandleFunc("GET /api/status", statusHandler)
	// mux.HandleFunc("POST /api/echo", echoHandler)

	mux.Get("/home", homeHandler)
	mux.Get("/echo", echoHandler)

	// Start the server
	fmt.Println("Server starting on http://localhost:8080")
	log.Fatal(mux.Listen(":8080"))
}

// Home page handler
func homeHandler(r *core.Request, w *core.Response) {
	w.SetHeader("Content-Type", "text/html")
	w.Write(`
        <html>
            <head><title>server Basic Server</title></head>
            <body>
                <h1>Welcome to server!</h1>
                <p>This is a basic server example.</p>
                <ul>
                    <li><a href="/hello">Hello World</a></li>
                    <li><a href="/hello/John">Hello John</a></li>
                    <li><a href="/api/status">API Status</a></li>
                </ul>
            </body>
        </html>
    `)
}

// Echo endpoint that returns the request body
func echoHandler(r *core.Request, w *core.Response) {
	var data map[string]interface{}

	response := map[string]interface{}{
		"echo":   data,
		"method": r.Method,
		"path":   r.Url,
	}

	w.JSON(response)
}
