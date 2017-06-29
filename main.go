package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"

	"github.com/sslampa/go-app/handlers"
	"github.com/sslampa/go-app/models"
	"github.com/sslampa/go-app/utility"
)

func main() {
	port := flag.String("p", "8080", "port to serve on")
	seed := flag.Bool("s", false, "drop and add seed data")
	flag.Parse()
	if *seed {
		models.Seed()
	}
	log.Printf("Serving on HTTP port: %s\n", *port)

	http.HandleFunc("/", sendIndex)
	http.HandleFunc("/users", sendUsers)
	http.HandleFunc("/login", sendLogin)
	http.HandleFunc("/signup/create", handlers.CreateUserHandler)
	http.HandleFunc("/signup", handlers.SignupHandler)
	http.Handle("/stylesheets/", http.StripPrefix("/stylesheets/", http.FileServer(http.Dir("stylesheets"))))

	err := http.ListenAndServe(":"+*port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func sendIndex(w http.ResponseWriter, r *http.Request) {
	tpl := makeTemplate()
	tpl.ParseFiles("./templates/index.gohtml")
	value := utility.GetFlash(w, r, "flash", "/")

	err := tpl.ExecuteTemplate(w, "base.gohtml", value)
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
	tpl := makeTemplate()
	tpl.ParseFiles("./templates/users.gohtml")

	err = tpl.ExecuteTemplate(w, "base.gohtml", users)
	if err != nil {
		log.Fatal(err)
	}
}

func sendLogin(w http.ResponseWriter, r *http.Request) {
	tpl := makeTemplate()
	tpl.ParseFiles("./templates/login.gohtml")

	err := tpl.ExecuteTemplate(w, "base.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func makeTemplate() *template.Template {
	tpl := template.Must(template.ParseFiles("./templates/base.gohtml",
		"./templates/footer.gohtml", "./templates/navbar.gohtml"))
	return tpl
}
