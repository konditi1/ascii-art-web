package main

import (
	"fmt"
	"net/http"

	"github.com/konditi1/ascii-art-web/internal/handlers"
)

func main() {

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/about", handlers.AboutHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../static"))))

	fmt.Println("Starting server on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
