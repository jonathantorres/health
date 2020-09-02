package main

import (
	"net/http"
	"strconv"
	"time"
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
	Text  string
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
		Text:  text,
		Class: class,
	}
}

func (blood *BloodReading) SqlDate() string {
	return blood.Date[:10]
}

func bloodAdd(res http.ResponseWriter, req *http.Request) {
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
	user := getUserFromSession(session)

	if req.Method == "POST" {
		req.ParseForm()
		nowDate := time.Now()
		sys, _ := strconv.Atoi(req.PostForm["sys"][0])
		dia, _ := strconv.Atoi(req.PostForm["dia"][0])
		pulse, _ := strconv.Atoi(req.PostForm["pulse"][0])
		date := req.PostForm["reading-date"][0]
		date += " " + nowDate.Format("15:04:05")
		if err := createBloodReading(db, user.Id, int32(sys), int32(dia), int32(pulse), date); err != nil {
			session.Set("errMsg", err.Error())
			http.Redirect(res, req, "/blood/add", http.StatusSeeOther)
			return
		}
		session.Set("okMsg", "Blood reading has been created successfully!")
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	setErrorAndSuccessMessages(session)
	res.Header().Set("Content-type", "text/html")
	appData.LayoutData["PageTitle"] = "Health - Blood Pressure Add Reading"
	appData.LayoutData["User"] = getUserFromSession(session)
	if err := renderView("views/blood/add.html", res); err != nil {
		serveViewError(res, err)
	}
	cleanupErrorAndSuccessMessages(session)
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
	db, err := initDb()
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

	user := getUserFromSession(session)
	reading, err := getBloodReading(db, user.Id, int64(readingId))
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

func bloodEdit(res http.ResponseWriter, req *http.Request) {
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
	readingId, err := getId(req.URL.Path)
	if err != nil {
		serve404(res, req)
		return
	}
	user := getUserFromSession(session)
	reading, err := getBloodReading(db, user.Id, int64(readingId))
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
		if err := updateBloodReading(db, user.Id, reading.Id, int32(sys), int32(dia), int32(pulse), date); err != nil {
			session.Set("errMsg", err.Error())
			http.Redirect(res, req, "/blood/edit/"+strconv.Itoa(readingId), http.StatusSeeOther)
			return
		}
		session.Set("okMsg", "Blood reading has been updated successfully!")
		http.Redirect(res, req, "/blood/edit/"+strconv.Itoa(readingId), http.StatusSeeOther)
		return
	}

	setErrorAndSuccessMessages(session)
	res.Header().Set("Content-type", "text/html")
	appData.LayoutData["PageTitle"] = "Health - Edit Blood Pressure Reading"
	appData.LayoutData["User"] = user
	appData.ViewData["Reading"] = reading
	if err := renderView("views/blood/edit.html", res); err != nil {
		serveViewError(res, err)
	}
	cleanupErrorAndSuccessMessages(session)
}

func bloodDelete(res http.ResponseWriter, req *http.Request) {
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
	readingId, err := getId(req.URL.Path)
	if err != nil {
		serve404(res, req)
		return
	}
	user := getUserFromSession(session)
	reading, err := getBloodReading(db, user.Id, int64(readingId))
	if err != nil {
		serve500(res, req, err.Error())
		return
	}
	if err = deleteBloodReading(db, user.Id, reading.Id); err != nil {
		session.Set("errMsg", err.Error())
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	session.Set("okMsg", "Blood pressure reading has been deleted!")
	http.Redirect(res, req, "/", http.StatusSeeOther)
}
