package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/mtslzr/pokeapi-go"
)

type PokemonDetailsData struct {
	Pokemon    any
	Evolutions any
}

func landingPageHandler(w http.ResponseWriter, r *http.Request) {
	pokemonData, err := pokeapi.Resource("pokemon", 0, 15)
	if err != nil {
		log.Panic("error fetching pokeapi")
	}
	templ := template.Must(template.ParseFiles("./templates/index.html"))
	templ.Execute(w, pokemonData)
}

func pokemonPageHandler(w http.ResponseWriter, r *http.Request) {
	pokemonName := r.URL.Path[len("/pokemon/"):]
	if pokemonName == "" {
		http.Error(w, "Pokemon name is missing", http.StatusBadRequest)
		return
	}

	pokemon, err := pokeapi.Pokemon(pokemonName)
	if err != nil {
		http.Error(w, "Pokemon not found", http.StatusNotFound)
		return
	}
	evolutions, err := pokeapi.EvolutionChain("1")
	if err != nil {
		http.Error(w, "Evolution chain not found", http.StatusNotFound)
		return
	}

	pokemonDetailsData := PokemonDetailsData{
		Pokemon:    pokemon,
		Evolutions: evolutions,
	}

	templ := template.Must(template.ParseFiles("./templates/pokemon.html"))
	templ.Execute(w, pokemonDetailsData)
}

func main() {
	http.HandleFunc("GET /{$}", landingPageHandler)
	http.HandleFunc("GET /pokemon/{name}", pokemonPageHandler)
	fmt.Println("Server running on port 8000")
	http.ListenAndServe(":8000", nil)
}
