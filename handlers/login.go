package handlers

import (
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	uuid "github.com/satori/go.uuid"
	"github.com/sslampa/go-app/models"
	"github.com/sslampa/go-app/utility"
)

// LoginHandler shows log in page
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var p Page
	loggedIn := models.UserLoggedIn(r)

	tpl := utility.MakeTemplate()
	tpl.ParseFiles("./templates/login.gohtml")
	value := utility.GetFlash(w, r, "flash", "/login")

	p.Message = value
	p.Users = loggedIn

	err := tpl.ExecuteTemplate(w, "base.gohtml", p)
	if err != nil {
		log.Fatal(err)
	}
}

// CreateLoginHandler creates stuff
func CreateLoginHandler(w http.ResponseWriter, r *http.Request) {
	var u models.User
	username := r.FormValue("username")
	password := r.FormValue("password")

	u, err := models.FindUser(username, "username")
	if err != nil {
		log.Fatal(err)
	}
	if !models.CheckUsername(&u) {
		utility.SetFlash(w, "flash", "Username/Password do not match", "/login")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		utility.SetFlash(w, "flash", "Username/Password do not match", "/login")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	sID := uuid.NewV4()
	var us models.UserSession
	us.SessionID = sID.String()
	us.UserID = u.ID
	models.CreateUserSession(&us)

	c := &http.Cookie{Name: "session", Value: sID.String(), MaxAge: 86400, Path: "/"}
	http.SetCookie(w, c)
	utility.SetFlash(w, "flash", "You have succesfully logged in", "/")
	http.Redirect(w, r, "/", 301)
}
