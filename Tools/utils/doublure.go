package utils

import "groupie/Tools/models"

func Duplicates(table []models.Artists, data models.Artists) bool {
	for _, t := range table {
		if t.ID == data.ID {
			return true
		}
	}
	return false
}