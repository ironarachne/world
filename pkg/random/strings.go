package random

import (
	"context"
	"fmt"

	"github.com/ironarachne/world/pkg/slices"
)

// String returns a random string from a slice of strings
func String(ctx context.Context, items []string) (string, error) {
	if len(items) == 0 {
		err := fmt.Errorf("tried to get a random value from an empty slice of strings")
		return "", err
	}
	if len(items) == 1 {
		return items[0], nil
	}
	return items[Intn(ctx, len(items))], nil
}

// StringSubset returns a random slice of strings from a slice of strings
func StringSubset(ctx context.Context, items []string, maxItems int) ([]string, error) {
	result := []string{}
	newItem := ""

	if len(items) == 0 {
		err := fmt.Errorf("tried to get a random value from an empty slice of strings")
		return []string{}, err
	}

	if maxItems > len(items) {
		maxItems = len(items)
	}
	for i := 0; i < maxItems; i++ {
		newItem = items[Intn(ctx, len(items))]
		if !slices.StringIn(newItem, result) {
			result = append(result, newItem)
		} else {
			i--
		}
	}

	return result, nil
}

// StringFromThresholdMap returns a random weighted string
func StringFromThresholdMap(ctx context.Context, items map[string]int) (string, error) {
	ceiling := 0

	for _, weight := range items {
		ceiling += weight
	}

	randomValue := Intn(ctx, ceiling)

	for item, weight := range items {
		randomValue -= weight
		if randomValue <= 0 {
			return item, nil
		}
	}

	err := fmt.Errorf("could not find value for %d", randomValue)
	return "", err
}
