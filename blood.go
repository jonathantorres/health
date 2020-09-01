package main

import (
	"net/http"
)

type BloodReading struct {
	Id        int64
	UserId    int64
	Systolic  int32
	Diastolic int32
	Pulse     int32
	Date      string
}

type BloodSeverity struct {
	Text string
	Class string
}

func (blood *BloodReading) Severity() *BloodSeverity {
	text := "N/A"
	class := "normal"

	if blood.Systolic <= 120 && blood.Diastolic <= 80 {
		text = "Normal"
		class = "primary"
	} else if (blood.Systolic > 120 && blood.Systolic <= 139) || (blood.Diastolic > 80 && blood.Diastolic <= 89) {
		text = "Pre Hypertension"
		class = "warning"
	} else if (blood.Systolic >= 140 && blood.Systolic <= 159) || (blood.Diastolic >= 90 && blood.Diastolic <= 99) {
		text = "Stage 1 Hypertension"
		class = "danger"
	} else if blood.Systolic >= 160 && blood.Diastolic >= 100 {
		text = "Stage 2 Hypertension"
		class = "danger"
	}

	return &BloodSeverity{
		Text: text,
		Class: class,
	}
}

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
	db, err := initDb()
	if err != nil {
		serve500(res, req, err.Error())
		return
	}
	res.Header().Set("Content-type", "text/html")

	user := getUserFromSession(session)
	readings, err := getBloodReadings(db, user.Id)
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
