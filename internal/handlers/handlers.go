package handlers

import (
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseGlob("../templates/*.html"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()

		argsPassed := r.FormValue("text")
		bannerFile := r.FormValue("file")
		colorSelect := r.FormValue("selectColor")

		var colorChoice string

		if colorSelect == "yes" {
			colorChoice = r.FormValue("color")
		} else if colorSelect == "no" {
			colorChoice = ""
		}

		data := map[string]interface{}{
			"argsPassed":  argsPassed,
			"bannerFile":  bannerFile,
			"colorChoice": colorChoice,
		}

		tmpl.ExecuteTemplate(w, "home.html", data)
	} else {
		tmpl.ExecuteTemplate(w, "home.html", nil)
	}

}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "about.html", nil)
}
