package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Pokemon struct {
	Type string
	Name string
}

func main() {
	data := map[string][]Pokemon{
		"Pokemons": {
			Pokemon{Type: "Fire", Name: "Charizard"},
			Pokemon{Type: "Grass", Name: "Bulbasaur"},
		},
	}
	handler := func(w http.ResponseWriter, r *http.Request) {
		templ := template.Must(template.ParseFiles("./templates/index.html"))
		templ.Execute(w, data)
	}

	http.HandleFunc("/", handler)
	fmt.Println("Server running on port 8000")
	http.ListenAndServe(":8000", nil)
}
