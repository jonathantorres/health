package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jonathantorres/health/internal/auth"
	"github.com/jonathantorres/health/internal/db"
	"github.com/jonathantorres/health/internal/health"
	"github.com/jonathantorres/health/internal/session"
)

type AppData struct {
	LayoutData map[string]interface{}
	ViewData   map[string]interface{}
}

var appData = AppData{
	LayoutData: make(map[string]interface{}),
	ViewData:   make(map[string]interface{}),
}

func main() {
	appData.LayoutData["PageTitle"] = "Health"
	appData.LayoutData["Version"] = version
	appData.LayoutData["Year"] = time.Now().Year()

	http.HandleFunc("/", root)
	http.HandleFunc("/login", auth.Login)
	http.HandleFunc("/logout", auth.Logout)
	http.HandleFunc("/register", auth.Register)
	http.HandleFunc("/reset", auth.ResetPassword)
	http.HandleFunc("/resetLink", auth.ResetPasswordLink)
	http.HandleFunc("/blood/add", health.BloodAdd)
	http.HandleFunc("/blood/all", health.BloodAll)
	http.HandleFunc("/blood/details/", health.BloodDetails)
	http.HandleFunc("/blood/edit/", health.BloodEdit)
	http.HandleFunc("/blood/delete/", health.BloodDelete)
	http.HandleFunc("/weight/add", health.WeightAdd)
	http.HandleFunc("/weight/all", health.WeightAll)
	http.HandleFunc("/weight/edit/", health.WeightEdit)
	http.HandleFunc("/weight/delete/", health.WeightDelete)
	log.Fatal(http.ListenAndServe(":7070", nil))
}

func root(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/" {
		index(res, req)
		return
	}
	serveStaticFile(res, req)
}

func index(res http.ResponseWriter, req *http.Request) {
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
	readings, readingsErr := db.GetBloodReadings(dbs, user.Id)
	entries, entriesErr := db.GetWeightEntries(dbs, user.Id)
	if readingsErr != nil {
		serve500(res, req, readingsErr.Error())
		return
	}
	if entriesErr != nil {
		serve500(res, req, entriesErr.Error())
		return
	}
	maxReadings := 10
	maxEntries := 10
	if len(readings) <= 10 {
		maxReadings = len(readings)
	}
	if len(entries) <= 10 {
		maxEntries = len(entries)
	}

	sess.SetErrorAndSuccessMessages(appData)
	appData.LayoutData["PageTitle"] = "Health - Dashboard"
	appData.LayoutData["User"] = user
	appData.ViewData["Readings"] = readings[:maxReadings]
	appData.ViewData["Entries"] = entries[:maxEntries]
	appData.ViewData["BloodHeading"] = "Blood Pressure Readings"
	appData.ViewData["WeightHeading"] = "Weight Entries"
	res.Header().Set("Content-type", "text/html")
	if err := renderView("views/index.html", res); err != nil {
		serveViewError(res, err)
	}
	sess.CleanupErrorAndSuccessMessages(app)
}

func renderView(name string, out io.Writer) error {
	var templates = []string{
		"views/app.html",
		"views/partials/flash_messages.html",
		"views/partials/footer.html",
		"views/partials/nav.html",
		"views/partials/blood_readings.html",
		"views/partials/weight_entries.html",
	}
	templates = append(templates, name)
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return fmt.Errorf("error parsing template: %s", err)
	}
	err = tmpl.ExecuteTemplate(out, "app", appData)
	if err != nil {
		return fmt.Errorf("error executing template: %s", err)
	}
	return nil
}

func getId(path string) (int, error) {
	pieces := strings.Split(path, "/")
	id, err := strconv.Atoi(pieces[len(pieces)-1])
	if err != nil {
		return 0, err
	}
	if id == 0 {
		return 0, fmt.Errorf("error: zero value")
	}
	return id, nil
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

func serve500(res http.ResponseWriter, req *http.Request, msg string) {
	res.Header().Set("Content-type", "text/html")
	res.WriteHeader(http.StatusInternalServerError)
	res.Write([]byte(msg))
}

func serve404(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "text/html")
	res.WriteHeader(http.StatusNotFound)
	res.Write([]byte("<p>404 page was not found</p>"))
}

func serveViewError(res http.ResponseWriter, err error) {
	res.WriteHeader(http.StatusInternalServerError)
	res.Write([]byte(fmt.Sprintf("error rendering view: %s", err)))
}
