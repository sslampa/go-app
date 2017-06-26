package models

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
)

var DB *sql.DB

type Specification struct {
	User string
	Pass string
}

func init() {
	fmt.Println("Here")
	var s Specification
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
}
