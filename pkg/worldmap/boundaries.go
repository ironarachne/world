package worldmap

import (
	"github.com/ironarachne/world/pkg/grid"
)

// Boundary is a political boundary or other area
type Boundary struct {
	Label    string
	Style    string
	Pattern  string
	Tiles    []Tile
	Edges    []grid.Edge
	Vertices []grid.Coordinate
}

// CalculateBoundaryEdges calculates the edges of a boundary given a tile pixel size and coordinates
func (worldMap WorldMap) CalculateBoundaryEdges(boundary Boundary) []grid.Edge {
	var edges []grid.Edge
	allEdges := []grid.Edge{}

	for _, t := range boundary.Tiles {
		edges = t.CalculateEdges(worldMap.TileHeight, worldMap.TileWidth)
		allEdges = append(allEdges, edges...)
	}

	boundaryEdges := grid.GetUniqueEdges(allEdges)

	return boundaryEdges
}

// CalculateBoundaryVertices calculates the vertices of a boundary given a tile pixel size and coordinates
func (worldMap WorldMap) CalculateBoundaryVertices(boundary Boundary) []grid.Coordinate {
	allVertices := []grid.Coordinate{}

	for _, e := range boundary.Edges {
		allVertices = append(allVertices, e.A)
		allVertices = append(allVertices, e.B)
	}

	return allVertices
}

// CreateBoundary creates a boundary given a set of tiles
func CreateBoundary(label string, style string, pattern string, tiles []Tile) Boundary {
	boundary := Boundary{}

	boundary.Label = label
	boundary.Style = style
	boundary.Pattern = pattern
	boundary.Tiles = tiles

	return boundary
}
