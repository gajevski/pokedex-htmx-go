package main

import (
	"fmt"
	"net/http"
	"pokedex/controllers"
)

func main() {
	http.HandleFunc("GET /{$}", controllers.PokemonListHandler)
	http.HandleFunc("GET /pokemon/{name}", controllers.PokemonPageHandler)
	fmt.Println("Server running on port 8000")
	http.ListenAndServe(":8000", nil)
}
