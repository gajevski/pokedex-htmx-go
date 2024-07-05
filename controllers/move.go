package controllers

import (
	"net/http"
	"text/template"
)

func MovePageController(w http.ResponseWriter, r *http.Request) {
	templ := template.Must(template.New("move.html").ParseFiles("./templates/move.html"))
	templ.Execute(w, r)
}
