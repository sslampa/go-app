package utility

import "html/template"

// MakeTemplate creates the base template for every page
func MakeTemplate() *template.Template {
	tpl := template.Must(template.ParseFiles("./templates/base.gohtml",
		"./templates/footer.gohtml", "./templates/navbar.gohtml"))
	return tpl
}
