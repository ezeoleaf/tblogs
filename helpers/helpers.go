package helpers

func IsIn(v int, s []int) (bool, int) {
	for ix, value := range s {
		if v == value {
			return true, ix
		}
	}

	return false, -1
}
