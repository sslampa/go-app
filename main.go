package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"

	"github.com/kelseyhightower/envconfig"
	"github.com/sslampa/go-app/models"
)

type Specification struct {
	User string
	Pass string
}

func main() {
	port := flag.String("p", "8080", "port to serve on")
	flag.Parse()
	log.Printf("Serving on HTTP port: %s\n", *port)

	var s Specification
	err := envconfig.Process("db", &s)
	if err != nil {
		log.Fatal(err.Error())
	}

	models.Init(s.User, s.Pass)

	http.HandleFunc("/", sendIndex)
	http.HandleFunc("/user", sendUser)
	http.HandleFunc("/login", sendLogin)
	http.HandleFunc("/signup", sendSignup)
	http.Handle("/stylesheets/", http.StripPrefix("/stylesheets/", http.FileServer(http.Dir("stylesheets"))))

	err = http.ListenAndServe(":"+*port, nil)
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
