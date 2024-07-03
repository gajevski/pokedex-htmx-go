package controllers

import (
	"net/http"
	"pokedex/shared/utils"
	"text/template"

	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

type PokemonDetailsData struct {
	Evolutions structs.EvolutionChain
	Pokemon    structs.Pokemon
}

func PokemonPageHandler(w http.ResponseWriter, r *http.Request) {
	pokemonName := extractPokemonName(w, r)
	pokemon := getPokemon(pokemonName, w)

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

func extractPokemonName(w http.ResponseWriter, r *http.Request) string {
	pokemonName := r.URL.Path[len("/pokemon/"):]
	if pokemonName == "" {
		http.Error(w, "Pokemon name is missing", http.StatusBadRequest)
		return ""
	}
	return pokemonName
}

func getPokemon(pokemonName string, w http.ResponseWriter) structs.Pokemon {
	pokemon, err := pokeapi.Pokemon(pokemonName)
	if err != nil {
		http.Error(w, "Pokemon not found", http.StatusNotFound)
		return structs.Pokemon{}
	}

	return pokemon
}