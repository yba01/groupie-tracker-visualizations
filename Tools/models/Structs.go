package models

type Artists struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type Relations struct {
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Complete struct {
	Art Artists
	Rel Relations
}

type Locationss struct {
	Index []struct {
		Id        int      `json:"id"`
		Locations []string `json:"locations"`
	} `json:"index"`
}

type Complete1 struct {
	Art_ists   []Artists
	Loca_tions Locationss
}
