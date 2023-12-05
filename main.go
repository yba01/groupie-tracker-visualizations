package main

import (
	Handlers "groupie/Tools/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", Handlers.Home)
	http.HandleFunc("/artistsinfos", Handlers.Artistinfos)
	http.HandleFunc("/search", Handlers.Search)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8081", nil)
}
