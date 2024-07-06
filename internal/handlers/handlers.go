package handlers

import (
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseGlob("../templates/*.html"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// r.ParseForm()
	// parsedArgs := r.FormValue("text")

	tmpl.ExecuteTemplate(w, "home.html", nil)
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "about.html", nil)
}
func PostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	userInput := r.FormValue("text")


	w.Write([]byte(userInput))
}
