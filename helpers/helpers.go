package helpers

import "github.com/ezeoleaf/tblogs/models"

// IsIn checks if an int value is inside an int slice
// Returns true if it's in the slice and the index of the value
// If not found, it returns false and -1
func IsIn(v int, s []int) (bool, int) {
	for ix, value := range s {
		if v == value {
			return true, ix
		}
	}

	return false, -1
}

// IsHash checks if a Post hash is inside a slice of Posts. This is used for check if you have a saved post
func IsHash(v string, s []models.Post) (bool, int) {
	for ix, value := range s {
		if v == value.Hash {
			return true, ix
		}
	}

	return false, -1
}
