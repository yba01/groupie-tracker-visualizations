package main

import (
	"fmt"
	Handlers "groupie/Tools/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", Handlers.Home)
	http.HandleFunc("/artistsinfos", Handlers.Artistinfos)
	http.HandleFunc("/search", Handlers.Search)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	fmt.Println("Starting at : http://localhost:8081")
	http.ListenAndServe(":8081", nil)
}
