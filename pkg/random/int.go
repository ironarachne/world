package random

import (
	"math/rand"
)

// IntFromThresholdMap returns a random weighted int
func IntFromThresholdMap(items map[int]int) int {
	result := 0
	ceiling := 0
	start := 0
	var thresholds = make(map[int]int)

	for item, weight := range items {
		ceiling += weight
		thresholds[item] = start
		start += weight
	}

	randomValue := rand.Intn(ceiling)

	for item, threshold := range thresholds {
		if threshold <= randomValue {
			result = item
		}
	}

	return result
}
