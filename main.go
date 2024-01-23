package main

import (
	"fmt"
	Handlers "groupie/Tools/handlers"
	"net/http"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 0 {
		fmt.Println("Usage: go run .")
		return
	}
	http.HandleFunc("/", Handlers.Home)
	http.HandleFunc("/artistsinfos", Handlers.Artistinfos)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	fmt.Println("Starting at : http://localhost:8081")
	http.ListenAndServe(":8081", nil)
}
