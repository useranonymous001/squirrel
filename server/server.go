package server

import "squirrel/core"

/*
Serve package handles the following things

	- listen to the connection and spawn a new go routine for each request
	- accepts the connection and then parse the req and res
	- serveMux for grouping the routes and the handlers
	- handle params from the routes
	- call the handler function and send the response to the client
	- allow middlewares integration: Global As well as Router Specific Middlewares
*/

// Serve Mux acts as HTTP Request Multiplexer.
// It matches the url of the incoming request with the registered routes
// in the route data structure and calls the handler for the pattern that
// closely matches to the url

type Middleware func(core.HandlerFunc) core.HandlerFunc

type route struct {
	method  string
	pattern string
	handler core.HandlerFunc
	// for route specific middleware
	// a route may have more than one middleware
	middleware []Middleware
}

type SquirrelMux struct {
	routes []route // grouping together the routes to the server mux
	// global middlewares
	// also an application have more than one middleware
	middleware []Middleware
}

// start or create an instance of server accepting the connection
// allows to have all communicattion between client and server
func SpawnServer() *SquirrelMux {
	return &SquirrelMux{}
}

// methods for using/defining the global middleware
// server.Use(Middleware)
func (sm *SquirrelMux) Use(mw Middleware) {
	sm.middleware = append(sm.middleware, mw)
}

// differnt http methods
// GET POST PUT DELETE

func (sm *SquirrelMux) Get(path string, handler core.HandlerFunc, mws ...Middleware) {
	sm.routes = append(sm.routes, route{
		method:     "GET",
		pattern:    path,
		handler:    handler,
		middleware: mws,
	})
}

func (sm *SquirrelMux) Post(path string, handler core.HandlerFunc, mws ...Middleware) {
	sm.routes = append(sm.routes, route{
		pattern:    path,
		method:     "POST",
		handler:    handler,
		middleware: mws,
	})
}

func (sm *SquirrelMux) PUT(path string, handler core.HandlerFunc, mws ...Middleware) {
	sm.routes = append(sm.routes, route{
		pattern:    path,
		method:     "PUT",
		handler:    handler,
		middleware: mws,
	})
}

func (sm *SquirrelMux) Delete(path string, handler core.HandlerFunc, mws ...Middleware) {
	sm.routes = append(sm.routes, route{
		pattern:    path,
		method:     "DELETE",
		handler:    handler,
		middleware: mws,
	})
}
