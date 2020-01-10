package math

import "testing"

func TestMax(t *testing.T) {
	a := 1
	b := 3

	result := Max(a, b)
	if result != b {
		t.Error("Failed to recognize 3 as greater than 1")
	}

	result = Max(b, a)
	if result != b {
		t.Error("Failed to recognize 3 as greater than 1")
	}
}

func TestMin(t *testing.T) {
	a := 1
	b := 3

	result := Min(a, b)
	if result != a {
		t.Error("Failed to recognize 1 as smaller than 3")
	}

	result = Min(b, a)
	if result != a {
		t.Error("Failed to recognize 1 as smaller than 3")
	}
}
