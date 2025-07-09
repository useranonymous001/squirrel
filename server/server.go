package server

import (
	"fmt"
	"log"
	"net"
	"squirrel/core"
	"strings"
)

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

/*

	Explanation:
	routeHandler = rt.middlewares[i](routeHandler)

	Our Middleware is something like this:

	--> type Middleware func(core.HandlerFunc) core.HandlerFunc


	So, it accepts a function of type HandlerFunc(req, res) and returns the same type

	so, here:
	rt.middlewares[i] is middleware of type Middleare and routeHandler is of type HandlerFunc

	So, Basically, it is:
		Middlware(HandlerFunc) {
			return HandlerFunc
		}

*/

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

func (sm *SquirrelMux) Listen(addr string) error {

	// listen to the tcp connection request
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("error listening to %s", addr)
	}

	log.Println("Listening at ", addr)

	for {

		conn, err := ln.Accept()
		if err != nil {
			return fmt.Errorf("error while accepting connection request: %v", err)
		}

		// if no err
		// for each connection, spawn a new go routine

		go func(conn net.Conn) {
			defer conn.Close()

			// parse the incoming request
			req, err := core.ParseRequest(conn)
			if err != nil {
				log.Fatal("Error while parsing request", err)
				return
			}

			// create new response object for the server to send back to the client
			// also for handler function
			res := core.NewResponse(&conn)

			// things to do:
			// get params if any
			// get the middlewares if any (both global and route specific)
			// send the response to the client

			params := map[string]string{} // for registering the params if any

			// go through each routes in the mux server and call the route that matches the specific pattern
			// matching includes:
			// simple mathcing the patterns to attaching the middlewares for that routes to matching the exact params
			// this is important
			for _, rt := range sm.routes {

				if rt.method != req.Method {
					continue // doesn't match with the incoming req, so skip and compare other route
				}

				if matched := matchPattern(rt.pattern, req.Path, params); matched {
					// do something
					req.Params = params

					// explanation of middleware handler at top
					// create handler along with available middlewares
					routeHandler := rt.handler

					// handling route specific middlewares
					// cause we are executing all middlewares in reverse order of its implementation
					for i := len(rt.middleware) - 1; i >= 0; i-- {
						routeHandler = rt.middleware[i](routeHandler) // wrapping the handler with global middleware
					}

					// handling global  middlewares
					for i := len(sm.middleware) - 1; i >= 0; i-- {
						routeHandler = sm.middleware[i](routeHandler) // keep on wrapping the handler with middlewares
						// so that middleware executes first and then the handler
					}

					// calling the handler function
					routeHandler(req, res)
					res.Send()
					return
				}
			}

			res.SetStatus(404)
			res.Write("Route Not Found")
			res.Send()

		}(conn)

	}
}

// pattern ==> represents the path registered in while creating routes
// path ==> represents the actual incoming path from the client and saved in the req.Path

// eg:
// pattern => /user/profile/:id
// path => req.Path => /user/Profile/23

func matchPattern(pattern, path string, params map[string]string) bool {

	patternSegment := strings.Split(pattern, "/") // [user profile :id]
	pathSegment := strings.Split(path, "/")       // [user profile 12]

	// incoming req route path doesn't matches with the pattern
	if len(patternSegment) != len(pathSegment) {
		return false
	}

	for i := 0; i < len(patternSegment); i++ {
		if strings.HasPrefix(patternSegment[i], ":") {
			paramName := strings.TrimPrefix(patternSegment[i], ":")
			params[paramName] = pathSegment[i]
			continue
		}
		// check of the each route path matches or not
		if patternSegment[i] != pathSegment[i] {
			return false
		}
	}

	return true
}
