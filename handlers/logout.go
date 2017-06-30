package handlers

import "net/http"

// LogoutHandler does logging out
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	nc := &http.Cookie{Name: "session", MaxAge: -86400, Path: "/"}
	http.SetCookie(w, nc)

	http.Redirect(w, r, "/", 302)
	return
}
