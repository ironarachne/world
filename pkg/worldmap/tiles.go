package worldmap

import (
	"github.com/ironarachne/world/pkg/geography"
	"math/rand"

	"github.com/ironarachne/world/pkg/grid"
)

// Tile is a map tile
type Tile struct {
	Coordinate  grid.Coordinate
	Points      []grid.Coordinate
	Edges       []grid.Edge
	Temperature int
	Humidity    int
	Elevation   int
	IsInhabited bool
	IsOcean     bool
	TileType    string
}

type tileGenerator struct {
	state    string
	tileType string
}

// CalculateTilePoints returns the pixel positions for the vertices of the tile
func (worldMap WorldMap) CalculateTilePoints(tile Tile) []grid.Coordinate {
	coords := []grid.Coordinate{}
	tileHeight := worldMap.TileHeight
	tileWidth := worldMap.TileWidth

	point := grid.Coordinate{X: tileWidth * tile.Coordinate.X, Y: tileHeight * tile.Coordinate.Y}
	coords = append(coords, point)
	point = grid.Coordinate{X: (tileWidth * tile.Coordinate.X) + tileWidth, Y: tileHeight * tile.Coordinate.Y}
	coords = append(coords, point)
	point = grid.Coordinate{X: (tileWidth * tile.Coordinate.X) + tileWidth, Y: (tileHeight * tile.Coordinate.Y) + tileHeight}
	coords = append(coords, point)
	point = grid.Coordinate{X: tileWidth * tile.Coordinate.X, Y: (tileHeight * tile.Coordinate.Y) + tileHeight}
	coords = append(coords, point)

	return coords
}

// CalculateEdges returns the pixel edges of a tile
func (tile Tile) CalculateEdges(tileHeight int, tileWidth int) []grid.Edge {
	edges := []grid.Edge{}

	edge1 := grid.Edge{A: tile.Points[0], B: tile.Points[1]}
	edges = append(edges, edge1)
	edge2 := grid.Edge{A: tile.Points[1], B: tile.Points[2]}
	edges = append(edges, edge2)
	edge3 := grid.Edge{A: tile.Points[2], B: tile.Points[3]}
	edges = append(edges, edge3)
	edge4 := grid.Edge{A: tile.Points[3], B: tile.Points[0]}
	edges = append(edges, edge4)

	return edges
}

// AllCoordinates gets a slice of all coordinates in the map
func (worldMap WorldMap) AllCoordinates() []grid.Coordinate {
	coords := []grid.Coordinate{}

	for _, row := range worldMap.Tiles {
		for _, t := range row {
			coords = append(coords, t.Coordinate)
		}
	}

	return coords
}

// AllLandCoordinates gets a slice of all coordinates that are not ocean
func (worldMap WorldMap) AllLandCoordinates() []grid.Coordinate {
	tiles := worldMap.Tiles
	coords := []grid.Coordinate{}

	for _, row := range tiles {
		for _, t := range row {
			if !t.IsOcean {
				coords = append(coords, t.Coordinate)
			}
		}
	}

	return coords
}

func (worldMap WorldMap) initializeTiles() [][]Tile {
	var tiles [][]Tile
	var c grid.Coordinate
	var row []Tile
	var tile Tile
	var points []grid.Coordinate
	var edges []grid.Edge

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
				Elevation:   0,
				IsOcean:     true,
				IsInhabited: false,
				TileType:    "ocean",
			}
			points = worldMap.CalculateTilePoints(tile)
			tile.Points = points
			edges = tile.CalculateEdges(worldMap.TileHeight, worldMap.TileWidth)
			tile.Edges = edges
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
			adjacentTiles = worldMap.GetAdjacentTiles(tile)
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
	coords := FindTilesOfType(tileType, worldMap.Tiles)

	for _, c := range coords {
		if !worldMap.isTileContiguous(6, worldMap.Tiles[c.Y][c.X]) {
			filteredCoords = append(filteredCoords, c)
		}
	}

	return filteredCoords
}

// FindTilesByCoordinates finds tiles that match the given slice of coordinates
func FindTilesByCoordinates(coords []grid.Coordinate, tiles [][]Tile) []Tile {
	var tile Tile

	filteredTiles := []Tile{}

	for _, c := range coords {
		tile = tiles[c.Y][c.X]
		filteredTiles = append(filteredTiles, tile)
	}

	return filteredTiles
}

// FindTilesOfType finds tiles that match the tile type
func FindTilesOfType(tileType string, tiles [][]Tile) []grid.Coordinate {
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

// GetAdjacentTiles returns a slice of tiles adjacent to the given tile
func (worldMap WorldMap) GetAdjacentTiles(tile Tile) []Tile {
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
		adjacentTiles = worldMap.GetAdjacentTiles(workingTile)
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

	variance := 0

	for y, row := range worldMap.Tiles {
		newRow = []Tile{}
		for _, tile := range row {
			variance = rand.Intn(3)
			if y <= arcticRange+variance || y >= worldMap.Height-arcticRange+variance {
				tile.Temperature -= rand.Intn(4)
			}
			if y >= equatorialRange+variance && y <= worldMap.Height-equatorialRange+variance {
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
				adjacentTiles = worldMap.GetAdjacentTiles(tile)
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
	var tileGeography geography.Area
	var equatorDistance int
	var err error
	tiles := [][]Tile{}
	newRow := []Tile{}

	for _, row := range worldMap.Tiles {
		newRow = []Tile{}
		for _, tile := range row {
			if !tile.IsOcean {
				equatorDistance = (worldMap.Height / 2) - tile.Coordinate.Y
				tileGeography, err = geography.GenerateSpecific(tile.Temperature, tile.Humidity, tile.Elevation, equatorDistance)
				if err != nil {
					panic(err) // TODO: Properly handle this error
				}
				tile.TileType = tileGeography.Biome.Name
			}
			newRow = append(newRow, tile)
		}
		tiles = append(tiles, newRow)
	}

	return tiles
}
