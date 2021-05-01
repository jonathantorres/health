package auth

import (
	"database/sql"
	"errors"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/jonathantorres/health/internal/db"
	"github.com/jonathantorres/health/internal/health"
	"github.com/jonathantorres/health/internal/session"
)

func Login(res http.ResponseWriter, req *http.Request) {
	sess := &session.Session{}
	sess.Start(res, req)
	if sess.LoggedIn() {
		http.Redirect(res, req, "/", http.StatusFound)
		return
	}
	dbs, err := db.InitDb()
	if err != nil {
		health.Serve500(res, req, err.Error())
		return
	}
	if req.Method == "POST" {
		req.ParseForm()
		email := req.PostForm["email"][0]
		pass := req.PostForm["password"][0]
		if err = Authenticate(dbs, res, req, sess, email, pass); err != nil {
			sess.Set("errMsg", err.Error())
			http.Redirect(res, req, "/login", http.StatusSeeOther)
			return
		}
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	health.App.SetErrorAndSuccessMessages(sess)
	res.Header().Set("Content-type", "text/html")
	health.App.LayoutData["PageTitle"] = "Health - Login"
	if err := health.RenderView("views/login.html", res); err != nil {
		health.ServeViewError(res, err)
	}
	health.App.CleanupErrorAndSuccessMessages(sess)
}

func Authenticate(dbs *sql.DB, res http.ResponseWriter, req *http.Request, sess *session.Session, email, pass string) error {
	sql := `
		SELECT id, name, last_name, email, password
		FROM users
		WHERE email = ?
		LIMIT 1
	`
	rows, err := dbs.Query(sql, email)
	if err != nil {
		log.Printf("err: %s", err)
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var userId int64
		var userName string
		var userLastName string
		var userEmail string
		var userPass string
		if err := rows.Scan(&userId, &userName, &userLastName, &userEmail, &userPass); err != nil {
			log.Printf("%s", err)
			break
		}
		if userEmail == email && bcrypt.CompareHashAndPassword([]byte(userPass), []byte(pass)) == nil {
			user := session.User{
				Id:       userId,
				Name:     userName,
				LastName: userLastName,
				Email:    userEmail,
			}
			sess.Start(res, req)
			sess.Set("user", user)
			return nil
		}
	}
	return errors.New("Invalid credentials")
}

func Logout(res http.ResponseWriter, req *http.Request) {
	sess := &session.Session{}
	sess.Start(res, req)
	if !sess.LoggedIn() {
		http.Redirect(res, req, "/login", http.StatusFound)
		return
	}
	sess.Destroy(res)
	http.Redirect(res, req, "/login", http.StatusSeeOther)
}

func Register(res http.ResponseWriter, req *http.Request) {
	sess := &session.Session{}
	sess.Start(res, req)
	if sess.LoggedIn() {
		http.Redirect(res, req, "/", http.StatusFound)
		return
	}
	dbs, err := db.InitDb()
	if err != nil {
		health.Serve500(res, req, err.Error())
		return
	}
	if req.Method == "POST" {
		req.ParseForm()
		name := req.PostForm["name"][0]
		lastName := req.PostForm["last_name"][0]
		email := req.PostForm["email"][0]
		pass := req.PostForm["password"][0]
		passConfirm := req.PostForm["password_confirmation"][0]

		if pass != passConfirm {
			sess.Set("errMsg", "Password and password confirmation must be the same")
			http.Redirect(res, req, "/register", http.StatusSeeOther)
			return
		}
		if err := db.RegisterUser(dbs, name, lastName, email, pass); err != nil {
			sess.Set("errMsg", err.Error())
			http.Redirect(res, req, "/register", http.StatusSeeOther)
			return
		}
		if err = Authenticate(dbs, res, req, sess, email, pass); err != nil {
			sess.Set("errMsg", err.Error())
			http.Redirect(res, req, "/login", http.StatusSeeOther)
			return
		}
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	health.App.SetErrorAndSuccessMessages(sess)
	res.Header().Set("Content-type", "text/html")
	health.App.LayoutData["PageTitle"] = "Health - Register"
	if err := health.RenderView("views/register.html", res); err != nil {
		health.ServeViewError(res, err)
	}
	health.App.CleanupErrorAndSuccessMessages(sess)
}

func ResetPassword(res http.ResponseWriter, req *http.Request) {
	health.App.LayoutData["PageTitle"] = "Health - Reset Password"
	if err := health.RenderView("views/reset_password.html", res); err != nil {
		health.ServeViewError(res, err)
	}
}

func ResetPasswordLink(res http.ResponseWriter, req *http.Request) {
	health.App.LayoutData["PageTitle"] = "Health - Reset Password"
	if err := health.RenderView("views/reset_password_email.html", res); err != nil {
		health.ServeViewError(res, err)
	}
}
