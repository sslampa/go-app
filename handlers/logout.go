package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(session)

	nc := &http.Cookie{Name: "session", MaxAge: -86400, Path: "/"}
	http.SetCookie(w, nc)

	http.Redirect(w, r, "/", 302)
	return
}
