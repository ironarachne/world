package grid

import "math/rand"

const NORTH = 0
const NORTHEAST = 1
const EAST = 2
const SOUTHEAST = 3
const SOUTH = 4
const SOUTHWEST = 5
const WEST = 6
const NORTHWEST = 7

// OppositeDirection returns the opposite of a given compass direction
func OppositeDirection(original int) int {
	opposite := 8 - original
	if original == NORTH {
		opposite = SOUTH
	} else if original == SOUTH {
		opposite = NORTH
	}

	return opposite
}

// RandomDirection returns a random compass direction
func RandomDirection() int {
	directions := []int{
		NORTH,
		NORTHEAST,
		EAST,
		SOUTHEAST,
		SOUTH,
		SOUTHWEST,
		WEST,
		NORTHWEST,
	}

	direction := directions[rand.Intn(len(directions))]

	return direction
}
