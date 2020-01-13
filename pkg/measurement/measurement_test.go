package measurement

import (
	"strconv"
	"testing"
)

func TestInchesToFeet(t *testing.T) {
	feet, inches := InchesToFeet(72)

	if feet != 6 {
		t.Error("failed to convert 72 inches to 6 feet, got " + strconv.Itoa(feet) + " instead")
	}

	if inches != 0 {
		t.Error("failed to convert 72 inches to 6 feet, 0 inches, got " + strconv.Itoa(inches) + " instead")
	}
}

func TestToString(t *testing.T) {
	height := ToString(72)

	if height != "6'0\"" {
		t.Error("failed to display 72 inches as 6'0\", got " + height + " instead")
	}
}
