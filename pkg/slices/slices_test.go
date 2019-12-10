package slices

import "testing"

func TestStringIn(t *testing.T) {
	slice := []string{
		"one",
		"two",
		"three",
	}

	if !StringIn("two", slice) {
		t.Error("failed to find element two in slice that contains it")
	}

	if StringIn("four", slice) {
		t.Error("found string in slice that doesn't contain it")
	}
}

func TestStringSliceWithin(t *testing.T) {
	slice1 := []string{
		"one",
		"two",
		"three",
	}

	slice2 := []string{
		"one",
		"two",
		"three",
		"four",
	}

	slice3 := []string{
		"six",
		"nine",
		"twelve",
	}

	if !StringSliceWithin(slice1, slice2) {
		t.Error("failed to find slice within slice that contains it")
	}

	if StringSliceWithin(slice1, slice3) {
		t.Error("found slice in slice that does not contain it")
	}
}

func TestStringSlicePartlyWithin(t *testing.T) {
	slice1 := []string{
		"one",
		"two",
		"three",
	}

	slice2 := []string{
		"one",
		"two",
		"four",
	}

	slice3 := []string{
		"six",
		"nine",
		"twelve",
	}

	if !StringSlicePartlyWithin(slice1, slice2) {
		t.Error("failed to find part of slice within target slice that contains some of it")
	}

	if StringSlicePartlyWithin(slice1, slice3) {
		t.Error("found slice elements in target slice that does not contain any of its elements")
	}
}
