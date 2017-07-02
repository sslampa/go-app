package handlers

import (
	"log"
	"net/http"

	"github.com/sslampa/go-app/models"
	"github.com/sslampa/go-app/utility"
)

// UsersHandler shows all users page
func UsersHandler(w http.ResponseWriter, r *http.Request) {
	var p Page
	loggedIn := models.UserLoggedIn(r)

	users, err := models.AllUsers()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	tpl := utility.MakeTemplate()
	tpl.ParseFiles("./templates/users.gohtml")

	p.User = loggedIn
	p.UsersData = users

	err = tpl.ExecuteTemplate(w, "base.gohtml", p)
	if err != nil {
		log.Fatal(err)
	}
}
