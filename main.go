package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/sslampa/go-app/handlers"
	"github.com/sslampa/go-app/models"
)

func main() {
	port := flag.String("p", "8080", "port to serve on")
	seed := flag.Bool("s", false, "drop and add seed data")
	flag.Parse()
	if *seed {
		models.Seed()
	}
	log.Printf("Serving on HTTP port: %s\n", *port)

	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/users", handlers.UsersHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/login/create", handlers.CreateLoginHandler)
	http.HandleFunc("/signup/create", handlers.CreateUserHandler)
	http.HandleFunc("/signup", handlers.SignupHandler)
	http.Handle("/stylesheets/", http.StripPrefix("/stylesheets/", http.FileServer(http.Dir("stylesheets"))))

	err := http.ListenAndServe(":"+*port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
