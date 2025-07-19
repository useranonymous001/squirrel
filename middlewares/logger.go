package middlewares

import (
	"fmt"
	"squirrel/core"
	"time"
)

// logger logs the request/response status of the current server
func Logger(next core.HandlerFunc) core.HandlerFunc {
	return func(req *core.Request, res *core.Response) {
		start := time.Now()
		next(req, res)
		duration := time.Since(start)
		statusCode := res.GetStatusCode()
		if statusCode == 0 {
			statusCode = 200
		}
		logMessage := fmt.Sprintf("[%s] - %s - %s - %d - %s", start.Format("2006-01-02 15:04:05"), req.Method, req.Path, statusCode, duration)

		fmt.Println(logMessage)
	}
}
