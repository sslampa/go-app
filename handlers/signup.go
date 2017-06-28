package handlers

import (
	"log"
	"net/http"

	"github.com/sslampa/go-app/models"
)

// SignupHandler does stuff
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	u := models.User{}
	u.Username = r.FormValue("username")
	u.Password = r.FormValue("password")
	u.FirstName = r.FormValue("first-name")
	u.LastName = r.FormValue("last-name")
	passwordConfirm := r.FormValue("password-confirm")

	if u.Password != passwordConfirm {
		http.Redirect(w, r, "/signup", http.StatusSeeOther)
		return
	}

	_, err := models.CreateUser(&u)
	if err != nil {
		log.Fatal(err)
	}

	c := &http.Cookie{Name: "flash", Value: "You signed up!", Path: "/"}
	http.SetCookie(w, c)
	http.Redirect(w, r, "/", 301)
}
