package health

import (
	"net/http"
	"strconv"
	"time"

	"github.com/jonathantorres/health/internal/db"
	"github.com/jonathantorres/health/internal/session"
)

func BloodAdd(res http.ResponseWriter, req *http.Request) {
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
		sys, _ := strconv.Atoi(req.PostForm["sys"][0])
		dia, _ := strconv.Atoi(req.PostForm["dia"][0])
		pulse, _ := strconv.Atoi(req.PostForm["pulse"][0])
		date := req.PostForm["reading-date"][0]
		date += " " + nowDate.Format("15:04:05")
		if err := db.CreateBloodReading(dbs, user.Id, int32(sys), int32(dia), int32(pulse), date); err != nil {
			sess.Set("errMsg", err.Error())
			http.Redirect(res, req, "/blood/add", http.StatusSeeOther)
			return
		}
		sess.Set("okMsg", "Blood reading has been created successfully!")
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	sess.SetErrorAndSuccessMessages(appData)
	res.Header().Set("Content-type", "text/html")
	appData.LayoutData["PageTitle"] = "Health - Blood Pressure Add Reading"
	appData.LayoutData["User"] = sess.GetUserFromSession()
	if err := renderView("views/blood/add.html", res); err != nil {
		serveViewError(res, err)
	}
	sess.CleanupErrorAndSuccessMessages(appData)
}

func BloodAll(res http.ResponseWriter, req *http.Request) {
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
	readings, err := db.GetBloodReadings(dbs, user.Id)
	if err != nil {
		serve500(res, req, err.Error())
		return
	}
	appData.LayoutData["PageTitle"] = "Health - Blood Pressure Readings"
	appData.LayoutData["User"] = user
	appData.ViewData["BloodHeading"] = "Blood Pressure Readings"
	appData.ViewData["Readings"] = readings
	if err := renderView("views/blood/all.html", res); err != nil {
		serveViewError(res, err)
	}
}

func BloodDetails(res http.ResponseWriter, req *http.Request) {
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
	readingId, err := getId(req.URL.Path)
	if err != nil {
		serve404(res, req)
		return
	}

	user := sess.GetUserFromSession()
	reading, err := db.GetBloodReading(dbs, user.Id, int64(readingId))
	if err != nil {
		serve500(res, req, err.Error())
		return
	}
	appData.LayoutData["PageTitle"] = "Health - Blood Pressure Reading Details"
	appData.LayoutData["User"] = user
	appData.ViewData["Reading"] = reading
	if err = renderView("views/blood/details.html", res); err != nil {
		serveViewError(res, err)
	}
}

func BloodEdit(res http.ResponseWriter, req *http.Request) {
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
	readingId, err := getId(req.URL.Path)
	if err != nil {
		serve404(res, req)
		return
	}
	user := sess.GetUserFromSession()
	reading, err := db.GetBloodReading(dbs, user.Id, int64(readingId))
	if err != nil {
		serve500(res, req, err.Error())
		return
	}

	if req.Method == "POST" {
		req.ParseForm()
		nowDate := time.Now()
		sys, _ := strconv.Atoi(req.PostForm["sys"][0])
		dia, _ := strconv.Atoi(req.PostForm["dia"][0])
		pulse, _ := strconv.Atoi(req.PostForm["pulse"][0])
		date := req.PostForm["reading-date"][0]
		date += " " + nowDate.Format("15:04:05")
		if err := db.UpdateBloodReading(dbs, user.Id, reading.Id, int32(sys), int32(dia), int32(pulse), date); err != nil {
			sess.Set("errMsg", err.Error())
			http.Redirect(res, req, "/blood/edit/"+strconv.Itoa(readingId), http.StatusSeeOther)
			return
		}
		sess.Set("okMsg", "Blood reading has been updated successfully!")
		http.Redirect(res, req, "/blood/edit/"+strconv.Itoa(readingId), http.StatusSeeOther)
		return
	}

	sess.SetErrorAndSuccessMessages(appData)
	res.Header().Set("Content-type", "text/html")
	appData.LayoutData["PageTitle"] = "Health - Edit Blood Pressure Reading"
	appData.LayoutData["User"] = user
	appData.ViewData["Reading"] = reading
	if err := renderView("views/blood/edit.html", res); err != nil {
		serveViewError(res, err)
	}
	sess.CleanupErrorAndSuccessMessages(appData)
}

func BloodDelete(res http.ResponseWriter, req *http.Request) {
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
	readingId, err := getId(req.URL.Path)
	if err != nil {
		serve404(res, req)
		return
	}
	user := sess.GetUserFromSession()
	reading, err := db.GetBloodReading(dbs, user.Id, int64(readingId))
	if err != nil {
		serve500(res, req, err.Error())
		return
	}
	if err = db.DeleteBloodReading(dbs, user.Id, reading.Id); err != nil {
		sess.Set("errMsg", err.Error())
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	sess.Set("okMsg", "Blood pressure reading has been deleted!")
	http.Redirect(res, req, "/", http.StatusSeeOther)
}
