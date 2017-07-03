package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sslampa/go-app/models"
	"github.com/sslampa/go-app/utility"
)

// UserShowHandler does things
func UserShowHandler(w http.ResponseWriter, r *http.Request) {
	var p Page
	loggedIn := models.UserLoggedIn(r)

	// path := r.URL.Path
	// re := regexp.MustCompile("^.*/user/([0-9]+)")
	// re.FindStringSubmatch(path)
	tpl := utility.MakeTemplate()
	tpl.ParseFiles("./templates/user.gohtml")

	p.User = loggedIn
	fmt.Println("Runs")

	err := tpl.ExecuteTemplate(w, "base.gohtml", p)
	if err != nil {
		log.Fatal(err)
	}
}
