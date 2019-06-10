package buildings

import (
	"github.com/ironarachne/world/pkg/random"
)

func getRandomRoofStyle() string {
	material := getRandomRoofMaterial()
	shape := getRandomRoofShape()

	return shape + " and made of " + material
}

func getRandomRoofShape() string {
	shapes := map[string]int{
		"flat":            20,
		"slanted":         6,
		"conical":         3,
		"steeply slanted": 1,
		"onion-shaped":    1,
	}

	return random.StringFromThresholdMap(shapes)
}

func getRandomRoofMaterial() string {
	materials := map[string]int{
		"beaten metal":    1,
		"ceramic tiles":   5,
		"stone tiles":     3,
		"thatch":          15,
		"wooden shingles": 5,
		"wooden tiles":    3,
	}

	return random.StringFromThresholdMap(materials)
}
