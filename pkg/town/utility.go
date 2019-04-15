package town

// inSlice checks to see if a string is in an array of strings
func inSlice(item string, collection []string) bool {
	for _, element := range collection {
		if item == element {
			return true
		}
	}
	return false
}
