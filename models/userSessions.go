package models

import "log"

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
