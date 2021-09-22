package geometry

import (
	"math"
)

// Triangle is a triangle
type Triangle struct {
	A  Point
	B  Point
	C  Point
	AB Edge
	BC Edge
	CA Edge
}

// TriangleMesh is a mesh of triangles
type TriangleMesh struct {
	Triangles []Triangle
	Nodes     []Point // the centroids of each triangle
	Vertices  []Point
	Edges     []Edge
}

// BowyerWatson returns a Delaunay triangulation of the given points using the Bowyer-Watson algorithm.
func BowyerWatson(points []Point) TriangleMesh {
	var circumcircle Circle
	var mesh TriangleMesh
	var badTriangles []Triangle
	var finalTriangles []Triangle
	var otherTriangles []Triangle
	var newTrianglePoints []Point
	var newTriangle Triangle
	var polygon Polygon
	var vertices []Point

	superTriangle := SuperTriangle(points)

	mesh.Triangles = append(mesh.Triangles, superTriangle)

	for _, p := range points {
		mesh.Vertices = append(mesh.Vertices, p)
		badTriangles = []Triangle{}
		for _, t := range mesh.Triangles {
			circumcircle = Circumcircle(t)
			if p.InCircle(circumcircle) {
				badTriangles = append(badTriangles, t)
			}
		}

		polygon = Polygon{}

		for _, t := range badTriangles {
			for _, e := range t.Edges() {
				otherTriangles = t.RemoveFrom(badTriangles)
				if !e.IsSharedByTriangles(otherTriangles) {
					polygon.Edges = append(polygon.Edges, e)
				}
			}
		}

		for _, t := range badTriangles {
			mesh.Triangles = t.RemoveFrom(mesh.Triangles)
		}

		for _, e := range polygon.Edges {
			newTrianglePoints = []Point{
				{
					X: e.A.X,
					Y: e.A.Y,
				},
				{
					X: e.B.X,
					Y: e.B.Y,
				},
				{
					X: p.X,
					Y: p.Y,
				},
			}
			newTriangle = TriangleFromPoints(newTrianglePoints)
			mesh.Triangles = append(mesh.Triangles, newTriangle)
		}
	}

	// Remove any triangles that have a vertex from the super-triangle from the triangulation
	superVertices := superTriangle.Vertices()
	match := false

	for _, t := range mesh.Triangles {
		vertices = t.Vertices()
		for _, v := range vertices {
			for _, s := range superVertices {
				if v.Equals(s) {
					match = true
				}
			}
		}

		if !match {
			finalTriangles = append(finalTriangles, t)
		} else {
			match = false
		}
	}

	mesh.Triangles = finalTriangles

	for _, t := range mesh.Triangles {
		mesh.Edges = append(mesh.Edges, t.Edges()...)
	}

	mesh.Vertices = GetUniquePoints(mesh.Vertices)
	mesh.Edges = GetUniqueEdges(mesh.Edges)

	for _, t := range mesh.Triangles {
		mesh.Nodes = append(mesh.Nodes, Centroid(t.Vertices()))
	}

	return mesh
}

// Circumcenter finds the circumcenter for a given triangle
func Circumcenter(triangle Triangle) Point {
	a := triangle.A
	b := triangle.B
	c := triangle.C

	// We use only two of the sides, since the bisector of the third will intersect at the same point

	// Calculate midpoints of each side (AB, BC)
	mab := Point{
		X: (a.X + b.X) / 2,
		Y: (a.Y + b.Y) / 2,
	}

	mbc := Point{
		X: (b.X + c.X) / 2,
		Y: (b.Y + c.Y) / 2,
	}

	// Calculate slopes of each side (AB, BC)
	sab := (b.Y - a.Y) / (b.X - a.X)
	sbc := (c.Y - b.Y) / (c.X - b.X)

	// Calculate the slope of the perpendicular bisector of each side (AB, BC)
	bab := -1 / sab
	bbc := -1 / sbc

	// Calculate the y-intercept of the perpendicular bisector of each side
	yab := -((bab * mab.X) - mab.Y)
	ybc := -((bbc * mbc.X) - mbc.Y)

	// Calculate the point where the lines intersect
	x := (ybc - yab) / (bab - bbc)
	y := (bab * x) + yab

	center := Point{
		X: x,
		Y: y,
	}

	return center
}

