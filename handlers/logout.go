package handlers

import (
	"log"
	"net/http"

	"github.com/sslampa/go-app/models"
)

// LogoutHandler does logging out
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil {
		log.Fatal(err)
	}
	nc := &http.Cookie{Name: "session", MaxAge: -86400, Path: "/"}
	http.SetCookie(w, nc)

	models.DeleteUserSession(c.Value)
	http.Redirect(w, r, "/", http.StatusSeeOther)
	return
}
