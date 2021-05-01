package health

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/jonathantorres/health/internal/db"
	"github.com/jonathantorres/health/internal/session"
)

var App = AppData{
	LayoutData: make(map[string]interface{}),
	ViewData:   make(map[string]interface{}),
}

type AppData struct {
	LayoutData map[string]interface{}
	ViewData   map[string]interface{}
}

func (a *AppData) CleanupErrorAndSuccessMessages(s *session.Session) {
	delete(a.LayoutData, "errMsg")
	delete(a.LayoutData, "okMsg")
	s.Remove("errMsg")
	s.Remove("okMsg")
}

func (a *AppData) SetErrorAndSuccessMessages(s *session.Session) {
	if errMsg, ok := s.Get("errMsg"); ok {
		a.LayoutData["errMsg"] = errMsg
	}
	if okMsg, ok := s.Get("okMsg"); ok {
		a.LayoutData["okMsg"] = okMsg
	}
}

func Index(res http.ResponseWriter, req *http.Request) {
	sess := &session.Session{}
	sess.Start(res, req)
	if !sess.LoggedIn() {
		http.Redirect(res, req, "/login", http.StatusFound)
		return
	}
	dbs, err := db.InitDb()
	if err != nil {
		Serve500(res, req, err.Error())
		return
	}
	user := sess.GetUserFromSession()
	readings, readingsErr := db.GetBloodReadings(dbs, user.Id)
	entries, entriesErr := db.GetWeightEntries(dbs, user.Id)
	if readingsErr != nil {
		Serve500(res, req, readingsErr.Error())
		return
	}
	if entriesErr != nil {
		Serve500(res, req, entriesErr.Error())
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

	App.SetErrorAndSuccessMessages(sess)
	App.LayoutData["PageTitle"] = "Health - Dashboard"
	App.LayoutData["User"] = user
	App.ViewData["Readings"] = readings[:maxReadings]
	App.ViewData["Entries"] = entries[:maxEntries]
	App.ViewData["BloodHeading"] = "Blood Pressure Readings"
	App.ViewData["WeightHeading"] = "Weight Entries"
	res.Header().Set("Content-type", "text/html")
	if err := renderView("views/index.html", res); err != nil {
		ServeViewError(res, err)
	}
	App.CleanupErrorAndSuccessMessages(sess)
}

func RenderView(name string, out io.Writer) error {
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
	err = tmpl.ExecuteTemplate(out, "app", App)
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
