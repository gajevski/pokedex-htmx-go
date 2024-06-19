package main

import (
	"html/template"
	"net/http"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		templ := template.Must(template.ParseFiles("index.html"))
		templ.Execute(w, nil)
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}
