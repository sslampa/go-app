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

	_, err := models.CreateUser(&u)
	if err != nil {
		log.Fatal(err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
