package core

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"net/url"
	"squirrel/cookies"
	"strconv"
	"strings"
)

// creating the structure of the request object
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

// func to parse the incoming request
// basically we're listening to the incoming connection and then
// parsing it by the server and returning the request parsed request object

// methods for request object

// req.Param(paramName string) => extract the param data from the url
func (r *Request) Param(paramName string) string {
	return r.Params[paramName]
}

// req.Query
// fetches Query params from the url
// returns array of query params
// handling multiple query values
func (r *Request) Query(queryName string) []string {
	return r.Queries[queryName]
}

// req.Original
// gets the original request url
func (r *Request) OriginalUrl() *url.URL {
	return r.Url
}

func ParseRequest(conn net.Conn) (*Request, error) {

	reader := bufio.NewReader(conn)      // returns a new [Reader] whose buffer has default size
	line, err := reader.ReadString('\n') // read until the first delim 'delimeter: \n'
	if err != nil {
		log.Fatal(err)
	}

	// extracting the first line of request
	// GET HTTP/1.1 OK.
	parts := strings.Fields(strings.TrimSpace(line))
	method, path := parts[0], parts[1]

	if method == "" {
		return nil, fmt.Errorf("invalid method ")
	}

	headers := map[string]string{}
	var contentLength int64
	var cookies []*cookies.Cookie

	// now read the rest of the connection request
	// parse it and add to headers as:
	// key: value
	// Content-Lenght: 123
	for {
		line, err := reader.ReadString('\n')
		if err != nil || line == "\r\n" {
			break
		}

		kv := strings.SplitN(strings.TrimSpace(line), ":", 2)
		if len(kv) == 2 {
			key := strings.TrimSpace(kv[0])
			value := strings.TrimSpace(kv[1])
			headers[key] = value

			if strings.EqualFold(key, "Content-Length") {
				contentLength, _ = strconv.ParseInt(value, 10, 64)
			}

			if strings.EqualFold(key, "Cookie") {
				cookies = ParseCookieHeader(value)
				fmt.Println(cookies)
			}

		}
	}

	// if the conn is post request
	// read the body as well
	body := io.NopCloser(strings.NewReader(""))
	if contentLength > 0 {
		bodyBuf := make([]byte, contentLength)
		_, err := io.ReadFull(body, bodyBuf)
		if err != nil {
			return nil, err
		}
		// converting the bytes into io.Reader type
		// cause, our body is actually io.ReaderCloser
		body = io.NopCloser(bytes.NewReader(bodyBuf))
	}

	query := map[string][]string{}
	u, _ := url.Parse(path)
	for k, v := range u.Query() {
		query[k] = v
	}

	actualPath := path
	if strings.HasSuffix(u.Path, "/") {
		actualPath = strings.TrimSuffix(u.Path, "/")
	}

	return &Request{
		Method:        method,
		Path:          actualPath, // just getting the pure path without query
		Url:           u,
		Body:          body,
		Headers:       headers,
		Close:         false,
		ContentLength: contentLength,
		Queries:       query,
		Cookies:       cookies,
	}, nil

}

// since the headers we receive the cookie mostly in name-value pair
// we don't need to parse it much
func ParseCookieHeader(header string) []*cookies.Cookie {

	parts := strings.Split(header, ";")

	var c []*cookies.Cookie

	for _, part := range parts {

		pair := strings.SplitN(strings.TrimSpace(part), "=", 2)

		if len(pair) == 2 {
			c = append(c, &cookies.Cookie{
				Name:  pair[0],
				Value: pair[1],
			})
		}

	}
	return c
}

// req.GetCookie
// gets the cookie by its name
func (r *Request) GetCookie(name string) *cookies.Cookie {

	for _, cookie := range r.Cookies {
		if cookie.Name == name {
			return cookie
		}
	}
	return nil
}
