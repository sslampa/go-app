package models

import (
	"log"
	"net/http"
)

// UserSession holds the sessions
type UserSession struct {
	SessionID string
	UserID    int
}

// InitUserSessions creates the UserSessions table in the db
func InitUserSessions() {
	tableQuery := `CREATE TABLE IF NOT EXISTS user_sessions (
    session_id TEXT PRIMARY KEY,
    user_id INT REFERENCES users(id))`

	_, err := DB.Exec(tableQuery)
	if err != nil {
		log.Fatal(err)
	}
}

// CreateUserSession makes a session
func CreateUserSession(us *UserSession) {
	userSessionInsert := "INSERT INTO user_sessions (session_id, user_id) VALUES ($1, $2)"

	_, err := DB.Exec(userSessionInsert, us.SessionID, us.UserID)
	if err != nil {
		log.Fatal(err)
	}
}

// UserLoggedIn checks for session and gets user
func UserLoggedIn(r *http.Request) []User {
	c, err := r.Cookie("session")
	if err != nil {
		log.Fatal(err)
	}
	query := "SELECT * FROM users WHERE id = (SELECT user_id FROM user_sessions WHERE session_id = $1)"
	rows, err := DB.Query(query, c.Value)
	if err != nil {
		log.Fatal(err)
	}

	users := make([]User, 0)
	for rows.Next() {
		u := User{}
		err = rows.Scan(&u.ID, &u.Username, &u.Password, &u.FirstName, &u.LastName)
		if err != nil {
			return users
		}
		users = append(users, u)
	}
	return users
}
