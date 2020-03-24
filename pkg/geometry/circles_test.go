package geometry

import "testing"

func TestPoint_InCircle(t *testing.T) {
	point := Point{
		X: 20,
		Y: 20,
	}

	circle1 := Circle{
		Center: Point{
			X: 10,
			Y: 10,
		},
		Radius: 50,
	}

	circle2 := Circle{
		Center: Point{
			X: 30,
			Y: 30,
		},
		Radius: 2,
	}

	if !point.InCircle(circle1) {
		t.Error("failed to detect point in encompassing circle")
	}

	if point.InCircle(circle2) {
		t.Error("failed to detect point not in external circle")
	}
}
