package middlewares

import "squirrel/core"

func Recover(next core.HandlerFunc) core.HandlerFunc {
	return func(r1 *core.Request, r2 *core.Response) {

		defer func() {
			if r1 := recover(); r1 != nil {
				r2.SetStatus(500)
				r2.Write("Internal Server Error")
				r2.Send()
			}
		}()

		next(r1, r2)
	}
}
