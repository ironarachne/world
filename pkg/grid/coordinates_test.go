package grid

import "testing"

func TestDistance(t *testing.T) {
	a := Coordinate{
		X: 0,
		Y: 1,
	}

	b := Coordinate{
		X: 0,
		Y: 2,
	}

	c := Coordinate{
		X: 1,
		Y: 4,
	}

	d := Coordinate{
		X: 12,
		Y: 4,
	}

	distance := Distance(a, b)

	if distance != 1 {
		t.Error("failed to calculate distance between (0,1) and (0,2) correctly")
	}

	distance = Distance(c, d)

	if distance != 11 {
		t.Error("failed to calculate distance between (1,4) and (12,4) correctly")
	}
}

func TestEquals(t *testing.T) {
	a := Coordinate{
		X: 0,
		Y: 1,
	}

	b := Coordinate{
		X: 0,
		Y: 1,
	}

	c := Coordinate{
		X: 1,
		Y: 1,
	}

	if !Equals(a, b) {
		t.Error("failed to see that (0,1) and (0,1) are equal")
	}

	if !Equals(b, a) {
		t.Error("failed to see that (0,1) and (0,1) are equal")
	}

	if Equals(a, c) {
		t.Error("failed to see that (0,1) and (1,1) are different")
	}
}

func TestGetUniqueCoordinates(t *testing.T) {
	coordinates := []Coordinate{
		{
			X: 0,
			Y: 0,
		},
		{
			X: 0,
			Y: 0,
		},
		{
			X: 0,
			Y: 1,
		},
		{
			X: 1,
			Y: 1,
		},
		{
			X: 1,
			Y: 0,
		},
	}

	expectedCount := 4

	unique := GetUniqueCoordinates(coordinates)

	if len(unique) != expectedCount {
		t.Error("failed to get unique coordinates")
	}
}

func TestInSlice(t *testing.T) {
	a := Coordinate{
		X: 0,
		Y: 1,
	}

	b := Coordinate{
		X: 13,
		Y: 1,
	}

	coords := []Coordinate{
		{
			X: 0,
			Y: 1,
		},
		{
			X: 3,
			Y: 2,
		},
	}

	if !InSlice(a, coords) {
		t.Error("failed to find (0,1) in slice of coordinates")
	}

	if InSlice(b, coords) {
		t.Error("falsely found (13,1) in slice of coordinates")
	}
}

func TestRemoveCoordinatesFromSlice(t *testing.T) {
	remove := []Coordinate{
		{
			X: 1,
			Y: 1,
		},
	}

	origin := []Coordinate{
		{
			X: 1,
			Y: 1,
		},
		{
			X: 2,
			Y: 2,
		},
		{
			X: 1,
			Y: 1,
		},
	}

	filtered := RemoveCoordinatesFromSlice(remove, origin)

	if len(filtered) != 1 {
		t.Error("failed to remove correct number of matching coordinates from slice")
	}
}

func TestSortCoordinatesByDistance(t *testing.T) {
	unsorted := []Coordinate{
		{
			X: 0,
			Y: 0,
		},
		{
			X: 10,
			Y: 10,
		},
		{
			X: 0,
			Y: 3,
		},
	}

	sorted := SortCoordinatesByDistance(unsorted)

	if !Equals(sorted[0], unsorted[0]) {
		t.Error("failed to put origin coordinates first in sort")
	}

	if !Equals(sorted[1], unsorted[2]) {
		t.Error("failed to put next nearest coordinates second in sort")
	}

	if !Equals(sorted[2], unsorted[1]) {
		t.Error("failed to put furthest coordinates last in sort")
	}
}

func TestNearestCoordinateIndex(t *testing.T) {
	a := Coordinate{
		X: 0,
		Y: 0,
	}

	coords := []Coordinate{
		{
			X: 20,
			Y: 10,
		},
		{
			X: 1,
			Y: 0,
		},
		{
			X: 15,
			Y: 0,
		},
	}

	nearest := NearestCoordinateIndex(a, coords)

	if nearest != 1 {
		t.Error("failed to find nearest coordinate in slice")
	}
}
