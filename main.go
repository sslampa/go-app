package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"

	"github.com/sslampa/go-app/models"
)

func main() {
	port := flag.String("p", "8080", "port to serve on")
	flag.Parse()
	log.Printf("Serving on HTTP port: %s\n", *port)

	http.HandleFunc("/", sendIndex)
	http.HandleFunc("/user", sendUser)
	http.HandleFunc("/users", sendUsers)
	http.HandleFunc("/login", sendLogin)
	http.HandleFunc("/signup", sendSignup)
	http.Handle("/stylesheets/", http.StripPrefix("/stylesheets/", http.FileServer(http.Dir("stylesheets"))))

	err := http.ListenAndServe(":"+*port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func sendIndex(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseGlob("./templates/*.gohtml"))

	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func sendUser(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseGlob("./templates/*.gohtml"))

	err := tpl.ExecuteTemplate(w, "user.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func sendUsers(w http.ResponseWriter, r *http.Request) {
	users, err := models.AllUsers()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	tpl := template.Must(template.ParseGlob("./templates/*.gohtml"))

	err = tpl.ExecuteTemplate(w, "users.gohtml", users)
	if err != nil {
		log.Fatal(err)
	}
}

func sendLogin(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseGlob("./templates/*.gohtml"))

	err := tpl.ExecuteTemplate(w, "login.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func sendSignup(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseGlob("./templates/*.gohtml"))

	err := tpl.ExecuteTemplate(w, "signup.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
