package world

import (
	"context"

	"github.com/ironarachne/world/pkg/random"
)

const OCEAN sir = 0
const LAND sir = 2

type sir int // 0 = susceptible, 1 = infected, 2 = retired

func generateLandSIR(ctx context.Context, tiles [][]Tile) [][]Tile {
	var x, y, transmission, recovered, infected, susceptible int
	var adjacent []coord
	var infectedTiles []coord

	numberOfInfectedTiles := len(tiles) / 30
	transmissionRate := 50

	height := len(tiles)
	width := len(tiles[0])

	processed := make([][]Tile, height)
	for i := 0; i < height; i++ {
		processed[i] = make([]Tile, width)
	}

	state := make([][]sir, height)
	for i := 0; i < height; i++ {
		state[i] = make([]sir, width)
	}

	for i := 0; i < numberOfInfectedTiles; i++ {
		x = random.Intn(ctx, width)
		y = random.Intn(ctx, height)

		state[y][x] = 1
	}

	infected = numberOfInfectedTiles

	for infected > 0 {
		infectedTiles = []coord{}
		for y, r := range state {
			for x, c := range r {
				if c == 1 {
					infectedTiles = append(infectedTiles, coord{x: x, y: y})
				}
			}
		}

		for _, i := range infectedTiles {
			adjacent = i.adjacent(width, height)
			for _, a := range adjacent {
				if state[a.y][a.x] == 0 {
					transmission = random.Intn(ctx, 100)
					if transmission < transmissionRate {
						state[a.y][a.x] = 1
					}
				}
			}
			state[i.y][i.x] = 2
		}

		recovered = 0
		infected = 0
		susceptible = 0

		for _, r := range state {
			for _, c := range r {
				if c == LAND {
					recovered++
				} else if c == 1 {
					infected++
				} else {
					susceptible++
				}
			}
		}
	}

	state = removeArtifactOceanTiles(state)

	for i := 0; i < height-1; i++ {
		for j := 0; j < width-1; j++ {
			if state[i][j] == 2 {
				processed[i][j].IsOcean = false
			} else {
				processed[i][j].IsOcean = true
			}
		}
	}

	return processed
}

func findCoordsOfType(tileType sir, state [][]sir) []coord {
	var filteredCoords []coord

	for y, r := range state {
		for x, c := range r {
			if c == tileType {
				filteredCoords = append(filteredCoords, coord{x: x, y: y})
			}
		}
	}

	return filteredCoords
}

func removeArtifactOceanTiles(state [][]sir) [][]sir {
	coords := findArtifactsOfType(OCEAN, state)

	for _, c := range coords {
		state[c.y][c.x] = LAND
	}

	return state
}

func findArtifactsOfType(tileType sir, state [][]sir) []coord {
	var filteredCoords []coord

	coords := findCoordsOfType(tileType, state)

	for _, c := range coords {
		if !isContiguous(12, c, state) {
			filteredCoords = append(filteredCoords, c)
		}
	}

	return filteredCoords
}

func isContiguous(threshold int, c coord, state [][]sir) bool {
	var adjacentTiles []coord
	var contiguousTiles []coord

	width := len(state[0])
	height := len(state)

	checkNext := true

	workingTile := c

	stateValue := state[c.y][c.x]

	for i := 1; i < threshold; i++ {
		adjacentTiles = workingTile.adjacent(width, height)

		if checkNext {
			checkNext = false
			for _, a := range adjacentTiles {
				if state[a.y][a.x] == stateValue {
					if !coordInSlice(a, contiguousTiles) {
						contiguousTiles = append(contiguousTiles, a)
					}
					workingTile = a
					checkNext = true
				}
			}
		}

		if len(contiguousTiles) >= threshold {
			return true
		}
	}

	return false
}

func smoothState(ctx context.Context, state [][]sir) [][]sir {
	var countOfSimilar, chanceOfChange int
	var adjacent []coord
	var d coord

	height := len(state)
	width := len(state[0])

	for y, r := range state {
		for x, c := range r {
			d = coord{x: x, y: y}
			countOfSimilar = 0
			adjacent = d.adjacent(width, height)
			for _, a := range adjacent {
				if state[a.y][a.x] == c {
					countOfSimilar++
				}
			}
			if countOfSimilar < 2 {
				chanceOfChange = random.Intn(ctx, 100)
				if chanceOfChange > 20 {
					if c == 2 {
						state[y][x] = 0
					} else {
						state[y][x] = 2
					}
				}
			}
		}
	}

	return state
}
