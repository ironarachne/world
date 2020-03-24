package geometry

import "testing"

func TestCentroid(t *testing.T) {
	points := []Point{
		{
			X: 0,
			Y: 0,
		},
		{
			X: 10,
			Y: 0,
		},
		{
			X: 10,
			Y: 10,
		},
		{
			X: 0,
			Y: 10,
		},
	}

	centroid := Centroid(points)

	if centroid.X != 5 || centroid.Y != 5 {
		t.Error("failed to get the centroid of a polygon")
	}
}

func TestDistance(t *testing.T) {
	a := Point{
		X: 0,
		Y: 1,
	}

	b := Point{
		X: 0,
		Y: 2,
	}

	c := Point{
		X: 1,
		Y: 4,
	}

	d := Point{
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
	a := Point{
		X: 0,
		Y: 1,
	}

	b := Point{
		X: 0,
		Y: 1,
	}

	c := Point{
		X: 1,
		Y: 1,
	}

	if !a.Equals(b) {
		t.Error("failed to see that (0,1) and (0,1) are equal")
	}

	if !b.Equals(a) {
		t.Error("failed to see that (0,1) and (0,1) are equal")
	}

	if a.Equals(c) {
		t.Error("failed to see that (0,1) and (1,1) are different")
	}
}

func TestGetUniqueCoordinates(t *testing.T) {
	coordinates := []Point{
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

	unique := GetUniquePoints(coordinates)

	if len(unique) != expectedCount {
		t.Error("failed to get unique coordinates")
	}
}

func TestInSlice(t *testing.T) {
	a := Point{
		X: 0,
		Y: 1,
	}

	b := Point{
		X: 13,
		Y: 1,
	}

	coords := []Point{
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
	remove := []Point{
		{
			X: 1,
			Y: 1,
		},
	}

	origin := []Point{
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

	filtered := RemovePoints(remove, origin)

	if len(filtered) != 1 {
		t.Error("failed to remove correct number of matching coordinates from slice")
	}
}

func TestSortCoordinatesByDistance(t *testing.T) {
	unsorted := []Point{
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

	sorted := SortPointsByDistance(unsorted)

	if !sorted[0].Equals(unsorted[0]) {
		t.Error("failed to put origin coordinates first in sort")
	}

	if !sorted[1].Equals(unsorted[2]) {
		t.Error("failed to put next nearest coordinates second in sort")
	}

	if !sorted[2].Equals(unsorted[1]) {
		t.Error("failed to put furthest coordinates last in sort")
	}
}

func TestNearestCoordinateIndex(t *testing.T) {
	a := Point{
		X: 0,
		Y: 0,
	}

	coords := []Point{
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

	nearest := NearestPointIndex(a, coords)

	if nearest != 1 {
		t.Error("failed to find nearest coordinate in slice")
	}
}
