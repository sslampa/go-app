package models

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Init(user, pass string) {
	dbInfo := fmt.Sprintf("user=%s password=%s dbname=go_app", user, pass)
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("You connected to your database")
	InitUsers(db)
}
