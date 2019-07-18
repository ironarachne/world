package random

import (
	"math/rand"
)

// DiceTotal rolls "num" number of dice with "sides" sides and returns the total result
func DiceTotal(num, sides int) int {
	roll := 0
	result := 0

	for i := 0; i < num; i++ {
		roll = rand.Intn(sides) + 1
		result += roll
	}

	return result
}
