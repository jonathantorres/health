package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"io/ioutil"
)

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/login", login)
	log.Fatal(http.ListenAndServe(":7070", nil))
}

func root(res http.ResponseWriter, req *http.Request) {
	if (req.URL.Path == "/") {
		index(res, req)
		return
	}
	http.FileServer(http.Dir("./public")).ServeHTTP(res, req)
}

func index(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "text/html")
	res.Write([]byte("do your own thing here"))
}

func login(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "text/html")
	loginHtml, err := ioutil.ReadFile("views/login.html")
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(fmt.Sprintf("error reading file: %s", err)))
		return
	}
	tmpl, err := template.New("login").Parse(string(loginHtml))
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(fmt.Sprintf("error parsing template: %s", err)))
		return
	}
	err = tmpl.Execute(res, nil)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(fmt.Sprintf("error executing template: %s", err)))
	}
}
