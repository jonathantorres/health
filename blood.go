package main

import (
	"net/http"
)

func bloodAdd(res http.ResponseWriter, req *http.Request) {
	session := &Session{}
	session.Start(res, req)
	if !loggedIn(session) {
		http.Redirect(res, req, "/login", http.StatusSeeOther)
		return
	}
	res.Header().Set("Content-type", "text/html")
	appData.LayoutData["PageTitle"] = "Health - Blood Pressure Add Reading"
	appData.LayoutData["User"] = getUserFromSession(session)
	if err := renderView("views/blood/add.html", res); err != nil {
		serveViewError(res, err)
	}
}

func bloodAll(res http.ResponseWriter, req *http.Request) {
	session := &Session{}
	session.Start(res, req)
	if !loggedIn(session) {
		http.Redirect(res, req, "/login", http.StatusSeeOther)
		return
	}
	res.Header().Set("Content-type", "text/html")
	appData.LayoutData["PageTitle"] = "Health - Blood Pressure Readings"
	appData.LayoutData["User"] = getUserFromSession(session)
	if err := renderView("views/blood/all.html", res); err != nil {
		serveViewError(res, err)
	}
}

func bloodDetails(res http.ResponseWriter, req *http.Request) {
	session := &Session{}
	session.Start(res, req)
	if !loggedIn(session) {
		http.Redirect(res, req, "/login", http.StatusSeeOther)
		return
	}
	res.Header().Set("Content-type", "text/html")
	appData.LayoutData["PageTitle"] = "Health - Blood Pressure Reading Details"
	appData.LayoutData["User"] = getUserFromSession(session)
	_, err := getId(req.URL.Path) // todo: use id here
	if err != nil {
		serve404(res, req)
		return
	}
	if err = renderView("views/blood/details.html", res); err != nil {
		serveViewError(res, err)
	}
}

func bloodEdit(res http.ResponseWriter, req *http.Request) {
	session := &Session{}
	session.Start(res, req)
	if !loggedIn(session) {
		http.Redirect(res, req, "/login", http.StatusSeeOther)
		return
	}
	res.Header().Set("Content-type", "text/html")
	appData.LayoutData["PageTitle"] = "Health - Edit Blood Pressure Reading"
	appData.LayoutData["User"] = getUserFromSession(session)
	_, err := getId(req.URL.Path) // todo: use id here
	if err != nil {
		serve404(res, req)
		return
	}
	if err := renderView("views/blood/edit.html", res); err != nil {
		serveViewError(res, err)
	}
}

func bloodDelete(res http.ResponseWriter, req *http.Request) {
	session := &Session{}
	session.Start(res, req)
	if !loggedIn(session) {
		http.Redirect(res, req, "/login", http.StatusSeeOther)
		return
	}
	res.Header().Set("Content-type", "text/html")
	_, err := getId(req.URL.Path) // todo: use id here
	if err != nil {
		serve404(res, req) // todo: probably redirect instead of serving 404 page
		return
	}
	res.Write([]byte("blood delete"))
}
