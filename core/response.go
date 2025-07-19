package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"squirrel/cookies"
	"strings"
)

// response object to send back to the client as a server
type Response struct {
	conn        net.Conn
	headers     map[string]string
	contentType string
	body        io.ReadCloser
	statusCode  int
	cookies     []*cookies.Cookie
}


// create a new response object
func NewResponse(conn *net.Conn) *Response {
	return &Response{
		conn:        *conn,
		contentType: "text/plain",
		statusCode:  200,
		headers:     map[string]string{},
	}
}

// res.SetHeader
// Sets header for response body
func (r *Response) SetHeader(key, value string) {
	r.headers[key] = value
}

// res.SetStatus
// sets the status code for the response
func (r *Response) SetStatus(status int) {
	r.statusCode = status
}


//res.GetStatusCode
//getter function to private field of struct Response
func (r *Response) GetStatusCode() int {
	return r.statusCode
}
// res.SetBody
// accepts io.Reader type
// works for any kind of stream data like:
// reading File or Stream
func (r *Response) SetBody(reader io.ReadCloser) {
	r.body = io.NopCloser(reader)
}

// add methods and send to the client
// res.Send() methods
// write to the client

func (r *Response) Write(body string) {
	r.body = io.NopCloser(strings.NewReader(body))
}

// res.WriteBytes(b []byte)
// accepts bytes of data and returns the io.ReadCloser
// works well for binary responses Images Files Compressed/gzipped data
func (r *Response) WriteBytes(b []byte) {
	r.body = io.NopCloser(bytes.NewReader(b))
}

// res.JSON()
// allows to send the response as json data
// accepts any kind of data and returns response in json format
func (r *Response) JSON(data interface{}) {
	b, err := json.MarshalIndent(data, "", "")
	if err != nil {
		r.conn.Write([]byte("HTTP/1.1 500 Internal Server Error\r\n\r\n"))
		return
	}

	r.SetHeader("Content-Type", "application/json")
	r.contentType = "application/json"

	r.WriteBytes(b)

}

func (r *Response) Send() {

	// check if body is empty
	if r.body == nil {
		r.Write("")
	}

	statusLine := fmt.Sprintf("HTTP/1.1 %d %s\r\n", r.statusCode, http.StatusText(r.statusCode))
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

	// set the cookie headers to the client if available
	for _, cookie := range r.cookies {
		cookieHeader := cookies.FormatSetCookie(cookie)
		r.SetHeader("Set-Cookie", cookieHeader)
		// r.conn.Write([]byte("Set-Cookie: " + cookieHeader + "\r\n"))
		r.conn.Write([]byte(fmt.Sprintf("Set-Cookie: %s\r\n", cookieHeader)))
	}

	r.conn.Write([]byte("\r\n")) // single blank line before writing the body

	r.conn.Write(bodyBuf.Bytes())
}

// conn.Write([]byte) Write writes data to the connection.

// res.SetCookie ()
// sets the cookie to the response
// and send to the client
func (r *Response) SetCookie(cookie *cookies.Cookie) {
	r.cookies = append(r.cookies, cookie)
}
