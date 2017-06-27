package models

import (
	"fmt"
	"log"
)

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
		username VARCHAR UNIQUE ,
		password VARCHAR,
		first_name VARCHAR,
		last_name VARCHAR)`

	_, err := DB.Exec(tableQuery)
	if err != nil {
		log.Fatal(err)
	}
}

// CreateUser does stuff
func CreateUser(u *User) error {
	user, _ := FindUser(u.Username, "username")
	if user != (User{}) {
		return fmt.Errorf("Username already exists")
	}

	userInsert := `INSERT INTO users (username, password, first_name, last_name)
		VALUES ($1, $2, $3, $4)`

	_, err := DB.Exec(userInsert, u.Username, u.Password, u.FirstName, u.LastName)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

// FindUser finds stuff
func FindUser(value string, column string) (User, error) {
	var userString string
	switch column {
	case "username":
		userString = `SELECT * FROM users WHERE username = $1`
	case "id":
		userString = `SELECT * FROM users WHERE id = $1`
	}

	rows, err := DB.Query(userString, value)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	user := User{}
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Username, &user.Password, &user.FirstName, &user.LastName)
		if err != nil {
			return user, err
		}
	}

	return user, nil
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
