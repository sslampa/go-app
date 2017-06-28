package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sslampa/go-app/models"
	"github.com/sslampa/go-app/utility"
)

// SignupHandler signs up a new user
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	u := models.User{}
	u.Username = r.FormValue("username")
	u.Password = r.FormValue("password")
	u.FirstName = r.FormValue("first-name")
	u.LastName = r.FormValue("last-name")
	passwordConfirm := r.FormValue("password-confirm")

	if u.Password != passwordConfirm {
		flashMessage := "Passwords do not match"
		utility.SetFlash(w, "flash", flashMessage, "/signup")
		http.Redirect(w, r, "/signup", http.StatusSeeOther)
		return
	}

	if models.CheckUsername(&u) {
		flashMessage := "Username already exists"
		utility.SetFlash(w, "flash", flashMessage, "/signup")
		http.Redirect(w, r, "/signup", http.StatusSeeOther)
		return
	}

	_, err := models.CreateUser(&u)
	if err != nil {
		log.Fatal(err)
	}

	flashMessage := fmt.Sprintf("You signed up, %v!", u.Username)
	utility.SetFlash(w, "flash", flashMessage, "/")
	http.Redirect(w, r, "/", 301)
}
