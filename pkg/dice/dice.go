package dice

import "math/rand"

// Dice is a representation of polyhedral dice
type Dice struct {
	Number int
	Sides  int
}

// Roll rolls "num" number of dice with "sides" sides and returns the total result
func Roll(dice Dice) int {
	roll := 0
	result := 0

	for i := 0; i < dice.Number; i++ {
		roll = rand.Intn(dice.Sides) + 1
		result += roll
	}

	return result
}
