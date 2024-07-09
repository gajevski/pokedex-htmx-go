package controllers

import (
	"net/http"
	"text/template"

	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

func MovePageController(w http.ResponseWriter, r *http.Request) {
	moveID := extractMoveID(w, r)
	moveData := getMove(moveID, w)
	templ := template.Must(template.New("move.html").ParseFiles("./templates/move.html"))
	templ.Execute(w, moveData)
}

func extractMoveID(w http.ResponseWriter, r *http.Request) string {
	moveID := r.URL.Path[len("/move/"):]
	if moveID == "" {
		http.Error(w, "Move id is missing", http.StatusBadRequest)
		return ""
	}
	return moveID
}

func getMove(moveID string, w http.ResponseWriter) structs.Move {
	move, err := pokeapi.Move(moveID)
	if err != nil {
		http.Error(w, "Pokemon not found", http.StatusNotFound)
		return structs.Move{}
	}

	return move
}
