package models

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
	// Used to connect to server
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
}

// Seed the db
func Seed() {
	dropQuery := `DROP TABLE users`
	_, err := DB.Exec(dropQuery)
	if err != nil {
		log.Fatal(err)
	}
	InitUsers()
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
	err = CreateUser(&user1)
	if err != nil {
		log.Fatal(err)
	}
	err = CreateUser(&user2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Seed file ran")
}
