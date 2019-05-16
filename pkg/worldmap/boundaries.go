package worldmap

import (
	"github.com/ironarachne/world/pkg/grid"
)

// Boundary is a political boundary or other area
type Boundary struct {
	Label           string
	TileCoordinates []grid.Coordinate
}

// CreateBoundary creates a boundary given a set of coordinates
func CreateBoundary(label string, coords []grid.Coordinate) Boundary {
	boundary := Boundary{}

	boundary.Label = label
	boundary.TileCoordinates = coords

	return boundary
}
