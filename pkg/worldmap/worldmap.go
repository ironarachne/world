package worldmap

import (
	"math/rand"

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
	worldMap.Tiles = worldMap.generateLand()
	for i := 0; i < 2; i++ {
		worldMap.Tiles = worldMap.improveLand()
	}

	return worldMap
}

func (worldMap WorldMap) initializeTiles() [][]Tile {
	var tiles [][]Tile
	var c grid.Coordinate
	var row []Tile
	var tile Tile

	for y := 0; y < worldMap.Height; y++ {
		row = []Tile{}
		for x := 0; x < worldMap.Width; x++ {
			c = grid.Coordinate{
				X: x,
				Y: y,
			}
			tile = Tile{
				Coordinate:  c,
				Temperature: 5,
				Humidity:    5,
				IsOcean:     true,
				IsInhabited: false,
			}
			row = append(row, tile)
		}
		tiles = append(tiles, row)
	}

	return tiles
}

func (worldMap WorldMap) improveLand() [][]Tile {
	var adjacentTiles []Tile
	var countOcean int
	var isLandChance int
	var row []Tile
	var tile Tile
	var tiles [][]Tile

	for _, r := range worldMap.Tiles {
		row = []Tile{}
		for _, t := range r {
			tile = t
			adjacentTiles = worldMap.getAdjacentTiles(tile)
			countOcean = 0
			for _, a := range adjacentTiles {
				if a.IsOcean {
					countOcean++
				}
			}
			isLandChance = 10 - countOcean
			if isLandChance > 6 {
				tile.IsOcean = false
			} else {
				tile.IsOcean = true
			}
			row = append(row, tile)
		}
		tiles = append(tiles, row)
	}

	return tiles
}

func (worldMap WorldMap) generateLand() [][]Tile {
	var tile Tile
	var tiles [][]Tile
	var row []Tile
	var isLandChance int

	sourceTiles := worldMap.Tiles

	for _, r := range sourceTiles {
		row = []Tile{}
		for _, t := range r {
			tile = t
			isLandChance = rand.Intn(10)
			if isLandChance > 6 {
				tile.IsOcean = false
			}
			row = append(row, tile)
		}
		tiles = append(tiles, row)
	}

	return tiles
}

func (worldMap WorldMap) getAdjacentTiles(tile Tile) []Tile {
	var adjacentTiles []Tile

	if tile.Coordinate.X > 0 {
		adjacentTiles = append(adjacentTiles, worldMap.Tiles[tile.Coordinate.Y][tile.Coordinate.X-1])
	}
	if tile.Coordinate.X < worldMap.Width-1 {
		adjacentTiles = append(adjacentTiles, worldMap.Tiles[tile.Coordinate.Y][tile.Coordinate.X+1])
	}
	if tile.Coordinate.Y > 0 {
		adjacentTiles = append(adjacentTiles, worldMap.Tiles[tile.Coordinate.Y-1][tile.Coordinate.X])
	}
	if tile.Coordinate.Y < worldMap.Height-1 {
		adjacentTiles = append(adjacentTiles, worldMap.Tiles[tile.Coordinate.Y+1][tile.Coordinate.X])
	}

	return adjacentTiles
}
