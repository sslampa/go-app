package handlers

import (
	"log"
	"net/http"

	"github.com/sslampa/go-app/utility"
)

// LoginHandler shows log in page
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	tpl := utility.MakeTemplate()
	tpl.ParseFiles("./templates/login.gohtml")

	err := tpl.ExecuteTemplate(w, "base.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// LoginCreateHandler creates stuff
// func LoginCreateHandler(w http.ResponseWriter, r *http.Request) {
// 	var u models.User
// 	username := r.FormValue("username")
// 	password := r.FormValue("password")
//
// 	u, err := models.FindUser(username, "username")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if !models.CheckUsername(&u) {
// 		http.Redirect(w, r, "/login", http.StatusSeeOther)
// 	}
// 	if
// }
