package main

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
)

func login(res http.ResponseWriter, req *http.Request) {
	session := &Session{}
	session.Start(res, req)
	if loggedIn(session) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	db, err := initDb()
	if err != nil {
		serve500(res, req, err.Error())
		return
	}
	if req.Method == "POST" {
		req.ParseForm()
		email := req.PostForm["email"][0]
		pass := req.PostForm["password"][0]
		if err = authenticate(db, res, req, session, email, pass); err != nil {
			session.Set("errMsg", err.Error())
			http.Redirect(res, req, "/login", http.StatusSeeOther)
			return
		}
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	setErrorAndSuccessMessages(session)
	res.Header().Set("Content-type", "text/html")
	appData.LayoutData["PageTitle"] = "Health - Login"
	if err := renderView("views/login.html", res); err != nil {
		serveViewError(res, err)
	}
	cleanupErrorAndSuccessMessages(session)
}

func authenticate(db *sql.DB, res http.ResponseWriter, req *http.Request, session *Session, email, pass string) error {
	sql := `
		SELECT email
		FROM users
		WHERE email = ?
		LIMIT 1
	`
	rows, err := db.Query(sql, email)
	if err != nil {
		log.Printf("err: %s", err)
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var userEmail string
		if err := rows.Scan(&userEmail); err != nil {
			log.Printf("%s", err)
			break
		}
		if userEmail == email {
			session.Start(res, req)
			session.Set("user", userEmail)
			return nil
		}
	}
	return errors.New("Invalid credentials")
}

func logout(res http.ResponseWriter, req *http.Request) {
	session := &Session{}
	session.Start(res, req)
	if !loggedIn(session) {
		http.Redirect(res, req, "/login", http.StatusSeeOther)
		return
	}
	session.Destroy(res)
	http.Redirect(res, req, "/login", http.StatusSeeOther)
	return
}

func register(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "text/html")
	appData.LayoutData["PageTitle"] = "Health - Register"
	if err := renderView("views/register.html", res); err != nil {
		serveViewError(res, err)
	}
}

func resetPassword(res http.ResponseWriter, req *http.Request) {
	appData.LayoutData["PageTitle"] = "Health - Reset Password"
	if err := renderView("views/reset_password.html", res); err != nil {
		serveViewError(res, err)
	}
}

func resetPasswordLink(res http.ResponseWriter, req *http.Request) {
	appData.LayoutData["PageTitle"] = "Health - Reset Password"
	if err := renderView("views/reset_password_email.html", res); err != nil {
		serveViewError(res, err)
	}
}

func loggedIn(session *Session) bool {
	_, ok := session.Get("user")
	if !ok {
		return false
	}
	// todo: use the user node here
	return true
}
