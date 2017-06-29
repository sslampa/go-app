package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/sslampa/go-app/models"
	"github.com/sslampa/go-app/utility"
)

// Page Holds what pages should look like
type Page struct {
	Body  *template.Template
	Flash string
}

// CreateUserHandler signs up a new user
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
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

	if len(u.Password) < 6 {
		flashMessage := "Password is not 6 characters or more"
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

// SignupHandler shows sign up page
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("./templates/base.gohtml", "./templates/signup.gohtml", "./templates/navbar.gohtml", "./templates/footer.gohtml"))
	value := utility.GetFlash(w, r, "flash", "/signup")

	err := tpl.ExecuteTemplate(w, "base.gohtml", value)
	if err != nil {
		log.Fatalln(err)
	}
}
