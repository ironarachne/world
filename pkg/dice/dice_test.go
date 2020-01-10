package dice

import "testing"

func TestParseFromString(t *testing.T) {
	diceString := "2d4"

	dice, err := ParseFromString(diceString)
	if err != nil {
		t.Error("Failed to parse dice string")
	}

	if dice.Number != 2 {
		t.Error("Failed to parse 2 as 2 dice")
	}

	if dice.Sides != 4 {
		t.Error("Failed to parse 4 as 4 sides")
	}

	diceString = "22d44"

	dice, err = ParseFromString(diceString)
	if dice.Number != 22 {
		t.Error("Failed to parse 22 as 22 dice")
	}

	if dice.Sides != 44 {
		t.Error("Failed to parse 44 as 44 sides")
	}
}
