# Squirrel Framework

A simple and minimalist web framework for Go.

## Overview

Squirrel Framework provides a lightweight HTTP server implementation with routing, middleware support, and request/response handling capabilities.

## Types

### Request
```go
type Request struct {
   Conn          net.Conn
	Method        string
	Path          string
	Body          io.ReadCloser // inorder to read anykind of data
	Headers       map[string]string
	Url           *url.URL
	Params        map[string]string
	ContentLength int64
	Close         bool
	Queries       map[string][]string
	Cookies       []*cookies.Cookie
}
```
Represents an HTTP request with methods for accessing request data.

### Response
```go
type Response struct {
   conn        net.Conn
	headers     map[string]string
	contentType string
	body        io.ReadCloser
	statusCode  int
	cookies     []*cookies.Cookie
}
```
Represents an HTTP response with methods for setting headers, status, and body content.

### HandlerFunc
```go
type HandlerFunc func(req *Request, res *Response)
```
Function signature for request handlers.

### route
```go
type route struct{
   method  string
	pattern string
	handler core.HandlerFunc
	// for route specific middleware
	// a route may have more than one middleware
	middleware []Middleware
}
```
Internal structure for route management.

### Middleware
```go
type Middleware func(HandlerFunc) HandlerFunc
```
Function signature for middleware that wraps handlers.

### SqurlMux
```go
type SqurlMux struct {
   // grouping together the routes 
   // the server mux
   routes []route 
	// global middlewares
	// also an application have more than on middleware
	middleware []Middleware
}
```
Main multiplexer for routing HTTP requests.

### Cookies
Cookies are small key-value pairs that the server sends to the client in a Set-Cookie header.
Then, on subsequent requests, the client includes them in the Cookie header.
```go
type Cookie struct {
	Name   string // name of the cookie
	Value  string // cookie value
	Quoted bool   // indicates whether the Value was initially Quoted or not

	Path       string    // optional
	Domain     string    // optional
	Expires    time.Time // optional
	RawExpires string    // optional

	// MaxAge = 0; means no 'MaxAge' attributes set
	// MaxAge < 0; means delete cookie now, equivalently MaxAge = 0
	// MaxAge > 0; means Max-Age attribute present and available in seconds
	MaxAge   int
	Secure   bool
	HttpOnly bool
	SameSite SameSite
	Raw      string
	Unparsed []string // Raw text of unparsed attribute-value pairs
}
```

## Methods

### Request Methods

#### `ReadBodyAsString() (string, error)`
Reads the request body and returns it as a string.

#### `Param(paramName string) string`
Retrieves a URL parameter by name.

#### `OriginalUrl() *Url.url`
Retrieves the original URL of an incoming request

#### `Query() []string` 
Retrieves a Query parameter by query name from the url
Support Multi value query parameter

#### `GetCookie(name string) *cookies.Cookie`
Gets the cookie by its name

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

#### `SetCookie(cookie *cookies.Cookie)`
Sets Set-Cookie response header and sends to the client

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

### `func FormatSetCookie(c Cookie) string`
Serialize the cookie struct into string to set in the Set-Cookie header

### `ParseCookieHeader(header string) []cookies.Cookie`
Parses the Cookie Header from the incoming request and sets in the req cookie

### `ServeStatic(prefix, dirpath string)`
Serves static files from the public directory

Usage
```go
server.ServeStatic("/assets", "public")
```


## Built-In Middlewares

### `func Logger(next core.HandlerFunc) core.HandlerFunc`
```go
import "middlewares"
server := server.SpawnServer()
server.Use(middlewares.Logger)
```


### `func Recover(next core.HandlerFunc) core.HandlerFunc`
It is by default enabled that tries to catch the un-intended panic of the server.

You can also create your own custom recover middleware and override the function using:

`func SetGlobalMiddleware(fn func(any, *core.Request, *core.Response))`

```go
import "middlewares"
middlewares.SetGlobalMiddleware(func(a any, r1 *core.Request, r2 *core.Response) {
		r2.SetStatus(500)
		r2.Write("Err: foo foo foo")
	})
```
You can disable the default recover handler using `DisableAutoRecover()`
```go
server.DisableAutoRecover()
```



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

[Not Added Yet]

## Contributing

We welcome contributions to the Squirrel Framework! Please follow these guidelines to help us maintain a high-quality codebase.

### Getting Started

1. Fork the repository on GitHub
2. Clone your fork locally:
   ```bash
   git clone https://github.com/yourusername/squirrel-framework.git
   cd squirrel-framework
   ```
3. Create a new branch for your feature or bugfix:
   ```bash
   git checkout -b feature/your-feature-name
   ```

### Development Guidelines

#### Code Style
- Follow standard Go formatting using `gofmt`
- Use meaningful variable and function names
- Add comments for exported functions and types
- Keep functions small and focused on a single responsibility

#### Testing
- Write unit tests for new features and bug fixes
- Ensure all tests pass before submitting:
  ```bash
  go test ./...
  ```
- Aim for good test coverage of new code
- Include both positive and negative test cases

#### Documentation
- Update the README.md if you add new features
- Add inline documentation for exported functions
- Include examples in your documentation when helpful

### Submitting Changes

1. **Commit your changes** with clear, descriptive commit messages:
   ```bash
   git commit -m "Add support for custom middleware ordering"
   ```

2. **Push to your fork**:
   ```bash
   git push origin feature/your-feature-name
   ```

3. **Create a Pull Request** with:
   - Clear title describing the change
   - Detailed description of what was changed and why
   - Reference to any related issues
   - Screenshots or examples if applicable

### Pull Request Guidelines

- Keep PRs focused on a single feature or bug fix
- Include tests for new functionality
- Update documentation as needed
- Ensure CI checks pass
- Respond to code review feedback promptly

### Issue Reporting

When reporting bugs or requesting features:

- Use the issue templates when available
- Provide clear reproduction steps for bugs
- Include Go version and operating system information
- Search existing issues before creating new ones

### Code of Conduct

- Be respectful and inclusive in all interactions
- Focus on constructive feedback
- Help maintain a welcoming environment for all contributors

### Questions?

If you have questions about contributing, feel free to:
- Open an issue with the "question" label
- Start a discussion in the GitHub Discussions section
- Contact the maintainers directly

Thank you for contributing to Squirrel Framework!
