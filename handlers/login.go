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
