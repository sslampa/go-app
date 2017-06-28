package utility

import "net/http"

// SetFlash creates a cookie that will be used for a flash message
func SetFlash(w http.ResponseWriter, name, value, path string) {
	c := &http.Cookie{Name: name, Value: value, Path: path}
	http.SetCookie(w, c)
}

// GetFlash gets and removes flash cookie
func GetFlash(w http.ResponseWriter, r *http.Request, name, path string) string {
	c, err := r.Cookie(name)
	if err != nil {
		return ""
	}

	message := c.Value
	nc := &http.Cookie{Name: name, Path: path, MaxAge: -1}
	http.SetCookie(w, nc)

	return message
}
