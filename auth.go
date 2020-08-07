package main

import (
	"fmt"
	"net/http"
)

func login(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "text/html")
	layoutData.PageTitle = "Health - Login"
	if err := renderView("views/login.html", res); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(fmt.Sprintf("error rendering view: %s", err)))
	}
}

func logout(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "text/html")
	res.Write([]byte("logout page"))
}

func register(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "text/html")
	res.Write([]byte("register page"))
}

func resetPassword(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "text/html")
	res.Write([]byte("resetPassword page"))
}

func resetPasswordLink(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "text/html")
	res.Write([]byte("resetPasswordLink page"))
}
