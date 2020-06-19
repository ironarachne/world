package random

import (
	"context"
	"fmt"
)

// IntFromThresholdMap returns a random weighted int
func IntFromThresholdMap(ctx context.Context, items map[int]int) (int, error) {
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
	return -1, err
}
