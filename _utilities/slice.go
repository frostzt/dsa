package utilities

// Compares two provided slices value by value returns true if each matches
// otherwise returns false!
func CompareSliceValueToValue(a, b []int) bool {
	lena := len(a)
	lenb := len(b)

	if lena != lenb {
		return false
	}

	for i := 0; i < lena; i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
