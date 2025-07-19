package internal

import (
	"fmt"
	"io"
	"mime"
	"os"
	"path"
	"path/filepath"
	"squirrel/core"
	"strings"
)

/*
	Logic Behind serving static files


	- Not creating the every file/folder inside as the independent route
	- instead, creating a virtual route (prefix, dir) and
	- this should dynamically lookup for the subpaths inside the dir

	- example
	- serveStatic("/assets", "public")

	- /assets will be the virtual route and the files/folder of the public
	- directory would be the subpaths automatically


*/

func ServeStaticHandler(prefix string, dirpath string) core.HandlerFunc {
	if prefix == "" {
		prefix = "/"
	}

	if dirpath == "" {
		dirpath = "./public"
	}

	return func(req *core.Request, res *core.Response) {

		if req.Method != "GET" && req.Method != "HEAD" {
			// other method than GET and HEAD are not allowed
			res.SetStatus(405)
			res.SetHeader("Allow", "GET, HEAD")
			res.SetHeader("Content-Length", "0")
			res.Send()
			return
		}

		// replace the route path with directory name
		// /asset => "public"
		// cuz  we are storing all the static files over here.
		relPath := strings.TrimPrefix(req.Path, prefix)

		// if only given, /assets || /assets
		// setting the default path to index.html
		if relPath == "" || relPath == "/" {
			relPath = "index.html"
		}

		// get the file/folder that the path is trying to access
		// we are cleaning the relPath to avoid
		// directory traveral attack like: "../../../core"

		fullpath := path.Join(dirpath, path.Clean(relPath))

		// checking if the path is still inside the public directory or not
		// preventing the directory traversal attack
		absBase, _ := filepath.Abs(dirpath)
		absTargetPath, _ := filepath.Abs(fullpath)

		if !strings.HasPrefix(absTargetPath, absBase) {
			// ohh-ohh, trying to getout of you limit xD
			res.SetStatus(403)
			res.Write("403 Forbidden")
			// res.Send()
			return
		}

		file, err := os.Open(fullpath)
		if err != nil {
			res.SetStatus(404)
			res.Write("404 File Not Found")
			res.Send()
			return
		}
		defer file.Close()

		mimetype := mime.TypeByExtension(filepath.Ext(fullpath))

		if mimetype != "" {
			res.SetHeader("Content-Type", mimetype)
		}

		body, err := io.ReadAll(file)
		if err != nil {
			res.SetStatus(500)
			res.Write("Error Parsing Body")
			res.Send()
			return
		}

		res.SetHeader("Content-Length", fmt.Sprint(len(body)))
		res.WriteBytes(body)
		res.Send()
	}
}
