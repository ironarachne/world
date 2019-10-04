package random

import (
	"fmt"
	"math/rand"
	"strconv"
)

// IntFromThresholdMap returns a random weighted int
func IntFromThresholdMap(items map[int]int) (int, error) {
	ceiling := 0

	for _, weight := range items {
		ceiling += weight
	}

	randomValue := rand.Intn(ceiling)

	for item, weight := range items {
		randomValue -= weight
		if randomValue <= 0 {
			return item, nil
		}
	}

	err := fmt.Errorf("Could not find value for " + strconv.Itoa(randomValue))
	return -1, err
}
