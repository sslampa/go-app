package models

import "log"

// UserSession holds the sessions
type UserSession struct {
	SessionID int
	UserID    int
}

// InitUserSessions creates the UserSessions table in the db
func InitUserSessions() {
	tableQuery := `CREATE TABLE IF NOT EXISTS user_sessions (
    SessionID SERIAL PRIMARY KEY,
    UserID INT REFERENCES users(id))`

	_, err := DB.Exec(tableQuery)
	if err != nil {
		log.Fatal(err)
	}
}
