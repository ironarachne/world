package world

import (
	"github.com/ojrac/opensimplex-go"
	"math"
	"math/rand"
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

func generateHeightmap(width int, height int) [][]float64 {
	var nx, ny, altitude float64
	cells := newCellGrid(width, height)
	noise := opensimplex.New(rand.Int63n(10000))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			nx = float64(x)/float64(width)-0.5
			ny = float64(y)/float64(height)-0.5
			altitude = noise.Eval2(3 * nx, 3 * ny) + (0.5 * noise.Eval2(18 * nx, 18 * ny)) + (0.25 * noise.Eval2(36 * nx, 36 * ny)) + (0.125 * noise.Eval2(72 * nx, 72 * ny))
			altitude = math.Pow(altitude, 0.74)
			cells[y][x] = altitude
		}
	}

	return cells
}

func generateLand(tiles [][]Tile) [][]Tile {
	height := len(tiles)
	width := len(tiles[0])

	heightmap := generateHeightmap(width, height)
	adjustment := normalizeHeightmap(heightmap)

	for y, row := range adjustment {
		for x, h := range row {
			if h > 20 {
				tiles[y][x].IsOcean = false
			}
			tiles[y][x].Altitude = int(h)
		}
	}

	return tiles
}

func normalizeHeightmap(cells [][]float64) [][]float64 {
	max := getMax(cells)
	min := getMin(cells)

	totalRange := max - min

	for y, row := range cells {
		for x, v := range row {
			cells[y][x] = ((v-min)/totalRange) * 100
		}
	}

	return cells
}

func getMax(cells [][]float64) float64 {
	max := 0.0

	for _, row := range cells {
		for _, v := range row {
			if v > max {
				max = v
			}
		}
	}

	return max
}

func getMin(cells [][]float64) float64 {
	min := 0.0

	for _, row := range cells {
		for _, v := range row {
			if v < min {
				min = v
			}
		}
	}

	return min
}