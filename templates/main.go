package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

type person struct {
	FirstName string
	LastName  string
}

func (p *person) Add() string {
	return fmt.Sprintf("Hello, %v %v!", p.FirstName, p.LastName)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	people := []person{
		person{
			FirstName: "Stephen",
			LastName:  "Lampa",
		},
		person{
			FirstName: "Joey",
			LastName:  "Lampa",
		},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "main.gohtml", people)
	if err != nil {
		log.Fatalln(err)
	}
}
