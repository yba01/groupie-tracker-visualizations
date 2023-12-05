package main

import (
	"encoding/json"
	"groupie/Tools/models"
	"net/http"
	"text/template"
)

func init() {
	models.Tm, models.Erro = template.ParseGlob("templates/*")
//first
	jsonfiles, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")

	if err != nil {
		models.Theartists = nil
		return
	}

	defer jsonfiles.Body.Close()

	json.NewDecoder(jsonfiles.Body).Decode(&models.Theartists)

//second	
	jsonfiles1, err1 := http.Get("https://groupietrackers.herokuapp.com/api/locations")

	if err1 != nil {
		models.Thelocations.Index = nil
		return
	}

	defer jsonfiles1.Body.Close()

	json.NewDecoder(jsonfiles1.Body).Decode(&models.Thelocations)

	

}
