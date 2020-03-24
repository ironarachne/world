package geometry

import (
	"testing"
)

func TestCircumcenter(t *testing.T) {
	trianglePoints := []Point{
		{
			X: 0,
			Y: 0,
		},
		{
			X: 10,
			Y: 10,
		},
		{
			X: 20,
			Y: 0,
		},
	}
	triangle := TriangleFromPoints(trianglePoints)

	circumcenter := Circumcenter(triangle)

	if circumcenter.X != 10 {
		t.Error("Circumcenter for X was wrong")
	}

	if circumcenter.Y != 0 {
		t.Error("Circumcenter for Y was wrong")
	}
}

func TestCircumcircle(t *testing.T) {
	trianglePoints := []Point{
		{
			X: 0,
			Y: 0,
		},
		{
			X: 10,
			Y: 10,
		},
		{
			X: 20,
			Y: 0,
		},
	}
	triangle := TriangleFromPoints(trianglePoints)

	circumcircle := Circumcircle(triangle)

	if circumcircle.Radius != 10 {
		t.Error("Radius for circumcircle is wrong")
	}

	if circumcircle.Center.X != 10 {
		t.Error("Centerpoint X for circumcircle is wrong")
	}

	if circumcircle.Center.Y != 0 {
		t.Error("Centerpoint Y for circumcircle is wrong")
	}

	trianglePoints = []Point{
		{
			X: -100000,
			Y: -100000,
		},
		{
			X: -100000,
			Y: 100000,
		},
		{
			X: 100000,
			Y: 0,
		},
	}
	triangle = TriangleFromPoints(trianglePoints)

	circumcircle = Circumcircle(triangle)

	if circumcircle.Radius != 125000 {
		t.Error("Radius for giant circumcircle is wrong")
	}
}

func TestSuperTriangle(t *testing.T) {
	points := []Point{
		{
			X: 0,
			Y: 10,
		},
		{
			X: 20,
			Y: 30,
		},
		{
			X: 40,
			Y: 20,
		},
	}

	triangle := SuperTriangle(points)

	for _, p := range points {
		if !PointInTriangle(p, triangle) {
			t.Error("point was not inside supertriangle")
		}
	}
}