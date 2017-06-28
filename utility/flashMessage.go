package utility

import "net/http"

// SetFlash does things
func SetFlash(w http.ResponseWriter, name, value, path string) {
	c := &http.Cookie{Name: name, Value: value, Path: path}
	http.SetCookie(w, c)
}

// GetFlash gets and removes cookie
func GetFlash(w http.ResponseWriter, r *http.Request, name string) string {
	c, err := r.Cookie(name)
	if err != nil {
		return ""
	}

	message := c.Value
	nc := &http.Cookie{Name: name, MaxAge: -1}
	http.SetCookie(w, nc)

	return message
}
