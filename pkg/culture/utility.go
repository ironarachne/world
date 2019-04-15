package culture

// inSlice checks to see if a string is in an array of strings
func inSlice(item string, collection []string) bool {
	for _, element := range collection {
		if item == element {
			return true
		}
	}
	return false
}

// sliceWithin checks to see if a slice of strings is wholly within another slice of strings
func sliceWithin(slice1 []string, slice2 []string) bool {
	if len(slice1) == 0 {
		return true
	}

	for _, i := range slice1 {
		if !inSlice(i, slice2) {
			return false
		}
	}
	return true
}

// slicePartlyWithin checks to see if at least one element of a slice is in the second slice
func slicePartlyWithin(slice1 []string, slice2 []string) bool {
	if len(slice1) == 0 {
		return true
	}

	for _, i := range slice1 {
		if inSlice(i, slice2) {
			return true
		}
	}
	return false
}
