package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
)

func main() {
	port := flag.String("p", "8080", "port to serve on")
	flag.Parse()

	http.HandleFunc("/", sendIndex)
	http.HandleFunc("/user", sendUser)
	http.Handle("/stylesheets/", http.StripPrefix("/stylesheets/", http.FileServer(http.Dir("stylesheets"))))

	log.Printf("Serving on HTTP port: %s\n", *port)
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
