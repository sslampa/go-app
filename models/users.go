package models

import "log"

// User has things
type User struct {
	ID        int
	Username  string
	Password  string
	FirstName string
	LastName  string
}

// InitUsers does stuff
func InitUsers() {
	tableQuery := `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username varchar,
		password varchar,
		first_name varchar,
		last_name varchar)`

	_, err := DB.Exec(tableQuery)
	if err != nil {
		log.Fatal(err)
	}
}

// CreateUser does stuff
func CreateUser(u *User) {
	userInsert := `INSERT INTO users (username, password, first_name, last_name)
	VALUES ($1, $2, $3, $4)`

	_, err := DB.Exec(userInsert, u.Username, u.Password, u.FirstName, u.LastName)
	if err != nil {
		log.Fatal(err)
	}
}

// AllUsers does stuff
func AllUsers() ([]User, error) {
	rows, err := DB.Query("SELECT id, username, password, first_name, last_name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]User, 0)
	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.ID, &user.Username, &user.Password, &user.FirstName, &user.LastName)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
