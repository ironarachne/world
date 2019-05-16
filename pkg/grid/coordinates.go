package grid

// Equals tests the equality of two coordinates
func Equals(a Coordinate, b Coordinate) bool {
	if a.X == b.X && a.Y == b.Y {
		return true
	}
	return false
}

// InSlice tests to see if a coordinate is in a slice of coordinates
func InSlice(c Coordinate, target []Coordinate) bool {
	for _, t := range target {
		if Equals(c, t) {
			return true
		}
	}
	return false
}

// RemoveCoordinatesFromSlice returns a slice of coordinates with a set removed
func RemoveCoordinatesFromSlice(coordinatesToRemove []Coordinate, target []Coordinate) []Coordinate {
	coords := []Coordinate{}

	for _, c := range target {
		if !InSlice(c, coordinatesToRemove) {
			coords = append(coords, c)
		}
	}

	return coords
}
