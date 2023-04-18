package world

import (
	"fmt"
	"github.com/ironarachne/world/pkg/culture"
	"github.com/ironarachne/world/pkg/geography/region"
)

// World is an entire fantasy world
type World struct {
	Name              string
	Tiles             [][]Tile
	Cultures          []culture.Culture
	GeographicRegions []region.Region
}

// Generate procedurally generates a world
func Generate() (World, error) {
	var err error
	world := World{}

	world.Tiles = initializeTiles(1024, 1024)
	world.Tiles, err = generateLand(world.Tiles)
	if err != nil {
		return World{}, fmt.Errorf("failed to generate world: %w", err)
	}

	return world, nil
}
