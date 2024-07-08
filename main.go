package main

import (
	"fmt"
	"net/http"
	"pokedex/controllers"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("GET /{$}", controllers.PokemonListHandler)
	router.HandleFunc("GET /pokemon/{name}", controllers.PokemonPageHandler)
	router.HandleFunc("GET /move/{id}", controllers.MovePageController)
	fmt.Println("Server running on port 8000")
	http.ListenAndServe(":8000", router)
}
