# Squirrel Framework

*** Types ***

    type Request struct {}
    type Response struct {}
    type HandlerFunc func(req *Request, res *Response)
    type route struct{}
    type Middleware func (HandlerFunc) HandlerFunc
    type SqurlMux struct {}


*** Methods ***

    type Request struct {}
        - ReadBodyAsString() (string error)
        - Param(paramName string) string
        - Query() // coming soon...

    type Response struct {}
        - SetHeader(key, value string)
        - SetStatus(status int)
        - Write (body string)
        - SetBody(reader io.ReadCloser)
        - WriteBytes(b []byte)
        - Close()
        - JSON(data interface{})
        - Send()

    type HandlerFunc func(req *Request, res *Response)

    type route struct{}

    type Middleware func (HandlerFunc) HandlerFunc

    type SqurlMux struct {}
        - Listen(addr string)
        - Use(mw Middleware)
        - Get(path, string, handler HandlerFunc, mws ...Middleware)
        - Post(path, string, handler HandlerFunc, mws ...Middleware)
        - Put(path, string, handler HandlerFunc, mws ...Middleware)
        - Delete(path, string, handler HandlerFunc, mws ...Middleware)    



*** Functions ***
    func SpawnServer() *SqurlMux
    func NewResponse(conn *net.Conn) *Response 
    func ParseRequest(conn *net.Conn) *(Response, error) 