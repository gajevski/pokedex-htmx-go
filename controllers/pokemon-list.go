package pokemonList

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"text/template"

	"github.com/mtslzr/pokeapi-go"
)

func Greet() {
	fmt.Println("Hello from pokemonList")
}

func PokemonListHandler(w http.ResponseWriter, r *http.Request) {
	pokemonData, err := pokeapi.Resource("pokemon", 0, 905)
	if err != nil {
		log.Panic("error fetching pokeapi")
	}
	templ := template.Must(template.New("index.html").Funcs(template.FuncMap{
		"extractId": extractId,
	}).ParseFiles("./templates/index.html"))
	templ.Execute(w, pokemonData)
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
