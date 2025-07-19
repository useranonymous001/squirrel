package examples

import (
	"fmt"
	"squirrel/cookies"
	"squirrel/core"
)

func SquirrelApp(req *core.Request, res *core.Response) {

	token := cookies.Cookie{
		Name:   "lazinerd",
		Value:  "hidden_cookie_inside_browser",
		Secure: true,
	}

	fmt.Printf("\ntoken: %v\n", token)

	res.SetCookie(token)
	// myCookie := req.GetCookie("auth_token")
	// myCookie := req.GetCookie("lazinerd")

	// res.Write(name[0] + name[1])
	// res.Write(myCookie.Value)
	// res.Write(myCookie.Name)
	// res.JSON(myCookie)
	res.Write("hello homie")
}

func RenderHtml(req *core.Request, res *core.Response) {
	html := "<h1>Is this even happening??</h1>"
	res.WriteBytes([]byte(html))
}
