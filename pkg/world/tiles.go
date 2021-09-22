package world

import (
	"context"
	"github.com/ironarachne/world/pkg/geography/region"
	"github.com/ironarachne/world/pkg/geometry"
	"math"

	"github.com/ojrac/opensimplex-go"

	"github.com/ironarachne/world/pkg/random"
)

// Tile is a world tile
type Tile struct {
	Altitude    int
	Temperature int
	Humidity    int
	IsOcean     bool
}

func initializeTiles(width int, height int) [][]Tile {
	tiles := make([][]Tile, height)
	for i := 0; i < height; i++ {
		tiles[i] = make([]Tile, width)
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			tiles[i][j] = Tile{
				IsOcean: true,
			}
		}
	}

	return tiles
}

func newCellGrid(width int, height int) [][]float64 {
	cells := make([][]float64, height)
	for i := 0; i < height; i++ {
		cells[i] = make([]float64, width)
	}

	return cells
}

func getDistanceToNearestOcean(x int, y int, tiles [][]Tile) int {
	var subjects []Tile

	xMax := x
	yMax := y

	if len(tiles)-yMax > yMax {
		yMax = len(tiles) - yMax
	}

	if len(tiles[0])-xMax > xMax {
		xMax = len(tiles[0]) - xMax
	}

	distance := xMax
	if yMax > distance {
		distance = yMax
	}

	a := geometry.Point{X: float64(x), Y: float64(y)}

	if tiles[y][x].IsOcean {
		return 0
	}

	for d := 0; d < len(tiles)-y; d++ {
		subjects = getTilesAtDistance(d, a, tiles)
		for _, s := range subjects {
			if s.IsOcean {
				return d
			}
		}
	}

	return distance
}

func getTilesAtDistance(distance int, origin geometry.Point, tiles [][]Tile) []Tile {
	var result []Tile

	x := int(origin.X)
	y := int(origin.Y)

	xMax := len(tiles[0]) - 1
	xMin := 0
	yMax := len(tiles) - 1
	yMin := 0

	if x+distance < len(tiles[0])-1 {
		xMax = x + distance
	}

	if x-distance > 0 {
		xMin = x - distance
	}

	if y+distance < len(tiles)-1 {
		yMax = y + distance
	}

	if y-distance > 0 {
		yMax = y - distance
	}

	for j := yMin; j < yMax; j++ {
		result = append(result, tiles[j][xMin])
	}

	for j := yMin; j < yMax; j++ {
		result = append(result, tiles[j][xMax])
	}

	for j := xMin; j < xMax; j++ {
		result = append(result, tiles[yMin][j])
	}

	for j := xMin; j < xMax; j++ {
		result = append(result, tiles[yMax][j])
	}

	return result
}

func generateHumidities(tiles [][]Tile) [][]Tile {
	var humidity int

	for y, row := range tiles {
		for x, h := range row {
			humidity = 100 - h.Altitude
			tiles[y][x].Humidity = humidity
		}
	}

	return tiles
}

func generateTemperatures(tiles [][]Tile, minLatitude int, maxLatitude int) [][]Tile {
	var latitude, modifier float64
	modifier = float64(minLatitude) / float64(maxLatitude)

	for y, row := range tiles {
		latitude = float64(y) * modifier
		for x, h := range row {
			tiles[y][x].Temperature = region.GetTemperature(int(math.Abs(latitude)), h.Altitude)
		}
	}

	return tiles
}

func generateHeightmap(ctx context.Context, width int, height int) [][]float64 {
	var nx, ny, altitude float64
	cells := newCellGrid(width, height)
	noise := opensimplex.New(random.Int63n(ctx, 10000))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			nx = float64(x) / float64(width)
			ny = float64(y) / float64(height)
			altitude = noise.Eval2(2*nx, 2*ny)
			altitude += 0.5 * noise.Eval2(4*nx, 4*ny)
			altitude += 0.25 * noise.Eval2(8*nx, 8*ny)
			altitude += 0.125 * noise.Eval2(16*nx, 16*ny)
			altitude += 0.0675 * noise.Eval2(32*nx, 32*ny)
			altitude += 0.03375 * noise.Eval2(64*nx, 64*ny)
			altitude = altitude * 100

			if altitude < -100 {
				altitude = -100
			}

			if altitude > 100 {
				altitude = 100
			}

			cells[y][x] = altitude
		}
	}

	return cells
}

func generateLand(ctx context.Context, tiles [][]Tile) ([][]Tile, error) {
	height := len(tiles)
	width := len(tiles[0])

	heightmap := generateHeightmap(ctx, width, height)

	for y, row := range heightmap {
		for x, h := range row {
			if h > 0 {
				tiles[y][x].IsOcean = false
			}
			tiles[y][x].Altitude = int(h)
		}
	}

	tiles = generateTemperatures(tiles, 0, 80)
	tiles = generateHumidities(tiles)

	return tiles, nil
}
