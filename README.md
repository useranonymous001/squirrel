# Squirrel Framework

A simple and minimalist web framework for Go.

## Overview

Squirrel Framework provides a lightweight HTTP server implementation with routing, middleware support, and request/response handling capabilities.

## Types

### Request
```go
type Request struct {}
```
Represents an HTTP request with methods for accessing request data.

### Response
```go
type Response struct {}
```
Represents an HTTP response with methods for setting headers, status, and body content.

### HandlerFunc
```go
type HandlerFunc func(req *Request, res *Response)
```
Function signature for request handlers.

### route
```go
type route struct{}
```
Internal structure for route management.

### Middleware
```go
type Middleware func(HandlerFunc) HandlerFunc
```
Function signature for middleware that wraps handlers.

### SqurlMux
```go
type SqurlMux struct {}
```
Main multiplexer for routing HTTP requests.

## Methods

### Request Methods

#### `ReadBodyAsString() (string, error)`
Reads the request body and returns it as a string.

#### `Param(paramName string) string`
Retrieves a URL parameter by name.

#### `Query()` 
*Coming soon...*

### Response Methods

#### `SetHeader(key, value string)`
Sets an HTTP header.

#### `SetStatus(status int)`
Sets the HTTP status code.

#### `Write(body string)`
Writes a string to the response body.

#### `SetBody(reader io.ReadCloser)`
Sets the response body from an io.ReadCloser.

#### `WriteBytes(b []byte)`
Writes bytes to the response body.

#### `Close()`
Closes the response.

#### `JSON(data interface{})`
Writes JSON data to the response.

#### `Send()`
Sends the response to the client.

### SqurlMux Methods

#### `Listen(addr string)`
Starts the server listening on the specified address.

#### `Use(mw Middleware)`
Adds middleware to the request processing pipeline.

#### `Get(path string, handler HandlerFunc, mws ...Middleware)`
Registers a GET route handler.

#### `Post(path string, handler HandlerFunc, mws ...Middleware)`
Registers a POST route handler.

#### `Put(path string, handler HandlerFunc, mws ...Middleware)`
Registers a PUT route handler.

#### `Delete(path string, handler HandlerFunc, mws ...Middleware)`
Registers a DELETE route handler.

## Functions

### `SpawnServer() *SqurlMux`
Creates and returns a new SqurlMux instance.

### `NewResponse(conn *net.Conn) *Response`
Creates a new Response instance from a network connection.

### `ParseRequest(conn *net.Conn) (*Request, error)`
Parses an HTTP request from a network connection.

## Installation

```bash
go get github.com/useranonymous001/squirrel
```

## Quick Start

```go
package main

import "github.com/useranonymous001/squirrel"

func main() {
    server := SpawnServer()
    
    server.Get("/", func(req *Request, res *Response) {
        res.Write("Hello, World!")
        res.Send()
    })
    
    server.Listen(":8080")
}
```

## License

[Your License Here]

## Contributing

[Contributing guidelines here]
