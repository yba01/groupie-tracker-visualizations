package Handlers

import (
	"encoding/json"
	"groupie/Tools/models"
	"groupie/Tools/utils"
	"net/http"
	"strconv"
	"strings"
)

func Home(rw http.ResponseWriter, r *http.Request) {

	if models.Erro != nil {
		utils.Error500(models.Tm, rw)
		return
	}

	if utils.ErrorUrl(r, "/") {
		utils.Error404(models.Tm, rw)
		return
	}

	if models.Thelocations.Index == nil {
		utils.Error500(models.Tm, rw)
		return
	}

	var THElocations []string

	for _, char := range models.Thelocations.Index {
		for _, char1 := range char.Locations {
			if !utils.DuplicatesV2(THElocations, char1) {
					THElocations = append(THElocations, char1)
			}
		}
	}

	Allstructs := models.Complete2{Art_ists: models.Theartists, Loca_tions: THElocations}

	models.Tm.ExecuteTemplate(rw, "index.html", Allstructs)
}

func Artistinfos(rw http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	ID, erreu := strconv.Atoi(id)

	if models.Erro != nil {
		utils.Error500(models.Tm, rw)
		return
	}
	if erreu != nil {
		utils.Error400(models.Tm, rw)
		return
	}

	if ID < 1 || ID > 52 {
		utils.Error400(models.Tm, rw)
		return
	}

	if utils.ErrorUrl(r, "/artistsinfos") {
		utils.Error404(models.Tm, rw)
		return
	}

	jsonfiles, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + id)

	if err != nil {
		utils.Error500(models.Tm, rw)
		return
	}

	defer jsonfiles.Body.Close()

	var artistsinf models.Artists

	err = json.NewDecoder(jsonfiles.Body).Decode(&artistsinf)

	if err != nil {
		utils.Error500(models.Tm, rw)
		return
	}

	jsonfiles1, err1 := http.Get("https://groupietrackers.herokuapp.com/api/relation/" + id)

	if err1 != nil {
		utils.Error500(models.Tm, rw)
		return
	}

	defer jsonfiles1.Body.Close()

	var relations models.Relations

	err1 = json.NewDecoder(jsonfiles1.Body).Decode(&relations)

	if err1 != nil {
		utils.Error500(models.Tm, rw)
		return
	}
	if len(r.URL.String()) > 19 {
		utils.Error400(models.Tm, rw)
		return
	}

	if strings.HasSuffix(r.URL.String(), "&") {
		utils.Error400(models.Tm, rw)
		return
	}

	Allstruct := models.Complete{Art: artistsinf, Rel: relations}

	models.Tm.ExecuteTemplate(rw, "artistsinfo.html", Allstruct)
}
