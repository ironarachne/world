package geometry

import (
	"context"
	"math"

	"github.com/ironarachne/world/pkg/random"
)

// Point is an x, y location
type Point struct {
	X float64
	Y float64
}

// Centroid returns a point that is in the center of a polygon defined by the given points
func Centroid(points []Point) Point {
	x := 0.0
	y := 0.0

	for _, p := range points {
		x += p.X
		y += p.Y
	}

	point := Point{
		X: x / float64(len(points)),
		Y: y / float64(len(points)),
	}

	return point
}

// Distance returns the distance between two points
func Distance(a Point, b Point) float64 {
	distance := math.Sqrt(math.Pow(b.X-a.X, 2) + math.Pow(b.Y-a.Y, 2))

	return distance
}

// Equals tests the equality of two points
func (a Point) Equals(b Point) bool {
	if a.X == b.X && a.Y == b.Y {
		return true
	}
	return false
}

// GetUniquePoints returns only unique points from a slice of points
func GetUniquePoints(points []Point) []Point {
	var unique []Point

	for _, c := range points {
		if !InSlice(c, unique) {
			unique = append(unique, c)
		}
	}

	return unique
}

// InSlice tests to see if a point is in a slice of points
func InSlice(c Point, target []Point) bool {
	for _, t := range target {
		if c.Equals(t) {
			return true
		}
	}
	return false
}

// RemovePoints returns a slice of points with a set removed
func RemovePoints(pointsToRemove []Point, target []Point) []Point {
	var coords []Point

	for _, c := range target {
		if !InSlice(c, pointsToRemove) {
			coords = append(coords, c)
		}
	}

	return coords
}

// SortPointsByDistance sorts points by distance
func SortPointsByDistance(source []Point) []Point {
	var sorted []Point
	var remaining []Point

	nearestIndex := 0

	remaining = make([]Point, len(source))
	copy(remaining, source)

	sorted = append(sorted, remaining[0])
	remaining = remaining[1:]

	last := sorted[0]

	for len(remaining) > 0 {
		nearestIndex = NearestPointIndex(last, remaining)
		last = remaining[nearestIndex]
		sorted = append(sorted, last)
		remaining[nearestIndex] = remaining[len(remaining)-1]
		remaining[len(remaining)-1] = Point{}
		remaining = remaining[:len(remaining)-1]
	}

	return sorted
}

// NearestPointIndex finds the index in a slice of points of the nearest point to the given one
func NearestPointIndex(p Point, coords []Point) int {
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

// GenerateRandomPoints generates a set of random points inside the given height and width according to the number given
func GenerateRandomPoints(ctx context.Context, height int, width int, number int) []Point {
	var points []Point
	var point Point

	for n := 0; n < number; n++ {
		point = RandomPoint(ctx, height, width)
		points = append(points, point)
	}

	points = GetUniquePoints(points)

	return points
}

// GenerateRandomPointsByDimensions generates a set of random points of the given height, width, and density. Density is a floating point number between 0 and 1.
func GenerateRandomPointsByDimensions(ctx context.Context, height int, width int, density float64) []Point {
	var points []Point
	var point Point

	numberOfPoints := int(float64(height*width) * density)

	for n := 0; n < numberOfPoints; n++ {
		point = RandomPoint(ctx, height, width)
		points = append(points, point)
	}

	points = GetUniquePoints(points)

	return points
}

// RandomPoint returns a random point in the given height and width
func RandomPoint(ctx context.Context, height int, width int) Point {
	x := float64(random.Intn(ctx, width))
	y := float64(random.Intn(ctx, height))

	p := Point{
		X: x,
		Y: y,
	}

	return p
}

// RemoveClumps removes all points from a given set of points that are closer than N distance
func RemoveClumps(points []Point, distance float64) []Point {
	var result []Point

	sorted := SortPointsByDistance(points)

	for i, p := range sorted {
		if i < len(sorted)-1 {
			if Distance(p, sorted[i+1]) > distance {
				result = append(result, p)
			}
		}
	}

	return result
}
