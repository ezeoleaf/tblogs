package helpers

import "github.com/ezeoleaf/tblogs/models"

func IsIn(v int, s []int) (bool, int) {
	for ix, value := range s {
		if v == value {
			return true, ix
		}
	}

	return false, -1
}

func IsHash(v string, s []models.Post) (bool, int) {
	for ix, value := range s {
		if v == value.Hash {
			return true, ix
		}
	}

	return false, -1
}
