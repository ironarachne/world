package random

import (
	"context"
	"testing"
)

func TestString(t *testing.T) {
	items := []string{
		"one",
		"two",
		"three",
	}

	randomItem, err := String(context.Background(), items)
	if err != nil {
		t.Error("Failed to get random item from string slice")
	}
	if randomItem != "one" && randomItem != "two" && randomItem != "three" {
		t.Error("Failed to get valid random string from slice")
	}

	item := []string{
		"one",
	}

	randomItem, err = String(context.Background(), item)
	if err != nil {
		t.Error("Failed to get random item from string slice")
	}
	if randomItem != "one" {
		t.Error("Failed to get valid random string from slice")
	}

	items = []string{}
	_, err = String(context.Background(), items)
	if err == nil {
		t.Error("Failed to error on empty string slice")
	}
}

func TestStringSubset(t *testing.T) {
	items := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
	}

	max := 3

	randomItems, err := StringSubset(context.Background(), items, max)
	if err != nil {
		t.Error("Failed to get random subset of string slice")
	}
	if len(randomItems) != 3 {
		t.Errorf("Failed to get specified number of random elements: got %d instead of 3", len(randomItems))
	}

	max = 7

	randomItems, err = StringSubset(context.Background(), items, max)
	if err != nil {
		t.Error("Failed to get random subset of string slice")
	}
	if len(randomItems) != len(items) {
		t.Errorf("Failed to get adjusted number of random elements: got %d instead of 5", len(randomItems))
	}

	items = []string{}
	_, err = StringSubset(context.Background(), items, max)
	if err == nil {
		t.Error("Failed to error on empty string slice")
	}
}

func TestStringFromThresholdMap(t *testing.T) {
	items := map[string]int{
		"one":   5,
		"two":   5,
		"three": 100,
	}

	randomItem, err := StringFromThresholdMap(context.Background(), items)
	if err != nil {
		t.Error("Failed to get random item from weighted string slice")
	}
	if randomItem != "three" {
		t.Errorf("Failed to get expected 'random' result from weighted string slice: got '%s' instead of 'three'", randomItem)
	}
}
