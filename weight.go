package main

import (
	"net/http"
)

func weightAdd(res http.ResponseWriter, req *http.Request) {
	if !loggedIn(res, req) {
		http.Redirect(res, req, "/login", http.StatusSeeOther)
		return
	}
	res.Header().Set("Content-type", "text/html")
	appData.LayoutData["PageTitle"] = "Health - Add Weight Entry"
	if err := renderView("views/weight/add.html", res); err != nil {
		serveViewError(res, err)
	}
}

func weightAll(res http.ResponseWriter, req *http.Request) {
	if !loggedIn(res, req) {
		http.Redirect(res, req, "/login", http.StatusSeeOther)
		return
	}
	res.Header().Set("Content-type", "text/html")
	appData.LayoutData["PageTitle"] = "Health - Weight Entries"
	if err := renderView("views/weight/all.html", res); err != nil {
		serveViewError(res, err)
	}
}

func weightEdit(res http.ResponseWriter, req *http.Request) {
	if !loggedIn(res, req) {
		http.Redirect(res, req, "/login", http.StatusSeeOther)
		return
	}
	res.Header().Set("Content-type", "text/html")
	appData.LayoutData["PageTitle"] = "Health - Edit Weight Entry"
	_, err := getId(req.URL.Path) // todo: use id here
	if err != nil {
		serve404(res, req)
		return
	}
	if err := renderView("views/weight/edit.html", res); err != nil {
		serveViewError(res, err)
	}
}

func weightDelete(res http.ResponseWriter, req *http.Request) {
	if !loggedIn(res, req) {
		http.Redirect(res, req, "/login", http.StatusSeeOther)
		return
	}
	res.Header().Set("Content-type", "text/html")
	_, err := getId(req.URL.Path) // todo: use id here
	if err != nil {
		serve404(res, req) // todo: probably redirect instead of serving 404 page
		return
	}
	res.Write([]byte("weight delete"))
}
