package models

import (
	"database/sql"
	"fmt"
	"log"
)

func InitUsers(db *sql.DB) {
	createUsersTable := `CREATE TABLE IF NOT EXISTS users (
    id integer PRIMARY KEY,
    username varchar,
    password varchar,
    first_name varchar,
    last_name varchar)`

	_, err := db.Exec(createUsersTable)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Users DB created!")
}
