/*
Package dice provides methods for simulating the rolling of polyhedral dice.
*/
package dice

import (
	"context"
	"fmt"
	"regexp"
	"strconv"

	"github.com/ironarachne/world/pkg/random"
)

// Dice is a representation of polyhedral dice
type Dice struct {
	Number int `json:"number" db:"number"`
	Sides  int `json:"sides" db:"sides"`
}

// Roll rolls "num" number of dice with "sides" sides and returns the total result
func Roll(ctx context.Context, dice Dice) int {
	roll := 0
	result := 0

	for i := 0; i < dice.Number; i++ {
		roll = random.Intn(ctx, dice.Sides) + 1
		result += roll
	}

	return result
}

// ParseFromString returns a Dice struct parsed from a string
func ParseFromString(input string) (Dice, error) {
	n := regexp.MustCompile(`\d+`)
	results := n.FindAllString(input, -1)
	if len(results) != 2 {
		err := fmt.Errorf("bad dice string: %s", input)
		return Dice{}, err
	}

	number, err := strconv.Atoi(results[0])
	if err != nil {
		err = fmt.Errorf("failed to get number of dice: %w", err)
		return Dice{}, err
	}

	sides, err := strconv.Atoi(results[1])
	if err != nil {
		err = fmt.Errorf("failed to get sides of dice: %w", err)
		return Dice{}, err
	}

	dice := Dice{
		Number: number,
		Sides:  sides,
	}

	return dice, nil
}
