package array

// Finds minimum of an array
//
//lint:file-ignore U1000 Ignore unused code since this is a DSA repository
func findMinimum(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	var currentMinimum int = arr[0]

	for i := 0; i < len(arr); i++ {
		if currentMinimum > arr[i] {
			currentMinimum = arr[i]
		}
	}

	return currentMinimum
}

// Finds maximum of an array
//
//lint:file-ignore U1000 Ignore unused code since this is a DSA repository
func findMaximum(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	var currentMaximum int = arr[0]

	for i := 0; i < len(arr); i++ {
		if currentMaximum < arr[i] {
			currentMaximum = arr[i]
		}
	}

	return currentMaximum
}
