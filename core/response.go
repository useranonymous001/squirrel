package core

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"strings"
)

// response object to send back to the client as a server
type Response struct {
	conn        net.Conn
	headers     map[string]string
	contentType string
	body        io.ReadCloser
	statusCode  int
	close       bool
}

var (
	statusText = map[int]string{
		200: "OK",
		404: "Not Found",
		302: "Redirect",
		500: "Internal Server Error",
		403: "Forbidden",
		401: "Bad Request",
		// TODO: add more status texts as needed
	}
)

// create a new response object
func NewResponse(conn *net.Conn) *Response {
	return &Response{
		conn:        *conn,
		contentType: "text/plain",
		statusCode:  200,
		headers:     map[string]string{},
	}
}

// add methods and send to the client
// res.Send() methods
// write to the client

func (r *Response) Write(body string) {
	r.body = io.NopCloser(strings.NewReader(body))
}

func (r *Response) Send() {

	// check if body is empty
	if r.body == nil {
		r.Write("")
	}

	statusLine := fmt.Sprintf("HTTP/1.1 %d %s\r\n", r.statusCode, statusText[r.statusCode])
	contentType := fmt.Sprintf("Content-Type: %s\r\n", r.contentType)

	// A Buffer is a variable-sized buffer of bytes with [Buffer.Read] and [Buffer.Write] methods

	var bodyBuf bytes.Buffer
	defer r.body.Close()

	// copy all the contents of body to the buffer
	// body need to written at the end after one blank line
	_, err := io.Copy(&bodyBuf, r.body)
	if err != nil {
		r.conn.Write([]byte("HTTP/1.1 500 Internal Server Error"))
		return
	}

	contentLength := fmt.Sprintf("Content-Length: %d\r\n", bodyBuf.Len())

	// write response object to the connection
	// along with body if available
	r.conn.Write([]byte(statusLine))
	r.conn.Write([]byte(contentType))
	r.conn.Write([]byte(contentLength))

	r.conn.Write([]byte("\r\n")) // single blank line before writing the body

	r.conn.Write(bodyBuf.Bytes())
}

// conn.Write([]byte) Write writes data to the connection.
