package models

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
	// Connects to server
	_ "github.com/lib/pq"
)

// DB used to execute DB commands
var DB *sql.DB

type specification struct {
	User string
	Pass string
}

func init() {
	var s specification
	err := envconfig.Process("db", &s)
	if err != nil {
		log.Fatal(err.Error())
	}

	dbInfo := fmt.Sprintf("user=%s password=%s dbname=go_app", s.User, s.Pass)
	DB, err = sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal(err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("You connected to your database")
	InitUsers()
	InitUserSessions()
}

// Seed the db
func Seed() {
	dropQuery := `DROP TABLE users, user_sessions`
	_, err := DB.Exec(dropQuery)
	if err != nil {
		log.Fatal(err)
	}
	InitUsers()
	InitUserSessions()
	user1 := User{
		Username:  "sslampa",
		Password:  "123456",
		FirstName: "Stephen",
		LastName:  "Lampa",
	}
	user2 := User{
		Username:  "tomanistor",
		Password:  "123456",
		FirstName: "Toma",
		LastName:  "Nistor",
	}
	_, _ = CreateUser(&user1)
	_, _ = CreateUser(&user2)

	fmt.Println("Seed file ran")
}
