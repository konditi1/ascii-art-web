package main

import (
	"fmt"
	"net/http"

	"github.com/konditi1/ascii-art-web/internal/handlers"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handlers.PostHandler(w, r)	
		} else {
			http.ServeFile(w, r, "../templates/home.html")
		}
	})
	http.HandleFunc("/about", handlers.AboutHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../static"))))

	fmt.Println("Starting server on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
