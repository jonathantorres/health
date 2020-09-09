package main

import (
	"net/http"
	"strconv"
	"time"
)

type WeightEntry struct {
	Id     int64
	UserId int64
	Weight float32
	Date   string
}

func (weight *WeightEntry) SqlDate() string {
	return weight.Date[:10]
}

func weightAdd(res http.ResponseWriter, req *http.Request) {
	session := &Session{}
	session.Start(res, req)
	if !loggedIn(session) {
		http.Redirect(res, req, "/login", http.StatusFound)
		return
	}
	db, err := initDb()
	if err != nil {
		serve500(res, req, err.Error())
		return
	}
	user := getUserFromSession(session)

	if req.Method == "POST" {
		req.ParseForm()
		nowDate := time.Now()
		weight, _ := strconv.ParseFloat(req.PostForm["weight"][0], 32)
		date := req.PostForm["entered-date"][0]
		date += " " + nowDate.Format("15:04:05")
		if err := createWeightEntry(db, user.Id, float32(weight), date); err != nil {
			session.Set("errMsg", err.Error())
			http.Redirect(res, req, "/weight/add", http.StatusSeeOther)
			return
		}
		session.Set("okMsg", "Weight entry has been created successfully!")
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	setErrorAndSuccessMessages(session)
	res.Header().Set("Content-type", "text/html")
	appData.LayoutData["PageTitle"] = "Health - Add Weight Entry"
	appData.LayoutData["User"] = user
	if err := renderView("views/weight/add.html", res); err != nil {
		serveViewError(res, err)
	}
	cleanupErrorAndSuccessMessages(session)
}

func weightAll(res http.ResponseWriter, req *http.Request) {
	session := &Session{}
	session.Start(res, req)
	if !loggedIn(session) {
		http.Redirect(res, req, "/login", http.StatusFound)
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
		http.Redirect(res, req, "/login", http.StatusFound)
		return
	}
	db, err := initDb()
	if err != nil {
		serve500(res, req, err.Error())
		return
	}
	entryId, err := getId(req.URL.Path)
	if err != nil {
		serve404(res, req)
		return
	}
	user := getUserFromSession(session)
	entry, err := getWeightEntry(db, user.Id, int64(entryId))
	if err != nil {
		serve500(res, req, err.Error())
		return
	}

	if req.Method == "POST" {
		req.ParseForm()
		nowDate := time.Now()
		weight, _ := strconv.ParseFloat(req.PostForm["weight"][0], 32)
		date := req.PostForm["entered-date"][0]
		date += " " + nowDate.Format("15:04:05")
		if err := updateWeightEntry(db, user.Id, entry.Id, float32(weight), date); err != nil {
			session.Set("errMsg", err.Error())
			http.Redirect(res, req, "/weight/edit/"+strconv.Itoa(entryId), http.StatusSeeOther)
			return
		}
		session.Set("okMsg", "Weight entry has been updated successfully!")
		http.Redirect(res, req, "/weight/edit/"+strconv.Itoa(entryId), http.StatusSeeOther)
		return
	}

	setErrorAndSuccessMessages(session)
	res.Header().Set("Content-type", "text/html")
	appData.LayoutData["PageTitle"] = "Health - Edit Weight Entry"
	appData.LayoutData["User"] = user
	appData.ViewData["Entry"] = entry
	if err := renderView("views/weight/edit.html", res); err != nil {
		serveViewError(res, err)
	}
	cleanupErrorAndSuccessMessages(session)
}

func weightDelete(res http.ResponseWriter, req *http.Request) {
	session := &Session{}
	session.Start(res, req)
	if !loggedIn(session) {
		http.Redirect(res, req, "/login", http.StatusFound)
		return
	}
	db, err := initDb()
	if err != nil {
		serve500(res, req, err.Error())
		return
	}
	entryId, err := getId(req.URL.Path)
	if err != nil {
		serve404(res, req)
		return
	}
	user := getUserFromSession(session)
	entry, err := getWeightEntry(db, user.Id, int64(entryId))
	if err != nil {
		serve500(res, req, err.Error())
		return
	}
	if err = deleteWeightEntry(db, user.Id, entry.Id); err != nil {
		session.Set("errMsg", err.Error())
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	session.Set("okMsg", "Weight entry has been deleted!")
	http.Redirect(res, req, "/", http.StatusSeeOther)
}
