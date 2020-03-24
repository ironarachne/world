package world

import (
	"github.com/ironarachne/world/pkg/culture"
	"github.com/ironarachne/world/pkg/geography/region"
)

// World is an entire fantasy world
type World struct {
	Name string
	Tiles [][]Tile
	Cultures []culture.Culture
	GeographicRegions []region.Region
}

// Generate procedurally generates a world
func Generate() World {
	world := World{}

	world.Tiles = initializeTiles(1280, 1024)
	world.Tiles = generateLand(world.Tiles)

	return world
}
