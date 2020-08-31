package main

import (
	"net/http"
)

func weightAdd(res http.ResponseWriter, req *http.Request) {
	session := &Session{}
	session.Start(res, req)
	if !loggedIn(session) {
		http.Redirect(res, req, "/login", http.StatusSeeOther)
		return
	}
	res.Header().Set("Content-type", "text/html")
	appData.LayoutData["PageTitle"] = "Health - Add Weight Entry"
	appData.LayoutData["User"] = getUserFromSession(session)
	if err := renderView("views/weight/add.html", res); err != nil {
		serveViewError(res, err)
	}
}

func weightAll(res http.ResponseWriter, req *http.Request) {
	session := &Session{}
	session.Start(res, req)
	if !loggedIn(session) {
		http.Redirect(res, req, "/login", http.StatusSeeOther)
		return
	}
	db, err := initDb()
	if err != nil {
		serve500(res, req, err.Error())
		return
	}
	res.Header().Set("Content-type", "text/html")

	user := getUserFromSession(session)
	entries, err := getWeightEntries(db, user.Id)
	if err != nil {
		serve500(res, req, err.Error())
		return
	}
	appData.LayoutData["PageTitle"] = "Health - Weight Entries"
	appData.LayoutData["User"] = user
	appData.ViewData["WeightHeading"] = "Weight Entries"
	appData.ViewData["Entries"] = entries
	if err := renderView("views/weight/all.html", res); err != nil {
		serveViewError(res, err)
	}
}

func weightEdit(res http.ResponseWriter, req *http.Request) {
	session := &Session{}
	session.Start(res, req)
	if !loggedIn(session) {
		http.Redirect(res, req, "/login", http.StatusSeeOther)
		return
	}
	res.Header().Set("Content-type", "text/html")
	appData.LayoutData["PageTitle"] = "Health - Edit Weight Entry"
	appData.LayoutData["User"] = getUserFromSession(session)
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
	res.Write([]byte("weight delete"))
}
