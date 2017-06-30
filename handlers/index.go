package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sslampa/go-app/models"
	"github.com/sslampa/go-app/utility"
)

// Page holds data to send to view
type Page struct {
	Message string
	Users   []models.User
}

// IndexHandler shows index page
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	var p Page
	loggedIn := models.UserLoggedIn(r)
	fmt.Println(loggedIn)
	tpl := utility.MakeTemplate()
	tpl.ParseFiles("./templates/index.gohtml")
	value := utility.GetFlash(w, r, "flash", "/")

	p.Message = value
	p.Users = loggedIn
	err := tpl.ExecuteTemplate(w, "base.gohtml", p)
	if err != nil {
		log.Fatalln(err)
	}
}
