package middlewares

import (
	"log"
	"squirrel/core"
)

// recover function is used for catching any kind of panics in the applications
var defaultErrorHanlder = func(err any, req *core.Request, res *core.Response) {
	log.Printf("Error at: (%s) %v", req.Path, err)
	res.SetStatus(500)
	res.Write("Internal Server Error\n")
	res.Send()
}

// Global Recover Function to Listen to any kind of errors / panics
// It is by-default set whenever the user spawns a new Server
// server = server.SpawnServer()
func Recover(next core.HandlerFunc) core.HandlerFunc {
	return func(req *core.Request, res *core.Response) {
		defer func() {
			if err := recover(); err != nil {
				defaultErrorHanlder(err, req, res)
			}
		}()
		next(req, res)
	}
}

// allow developers to override the custom global middlewares overridding the default ones
// SetGlobalErrorHandler(func(err any, req *core.Request, res *core.Response)
func SetGlobalMiddleware(fn func(any, *core.Request, *core.Response)) {
	defaultErrorHanlder = fn
}

// todo: built-in function to disable the default "Recover()" middleware
// func DisableAutoRecover() {
// }
