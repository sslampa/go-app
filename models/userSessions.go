package models

import (
	"database/sql"
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

// DeleteUserSession deletes a session
func DeleteUserSession(sid string) {
	userSessionDelete := "DELETE FROM user_sessions WHERE session_id = $1"

	_, err := DB.Exec(userSessionDelete, sid)
	if err != nil {
		log.Fatal(err)
	}
}

// UserLoggedIn checks for session and gets user
func UserLoggedIn(r *http.Request) User {
	u := User{}
	c, err := r.Cookie("session")
	if err != nil {
		return u
	}

	query := "SELECT * FROM users WHERE id = (SELECT user_id FROM user_sessions WHERE session_id = $1)"
	row := DB.QueryRow(query, c.Value)
	err = row.Scan(&u.ID, &u.Username, &u.Password, &u.FirstName, &u.LastName)
	if err == sql.ErrNoRows {
		return User{}
	}
	return u
}
