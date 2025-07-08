package core

// client side handling
// handler func
// handler function for each route
// accepts the request and write response to the client

type HandlerFunc func(*Request, *Response)
