package words

import "testing"

func TestNumberToWord(t *testing.T) {
	number := 0

	result, err := NumberToWord(number)
	if err != nil {
		t.Error("failed to turn 0 into zero")
	}
	if result != "zero" {
		t.Error("failed to turn 0 into zero")
	}

	number = 1

	result, err = NumberToWord(number)
	if err != nil {
		t.Error("failed to turn 1 into one")
	}
	if result != "one" {
		t.Error("failed to turn 1 into one")
	}
}
