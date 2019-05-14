package worldmap

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/climate"
	"github.com/ironarachne/world/pkg/grid"
)

// Tile is a map tile
type Tile struct {
	Coordinate  grid.Coordinate
	Temperature int
	Humidity    int
	IsInhabited bool
	IsOcean     bool
	TileType    string
}

type tileGenerator struct {
	state    string
	tileType string
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
				TileType:    "ocean",
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

func (worldMap WorldMap) generateLandRandom() [][]Tile {
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

func generateInitialActiveTiles(number int, width int, height int) []grid.Coordinate {
	var x int
	var y int
	var coord grid.Coordinate

	coords := []grid.Coordinate{}

	for i := 0; i < number; i++ {
		x = rand.Intn(width)
		y = rand.Intn(height)
		coord = grid.Coordinate{X: x, Y: y}
		coords = append(coords, coord)
	}

	return coords
}

func (worldMap WorldMap) generateLandSIR() [][]Tile {
	var modelRow []tileGenerator
	var model tileGenerator
	var models [][]tileGenerator

	sourceTiles := worldMap.Tiles

	for _, row := range sourceTiles {
		modelRow = []tileGenerator{}
		for i := 0; i < len(row); i++ {
			model = tileGenerator{
				state:    "susceptible",
				tileType: "ocean",
			}
			modelRow = append(modelRow, model)
		}
		models = append(models, modelRow)
	}

	activeTiles := generateInitialActiveTiles(3, worldMap.Width, worldMap.Height)

	for _, a := range activeTiles {
		models[a.Y][a.X].state = "active"
		models[a.Y][a.X].tileType = "land"
	}

	countActive := 1

	for countActive > 0 {
		for _, a := range activeTiles {
			if a.X > 0 {
				models[a.Y][a.X-1] = evaluateTile(models[a.Y][a.X-1])
			}
			if a.X < worldMap.Width-1 {
				models[a.Y][a.X+1] = evaluateTile(models[a.Y][a.X+1])
			}
			if a.Y > 0 {
				models[a.Y-1][a.X] = evaluateTile(models[a.Y-1][a.X])
			}
			if a.Y < worldMap.Height-1 {
				models[a.Y+1][a.X] = evaluateTile(models[a.Y+1][a.X])
			}
			models[a.Y][a.X].state = "removed"
		}
		activeTiles = []grid.Coordinate{}
		countActive = 0
		for y, row := range models {
			for x, tile := range row {
				if tile.state == "active" {
					activeTiles = append(activeTiles, grid.Coordinate{X: x, Y: y})
					countActive++
				}
			}
		}
	}

	for y, row := range models {
		for x, model := range row {
			if model.tileType == "land" {
				sourceTiles[y][x].IsOcean = false
				sourceTiles[y][x].TileType = "land"
			}
		}
	}

	return sourceTiles
}

func evaluateTile(tile tileGenerator) tileGenerator {
	chance := rand.Intn(100)
	if tile.state == "susceptible" {
		if chance > 50 {
			tile.state = "active"
			tile.tileType = "land"
		}
	}

	return tile
}

func (worldMap WorldMap) findArtifactTilesOfType(tileType string) []grid.Coordinate {
	filteredCoords := []grid.Coordinate{}
	coords := findTilesOfType(tileType, worldMap.Tiles)

	for _, c := range coords {
		if !worldMap.isTileContiguous(6, worldMap.Tiles[c.Y][c.X]) {
			filteredCoords = append(filteredCoords, c)
		}
	}

	return filteredCoords
}

func findTilesOfType(tileType string, tiles [][]Tile) []grid.Coordinate {
	coords := []grid.Coordinate{}

	for y, row := range tiles {
		for x, tile := range row {
			if tile.TileType == tileType {
				coords = append(coords, grid.Coordinate{X: x, Y: y})
			}
		}
	}

	return coords
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

func (worldMap WorldMap) isTileContiguous(level int, tile Tile) bool {
	var adjacentTiles []Tile
	var contiguousTiles []Tile

	contiguousDistance := 0
	checkNext := true

	workingTile := tile

	for i := 1; i <= level; i++ {
		adjacentTiles = worldMap.getAdjacentTiles(workingTile)
		if checkNext {
			checkNext = false
			for _, a := range adjacentTiles {
				if a.TileType == tile.TileType {
					if !isTileInSlice(a, contiguousTiles) {
						contiguousTiles = append(contiguousTiles, a)
					}
					workingTile = a
					checkNext = true
				}
			}
		}
	}

	contiguousDistance = len(contiguousTiles)
	if contiguousDistance >= level {
		return true
	}

	return false
}

func isTileInSlice(tile Tile, tiles []Tile) bool {
	for _, t := range tiles {
		if tile.Coordinate.X == t.Coordinate.X && tile.Coordinate.Y == t.Coordinate.Y {
			return true
		}
	}
	return false
}

func (worldMap WorldMap) removeArtifactOceanTiles() [][]Tile {
	tiles := worldMap.Tiles
	coords := worldMap.findArtifactTilesOfType("ocean")

	for _, c := range coords {
		tiles[c.Y][c.X].TileType = "land"
		tiles[c.Y][c.X].IsOcean = false
	}

	return tiles
}

func (worldMap WorldMap) setTileTemperatures() [][]Tile {
	arcticRange := int(float64(worldMap.Height) * 0.1)
	equatorialRange := int(float64(worldMap.Height) * 0.3)

	tiles := [][]Tile{}
	newRow := []Tile{}

	for y, row := range worldMap.Tiles {
		newRow = []Tile{}
		for _, tile := range row {
			if y <= arcticRange || y >= worldMap.Height-arcticRange {
				tile.Temperature -= rand.Intn(4)
			}
			if y >= equatorialRange && y <= worldMap.Height-equatorialRange {
				tile.Temperature += rand.Intn(4)
			}
			newRow = append(newRow, tile)
		}
		tiles = append(tiles, newRow)
	}

	return tiles
}

func (worldMap WorldMap) setTileHumidities() [][]Tile {
	newRow := []Tile{}
	tiles := [][]Tile{}
	adjacentTiles := []Tile{}
	averageSurroundingHumidity := 0
	isCoastal := false
	totalAdjacentTiles := 0
	totalHumidity := 0
	upOrDown := 0

	for _, row := range worldMap.Tiles {
		newRow = []Tile{}
		for _, tile := range row {
			if !tile.IsOcean {
				adjacentTiles = worldMap.getAdjacentTiles(tile)
				totalAdjacentTiles = len(adjacentTiles)
				totalHumidity = 0
				isCoastal = false
				for _, a := range adjacentTiles {
					if a.TileType == "ocean" {
						isCoastal = true
					} else {
						totalHumidity += a.Humidity
					}
				}
				if isCoastal {
					tile.Humidity += rand.Intn(6)
				} else {
					averageSurroundingHumidity = int(totalHumidity / totalAdjacentTiles)
					upOrDown = rand.Intn(50)
					if upOrDown > 25 {
						tile.Humidity = averageSurroundingHumidity + 1
					} else {
						tile.Humidity = averageSurroundingHumidity - 1
					}
				}
			}
			newRow = append(newRow, tile)
		}
		tiles = append(tiles, newRow)
	}

	return tiles
}

func (worldMap WorldMap) setTileTypes() [][]Tile {
	tiles := [][]Tile{}
	newRow := []Tile{}

	for _, row := range worldMap.Tiles {
		newRow = []Tile{}
		for _, tile := range row {
			if !tile.IsOcean {
				tile.TileType = climate.GetClimateNameForConditions(tile.Humidity, tile.Temperature)
			}
			newRow = append(newRow, tile)
		}
		tiles = append(tiles, newRow)
	}

	return tiles
}
