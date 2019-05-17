package worldmap

import (
	"github.com/ironarachne/world/pkg/grid"
)

// Boundary is a political boundary or other area
type Boundary struct {
	Label           string
	Style           string
	Pattern         string
	TileCoordinates []grid.Coordinate
}

// CreateBoundary creates a boundary given a set of coordinates
func CreateBoundary(label string, style string, pattern string, coords []grid.Coordinate) Boundary {
	boundary := Boundary{}

	boundary.Label = label
	boundary.Style = style
	boundary.Pattern = pattern
	boundary.TileCoordinates = coords

	return boundary
}
