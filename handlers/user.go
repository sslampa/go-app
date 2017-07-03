package handlers

import (
	"log"
	"net/http"
	"regexp"

	"github.com/sslampa/go-app/models"
	"github.com/sslampa/go-app/utility"
)

// UserShowHandler does things
func UserShowHandler(w http.ResponseWriter, r *http.Request) {
	var p Page
	loggedIn := models.UserLoggedIn(r)

	path := r.URL.Path
	re := regexp.MustCompile("^.*/user/([0-9]+)")
	match := re.FindStringSubmatch(path)
	u, err := models.FindUser(match[1], "id")
	if err != nil {
		log.Fatal(err)
	}

	tpl := utility.MakeTemplate()
	tpl.ParseFiles("./templates/user.gohtml")

	p.User = loggedIn
	p.UserData = u

	err = tpl.ExecuteTemplate(w, "base.gohtml", p)
	if err != nil {
		log.Fatal(err)
	}
}
