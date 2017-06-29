package handlers

import (
	"log"
	"net/http"

	"github.com/sslampa/go-app/utility"
)

// IndexHandler shows index page
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tpl := utility.MakeTemplate()
	tpl.ParseFiles("./templates/index.gohtml")
	value := utility.GetFlash(w, r, "flash", "/")

	err := tpl.ExecuteTemplate(w, "base.gohtml", value)
	if err != nil {
		log.Fatalln(err)
	}
}
