package handlers

import (
	"log"
	"net/http"

	"github.com/sslampa/go-app/models"
	"github.com/sslampa/go-app/utility"
)

// Page holds data to send to view
type Page struct {
	Message   string
	User      models.User
	UsersData []models.User
}

// IndexHandler shows index page
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	var p Page
	loggedIn := models.UserLoggedIn(r)

	tpl := utility.MakeTemplate()
	tpl.ParseFiles("./templates/index.gohtml")
	value := utility.GetFlash(w, r, "flash", "/")

	p.Message = value
	p.User = loggedIn

	err := tpl.ExecuteTemplate(w, "base.gohtml", p)
	if err != nil {
		log.Fatalln(err)
	}
}
