package main

import (
	"net/http"
)

func weightAdd(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "text/html")
	layoutData.PageTitle = "Health - Add Weight Entry"
	if err := renderView("views/weight/add.html", res); err != nil {
		serveViewError(res, err)
	}
}

func weightAll(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "text/html")
	layoutData.PageTitle = "Health - Weight Entries"
	if err := renderView("views/weight/all.html", res); err != nil {
		serveViewError(res, err)
	}
}

func weightEdit(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "text/html")
	layoutData.PageTitle = "Health - Edit Weight Entry"
	if err := renderView("views/weight/edit.html", res); err != nil {
		serveViewError(res, err)
	}
}

func weightDelete(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "text/html")
	res.Write([]byte("weight delete"))
}
