package examples

import (
	"squirrel/core"
)

func SquirrelApp(req *core.Request, res *core.Response) {
	// name := req.Query("name")

	// cookie := &cookies.Cookie{
	// 	Name:  "auth_token",
	// 	Value: "what the fuck is this token",
	// }

	// res.SetCookie(cookie)

	myCookie := req.GetCookie("auth_token")
	// res.Write(name[0] + name[1])
	res.Write(myCookie.Value)
	res.Write(myCookie.Name)
	res.Write(myCookie.Value)
	res.JSON(myCookie)
}
