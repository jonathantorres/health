package main

import (
	"net/http"
)

func login(res http.ResponseWriter, req *http.Request) {
	if loggedIn(res, req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	res.Header().Set("Content-type", "text/html")
	layoutData.PageTitle = "Health - Login"
	if err := renderView("views/login.html", res); err != nil {
		serveViewError(res, err)
	}
}

func logout(res http.ResponseWriter, req *http.Request) {
	if !loggedIn(res, req) {
		http.Redirect(res, req, "/login", http.StatusSeeOther)
		return
	}
	res.Header().Set("Content-type", "text/html")
	res.Write([]byte("logout page"))
}

func register(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "text/html")
	layoutData.PageTitle = "Health - Register"
	if err := renderView("views/register.html", res); err != nil {
		serveViewError(res, err)
	}
}

func resetPassword(res http.ResponseWriter, req *http.Request) {
	layoutData.PageTitle = "Health - Reset Password"
	if err := renderView("views/reset_password.html", res); err != nil {
		serveViewError(res, err)
	}
}

func resetPasswordLink(res http.ResponseWriter, req *http.Request) {
	layoutData.PageTitle = "Health - Reset Password"
	if err := renderView("views/reset_password_email.html", res); err != nil {
		serveViewError(res, err)
	}
}

func loggedIn(res http.ResponseWriter, req *http.Request) bool {
	session := &Session{}
	session.Start(res, req)
	if _, ok := session.Get("user"); !ok {
		return false
	}
	// todo: use the user node here
	return true
}
