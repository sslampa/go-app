package models

import (
	"fmt"
	"log"
)

// User holds the values for a user of the site
type User struct {
	ID        int
	Username  string
	Password  string
	FirstName string
	LastName  string
}

// InitUsers creates the Users table in the db
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

// CreateUser adds a user into the Users db
func CreateUser(u *User) (User, error) {
	userInsert := "INSERT INTO users (username, password, first_name, last_name) VALUES ($1,$2,$3,$4)"

	result, err := DB.Exec(userInsert, u.Username, u.Password, u.FirstName, u.LastName)
	if err != nil {
		log.Fatal(err)
	}
	id, _ := result.LastInsertId()
	strID := fmt.Sprintf("%v", id)

	createdUser, err := FindUser(strID, "id")
	if err != nil {
		log.Fatal(err)
	}

	return createdUser, nil
}

// FindUser returns a single user found by either username or id
func FindUser(value string, column string) (User, error) {
	var userString string
	switch column {
	case "username":
		userString = "SELECT * FROM users WHERE username = $1"
	case "id":
		userString = "SELECT * FROM users WHERE id = $1"
	}
	rows, err := DB.Query(userString, value)
	if err != nil {
		log.Fatal(err)
	}

	user := User{}
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Username, &user.Password, &user.FirstName, &user.LastName)
		if err != nil {
			return user, err
		}
	}

	return user, nil
}

// AllUsers returns all users
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

// CheckUsername checks to see if there's a unique username
func CheckUsername(u *User) bool {
	user, _ := FindUser(u.Username, "username")

	if user != (User{}) {
		return true
	}

	return false
}

// CheckPassword checks to see if the password is valid
func CheckPassword(pw, pwConfirm string) bool {
	if pw != pwConfirm || len(pw) < 6 {
		return true
	}

	return false
}
