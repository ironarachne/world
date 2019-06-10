package buildings

import (
	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/words"
)

func getRandomWindowStyle() string {
	descriptors := []string{}
	descriptors = append(descriptors, getRandomWindowSize())
	descriptors = append(descriptors, getRandomWindowShape())

	return words.CombinePhrases(descriptors)
}

func getRandomWindowShape() string {
	shapes := map[string]int{
		"rectangular": 20,
		"square":      6,
		"round":       3,
		"oval":        1,
		"triangular":  1,
	}

	return random.StringFromThresholdMap(shapes)
}

func getRandomWindowSize() string {
	sizes := map[string]int{
		"large":  5,
		"medium": 15,
		"small":  3,
	}

	return random.StringFromThresholdMap(sizes)
}
