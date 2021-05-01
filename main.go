package main

import (
	"log"
	"net/http"
	"time"

	"github.com/jonathantorres/health/internal/auth"
	"github.com/jonathantorres/health/internal/health"
)

func main() {
	health.App.LayoutData["PageTitle"] = "Health"
	health.App.LayoutData["Version"] = health.Version
	health.App.LayoutData["Year"] = time.Now().Year()

	http.HandleFunc("/", health.Root)
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
