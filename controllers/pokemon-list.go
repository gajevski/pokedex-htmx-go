package controllers

import (
	"log"
	"net/http"
	"pokedex/shared/utils"
	"text/template"

	"github.com/mtslzr/pokeapi-go"
)

func PokemonListHandler(w http.ResponseWriter, r *http.Request) {
	pokemonData, err := pokeapi.Resource("pokemon", 0, 905)
	if err != nil {
		log.Panic("error fetching pokeapi")
	}
	templ := template.Must(template.New("index.html").Funcs(template.FuncMap{
		"extractId": utils.ExtractID,
	}).ParseFiles("../ui/html/index.html"))
	templ.Execute(w, pokemonData)
}
