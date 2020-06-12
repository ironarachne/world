package geometry

import (
	"context"

	"github.com/ironarachne/world/pkg/random"
)

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
func RandomDirection(ctx context.Context) int {
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

	direction := directions[random.Intn(ctx, len(directions))]

	return direction
}

// DirectionWord returns the string description of a direction
func DirectionWord(direction int) string {
	if direction == NORTH {
		return "north"
	} else if direction == NORTHEAST {
		return "northeast"
	} else if direction == EAST {
		return "east"
	} else if direction == SOUTHEAST {
		return "southeast"
	} else if direction == SOUTH {
		return "south"
	} else if direction == SOUTHWEST {
		return "southwest"
	} else if direction == WEST {
		return "west"
	}

	return "northwest"
}
