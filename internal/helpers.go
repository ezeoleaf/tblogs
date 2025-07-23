package internal

import (
	"crypto/md5"
	"encoding/hex"
)

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

// GetHash returns the hash of a slice of strings
func GetHash(textToHash []string) string {
	toHash := ""

	for _, text := range textToHash {
		toHash += text
	}

	hashed := md5.Sum([]byte(toHash))
	hash := hex.EncodeToString(hashed[:])

	return hash
}
