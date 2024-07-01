package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

type PokemonDetailsData struct {
	Evolutions structs.EvolutionChain
	Pokemon    structs.Pokemon
}

func landingPageHandler(w http.ResponseWriter, r *http.Request) {
	pokemonData, err := pokeapi.Resource("pokemon", 0, 905)
	if err != nil {
		log.Panic("error fetching pokeapi")
	}
	templ := template.Must(template.New("index.html").Funcs(template.FuncMap{
		"extractId": extractId,
	}).ParseFiles("./templates/index.html"))
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
	pokemonSpecies, err := pokeapi.PokemonSpecies(pokemonName)
	if err != nil {
		http.Error(w, "Pokemon species not found", http.StatusNotFound)
		return
	}
	evolutionId, err := extractId(pokemonSpecies.EvolutionChain.URL)
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
		"incrementId":         incrementId,
		"decrementId":         decrementId,
		"extractId":           extractId,
		"preparePokemonImage": preparePokemonImage,
	}).ParseFiles("./templates/pokemon.html"))

	templ.Execute(w, pokemonDetailsData)
}

func main() {
	http.HandleFunc("GET /{$}", landingPageHandler)
	http.HandleFunc("GET /pokemon/{name}", pokemonPageHandler)
	fmt.Println("Server running on port 8000")
	http.ListenAndServe(":8000", nil)
}

func extractId(urlString string) (string, error) {
	u, err := url.Parse(urlString)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return "", err
	}
	path := u.Path
	parts := strings.Split(path, "/")
	if len(parts) > 1 {
		number := parts[len(parts)-2]
		return number, nil
	} else {
		fmt.Println("URL path format is incorrect")
	}
	return "", nil
}

func incrementId(id int) int {
	return id + 1
}

func decrementId(id int) int {
	return id - 1
}

func preparePokemonImage(id int) string {
	if id < 10 {
		return fmt.Sprintf("00%d", id)
	}
	if id < 100 && id >= 10 {
		return fmt.Sprintf("0%d", id)
	}
	return fmt.Sprintf("%d", id)
}
