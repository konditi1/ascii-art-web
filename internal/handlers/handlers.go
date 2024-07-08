package handlers

import (
	"html/template"
	"net/http"

	GenerateASCII "github.com/konditi1/ascii-art-web/resourses/asciiOutput"
)

var tmpl = template.Must(template.ParseGlob("../templates/*.html"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}

		argsPassed := r.FormValue("text")
		bannerFile := r.FormValue("file")
		colorSelect := r.FormValue("selectColor")

		var colorChoice string

		if colorSelect == "yes" {
			colorChoice = r.FormValue("color")
		} else if colorSelect == "no" {
			colorChoice = ""
		}

		asciiArt, err := GenerateASCII.GenerateASCII(colorChoice, argsPassed, bannerFile)
		if err != nil {
			http.Error(w, "Error generating ASCII art", http.StatusInternalServerError)
		}

		data := map[string]interface{}{
			"AsciiArt":    asciiArt,
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
