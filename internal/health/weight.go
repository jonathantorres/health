package health

import (
	"net/http"
	"strconv"
	"time"

	"github.com/jonathantorres/health/internal/db"
	"github.com/jonathantorres/health/internal/session"
)

func WeightAdd(res http.ResponseWriter, req *http.Request) {
	sess := &session.Session{}
	sess.Start(res, req)
	if !sess.LoggedIn() {
		http.Redirect(res, req, "/login", http.StatusFound)
		return
	}
	dbs, err := db.InitDb()
	if err != nil {
		serve500(res, req, err.Error())
		return
	}
	user := sess.GetUserFromSession()

	if req.Method == "POST" {
		req.ParseForm()
		nowDate := time.Now()
		weight, _ := strconv.ParseFloat(req.PostForm["weight"][0], 32)
		date := req.PostForm["entered-date"][0]
		date += " " + nowDate.Format("15:04:05")
		if err := db.CreateWeightEntry(dbs, user.Id, float32(weight), date); err != nil {
			sess.Set("errMsg", err.Error())
			http.Redirect(res, req, "/weight/add", http.StatusSeeOther)
			return
		}
		sess.Set("okMsg", "Weight entry has been created successfully!")
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	sess.SetErrorAndSuccessMessages(appData)
	res.Header().Set("Content-type", "text/html")
	appData.LayoutData["PageTitle"] = "Health - Add Weight Entry"
	appData.LayoutData["User"] = user
	if err := renderView("views/weight/add.html", res); err != nil {
		serveViewError(res, err)
	}
	sess.CleanupErrorAndSuccessMessages(appData)
}

func WeightAll(res http.ResponseWriter, req *http.Request) {
	sess := &session.Session{}
	sess.Start(res, req)
	if !sess.LoggedIn() {
		http.Redirect(res, req, "/login", http.StatusFound)
		return
	}
	dbs, err := db.InitDb()
	if err != nil {
		serve500(res, req, err.Error())
		return
	}
	res.Header().Set("Content-type", "text/html")

	user := sess.GetUserFromSession()
	entries, err := db.GetWeightEntries(dbs, user.Id)
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

func WeightEdit(res http.ResponseWriter, req *http.Request) {
	sess := &session.Session{}
	sess.Start(res, req)
	if !sess.LoggedIn() {
		http.Redirect(res, req, "/login", http.StatusFound)
		return
	}
	dbs, err := db.InitDb()
	if err != nil {
		serve500(res, req, err.Error())
		return
	}
	entryId, err := getId(req.URL.Path)
	if err != nil {
		serve404(res, req)
		return
	}
	user := sess.GetUserFromSession()
	entry, err := db.GetWeightEntry(dbs, user.Id, int64(entryId))
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
		if err := db.UpdateWeightEntry(dbs, user.Id, entry.Id, float32(weight), date); err != nil {
			sess.Set("errMsg", err.Error())
			http.Redirect(res, req, "/weight/edit/"+strconv.Itoa(entryId), http.StatusSeeOther)
			return
		}
		sess.Set("okMsg", "Weight entry has been updated successfully!")
		http.Redirect(res, req, "/weight/edit/"+strconv.Itoa(entryId), http.StatusSeeOther)
		return
	}

	sess.SetErrorAndSuccessMessages(appData)
	res.Header().Set("Content-type", "text/html")
	appData.LayoutData["PageTitle"] = "Health - Edit Weight Entry"
	appData.LayoutData["User"] = user
	appData.ViewData["Entry"] = entry
	if err := renderView("views/weight/edit.html", res); err != nil {
		serveViewError(res, err)
	}
	sess.CleanupErrorAndSuccessMessages(appData)
}

func WeightDelete(res http.ResponseWriter, req *http.Request) {
	sess := &session.Session{}
	sess.Start(res, req)
	if !sess.LoggedIn() {
		http.Redirect(res, req, "/login", http.StatusFound)
		return
	}
	dbs, err := db.InitDb()
	if err != nil {
		serve500(res, req, err.Error())
		return
	}
	entryId, err := getId(req.URL.Path)
	if err != nil {
		serve404(res, req)
		return
	}
	user := sess.GetUserFromSession()
	entry, err := db.GetWeightEntry(dbs, user.Id, int64(entryId))
	if err != nil {
		serve500(res, req, err.Error())
		return
	}
	if err = db.DeleteWeightEntry(dbs, user.Id, entry.Id); err != nil {
		sess.Set("errMsg", err.Error())
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	sess.Set("okMsg", "Weight entry has been deleted!")
	http.Redirect(res, req, "/", http.StatusSeeOther)
}
