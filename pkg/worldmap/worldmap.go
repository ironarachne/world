package worldmap

import (
	"github.com/ironarachne/world/pkg/grid"
)

// WorldMap is a world map. It contains tiles.
type WorldMap struct {
	Height int
	Width  int
	Tiles  [][]Tile
}

// Tile is a map tile
type Tile struct {
	Coordinate  grid.Coordinate
	Temperature int
	Humidity    int
	IsInhabited bool
	IsOcean     bool
}

// Generate creates a world map
func Generate(height int, width int) WorldMap {
	worldMap := WorldMap{}

	worldMap.Height = height
	worldMap.Width = width

	tiles := worldMap.initializeTiles()
	worldMap.Tiles = tiles
	worldMap.Tiles = worldMap.generateLandSIR()

	return worldMap
}
