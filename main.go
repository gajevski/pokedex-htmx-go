package main

import (
	"fmt"
	"html/template"
	"net/http"
	pokemonList "pokedex/controllers"
	"pokedex/shared/utils"

	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

type PokemonDetailsData struct {
	Evolutions structs.EvolutionChain
	Pokemon    structs.Pokemon
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
	pokemonSpecies, err := pokeapi.PokemonSpecies(pokemonName)
	if err != nil {
		http.Error(w, "Pokemon species not found", http.StatusNotFound)
		return
	}
	evolutionId, err := utils.ExtractID(pokemonSpecies.EvolutionChain.URL)
	if err != nil {
		http.Error(w, "Evolution chain URL not found", http.StatusNotFound)
		return
	}
	evolutions, err := pokeapi.EvolutionChain(evolutionId)
	if err != nil {
		http.Error(w, "Evolution chain not found", http.StatusNotFound)
		return
	}

	pokemonDetailsData := PokemonDetailsData{
		Pokemon:    pokemon,
		Evolutions: evolutions,
	}

	templ := template.Must(template.New("pokemon.html").Funcs(template.FuncMap{
		"incrementId":         utils.IncrementID,
		"decrementId":         utils.DecrementID,
		"extractId":           utils.ExtractID,
		"preparePokemonImage": utils.PreparePokemonImage,
	}).ParseFiles("./templates/pokemon.html"))

	templ.Execute(w, pokemonDetailsData)
}

func main() {
	http.HandleFunc("GET /{$}", pokemonList.PokemonListHandler)
	http.HandleFunc("GET /pokemon/{name}", pokemonPageHandler)
	fmt.Println("Server running on port 8000")
	http.ListenAndServe(":8000", nil)
}
