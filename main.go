package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Fprintf(w, "Page not found!")
	} else {
		template.Execute(w, nil)
	}
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("The server is listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
