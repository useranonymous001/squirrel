package examples

import "squirrel/core"

func SquirrelApp(req *core.Request, res *core.Response) {
	name := req.Query("name")
	res.Write(name[0] + name[1])
}
