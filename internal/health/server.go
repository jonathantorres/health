package health

import (
	"fmt"
	"net/http"
	"os"
)

func Root(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/" {
		Index(res, req)
		return
	}
	serveStaticFile(res, req)
}

func Serve500(res http.ResponseWriter, req *http.Request, msg string) {
	res.Header().Set("Content-type", "text/html")
	res.WriteHeader(http.StatusInternalServerError)
	res.Write([]byte(msg))
}

func ServeViewError(res http.ResponseWriter, err error) {
	res.WriteHeader(http.StatusInternalServerError)
	res.Write([]byte(fmt.Sprintf("error rendering view: %s", err)))
}

func serveStaticFile(res http.ResponseWriter, req *http.Request) {
	path := "./public" + req.URL.Path
	fi, err := os.Stat(path)
	if err != nil {
		serve404(res, req)
		return
	}
	if fi.IsDir() {
		serve404(res, req)
		return
	}
	http.ServeFile(res, req, path)
}

func serve404(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "text/html")
	res.WriteHeader(http.StatusNotFound)
	res.Write([]byte("<p>404 page was not found</p>"))
}
