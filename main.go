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
)

type LayoutData struct {
	PageTitle string
}

var layoutData = LayoutData{
	PageTitle: "Health",
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/register", register)
	http.HandleFunc("/reset", resetPassword)
	http.HandleFunc("/resetLink", resetPasswordLink)
	http.HandleFunc("/blood/add", bloodAdd)
	http.HandleFunc("/blood/all", bloodAll)
	http.HandleFunc("/blood/details/", bloodDetails)
	http.HandleFunc("/blood/edit/", bloodEdit)
	http.HandleFunc("/blood/delete/", bloodDelete)
	http.HandleFunc("/weight/add", weightAdd)
	http.HandleFunc("/weight/all", weightAll)
	http.HandleFunc("/weight/edit/", weightEdit)
	http.HandleFunc("/weight/delete/", weightDelete)
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
	if !loggedIn(res, req) {
		http.Redirect(res, req, "/login", http.StatusSeeOther)
		return
	}
	layoutData.PageTitle = "Health - Dashboard"
	log.Println(sessionData)
	res.Header().Set("Content-type", "text/html")
	if err := renderView("views/index.html", res); err != nil {
		serveViewError(res, err)
	}
}

func renderView(name string, out io.Writer) error {
	var templates = []string{
		"views/app.html",
		// "views/partials/blood_readings.html",
		// "views/partials/flash_messages.html",
		// "views/partials/weight_entries.html",
		"views/partials/footer.html",
		"views/partials/nav.html",
	}
	templates = append(templates, name)
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return fmt.Errorf("error parsing template: %s", err)
	}
	err = tmpl.ExecuteTemplate(out, "app", layoutData)
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

func serve404(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "text/html")
	res.WriteHeader(http.StatusNotFound)
	res.Write([]byte("<p>404 page was not found</p>"))
}

func serveViewError(res http.ResponseWriter, err error) {
	res.WriteHeader(http.StatusInternalServerError)
	res.Write([]byte(fmt.Sprintf("error rendering view: %s", err)))
}
