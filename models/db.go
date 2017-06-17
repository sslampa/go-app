package models

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
)

var db *sql.DB

type Specification struct {
	User string
	Pass string
}

func init() {
	var s Specification
	err := envconfig.Process("db", &s)
	if err != nil {
		log.Fatal(err.Error())
	}

	dbInfo := fmt.Sprintf("user=%s password=%s dbname=go_app", s.User, s.Pass)
	db, err = sql.Open("postgres", dbInfo)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("You connected to your database")
}
