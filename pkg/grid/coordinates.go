package grid

import (
	"math"
)

// Coordinate is an x, y location
type Coordinate struct {
	X int
	Y int
}

// Distance returns the distance between two points
func Distance(a Coordinate, b Coordinate) float64 {
	distance := math.Sqrt(math.Pow(float64(b.X-a.X), 2) + math.Pow(float64(b.Y-a.Y), 2))

	return distance
}

// Equals tests the equality of two coordinates
func Equals(a Coordinate, b Coordinate) bool {
	if a.X == b.X && a.Y == b.Y {
		return true
	}
	return false
}

// GetUniqueCoordinates returns only unique coordinates from a slice of coordinates
func GetUniqueCoordinates(coords []Coordinate) []Coordinate {
	unique := []Coordinate{}

	for _, c := range coords {
		if !InSlice(c, unique) {
			unique = append(unique, c)
		}
	}

	return unique
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

// SortCoordinatesByDistance sorts coordinates by distance
func SortCoordinatesByDistance(source []Coordinate) []Coordinate {
	var sorted []Coordinate
	var remaining []Coordinate

	nearestIndex := 0

	remaining = make([]Coordinate, len(source))
	copy(remaining, source)

	sorted = append(sorted, remaining[0])
	remaining = remaining[1:]

	last := sorted[0]

	for len(remaining) > 0 {
		nearestIndex = NearestCoordinateIndex(last, remaining)
		last = remaining[nearestIndex]
		sorted = append(sorted, last)
		remaining[nearestIndex] = remaining[len(remaining)-1]
		remaining[len(remaining)-1] = Coordinate{}
		remaining = remaining[:len(remaining)-1]
	}

	return sorted
}

// NearestCoordinateIndex finds the index in a slice of coordinates of the nearest coordinate to the given one
func NearestCoordinateIndex(p Coordinate, coords []Coordinate) int {
	var distance float64

	smallestDistance := Distance(p, coords[0])
	smallest := 0

	for i, c := range coords {
		distance = Distance(p, c)
		if distance < smallestDistance {
			smallest = i
			smallestDistance = distance
		}
	}

	return smallest
}