// Circumcircle calculates the circumcircle for a given triangle
func Circumcircle(triangle Triangle) Circle {
	a := triangle.AB.Length()
	b := triangle.BC.Length()
	c := triangle.CA.Length()

	radius := (a * b * c) / math.Sqrt((a+b+c)*(b+c-a)*(c+a-b)*(a+b-c))
	center := Circumcenter(triangle)

	circle := Circle{
		Center: center,
		Radius: radius,
	}

	return circle
}

// Edges returns all the edges of a triangle
func (t Triangle) Edges() []Edge {
	edges := []Edge{
		t.AB,
		t.BC,
		t.CA,
	}

	return edges
}

// Equals checks to see if a triangle equals another triangle
func (t Triangle) Equals(other Triangle) bool {
	if t.A.Equals(other.A) && t.B.Equals(other.B) && t.C.Equals(other.C) {
		return true
	}

	return false
}

// RemoveFrom removes a triangle from a slice of triangles
func (t Triangle) RemoveFrom(triangles []Triangle) []Triangle {
	var newTriangles []Triangle

	for _, o := range triangles {
		if !t.Equals(o) {
			newTriangles = append(newTriangles, o)
		}
	}

	return newTriangles
}

// IsEdgeSharedByTriangles checks to see if the given edge is shared by any of the triangles in the given slice
func (e Edge) IsSharedByTriangles(triangles []Triangle) bool {
	var edges []Edge

	for _, t := range triangles {
		edges = append(edges, t.Edges()...)
	}

	for _, other := range edges {
		if e.Equals(other) {
			return true
		}
	}

	return false
}

// PointInTriangle determines if the given point is in the given triangle
func PointInTriangle(pt Point, t Triangle) bool {
	v1 := t.A
	v2 := t.B
	v3 := t.C
	var d1, d2, d3 float64
	var hasNeg, hasPos bool

	d1 = sign(pt, v1, v2)
	d2 = sign(pt, v2, v3)
	d3 = sign(pt, v3, v1)

	hasNeg = (d1 < 0) || (d2 < 0) || (d3 < 0)
	hasPos = (d1 > 0) || (d2 > 0) || (d3 > 0)

	return !(hasNeg && hasPos)
}

// SuperTriangle returns a triangle big enough to contain all of the given points
func SuperTriangle(points []Point) Triangle {
	minY := 0.0
	maxY := 0.0
	minX := 0.0
	maxX := 0.0

	for _, p := range points {
		if p.X < minX {
			minX = p.X
		}
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y < minY {
			minY = p.Y
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}

	leftBound := Point{
		X: minX - 5000,
		Y: minY - 5000,
	}

	rightBound := Point{
		X: maxX * 50,
		Y: maxY * 50,
	}

	corner := Point{
		X: leftBound.X,
		Y: 0,
	}

	trianglePoints := []Point{
		leftBound,
		rightBound,
		corner,
	}

	triangle := TriangleFromPoints(trianglePoints)

	return triangle
}

// ToPolygon turns a triangle into a polygon
func (t Triangle) ToPolygon() Polygon {
	var edges []Edge
	var points []Point

	edges = append(edges, t.AB)
	edges = append(edges, t.BC)
	edges = append(edges, t.CA)

	points = append(points, t.A)
	points = append(points, t.B)
	points = append(points, t.C)

	polygon := Polygon{
		Edges:  edges,
		Points: points,
	}

	return polygon
}

// Vertices returns a slice of the points of the given triangle
func (t Triangle) Vertices() []Point {
	vertices := []Point{
		t.A,
		t.B,
		t.C,
	}

	return vertices
}

// TriangleFromPoints returns a triangle based on a set of points
func TriangleFromPoints(points []Point) Triangle {
	polygon := PolygonFromPoints(points)

	triangle := Triangle{
		A:  polygon.Points[0],
		B:  polygon.Points[1],
		C:  polygon.Points[2],
		AB: polygon.Edges[0],
		BC: polygon.Edges[1],
		CA: polygon.Edges[2],
	}

	return triangle
}

func sign(p1 Point, p2 Point, p3 Point) float64 {
	return (p1.X-p3.X)*(p2.Y-p3.Y) - (p2.X-p3.X)*(p1.Y-p3.Y)
}
