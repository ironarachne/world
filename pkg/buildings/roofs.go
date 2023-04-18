package buildings

import (
	"fmt"
	"github.com/ironarachne/world/pkg/random"
)

func getRandomRoofStyle() (string, error) {
	material, err := getRandomRoofMaterial()
	if err != nil {
		err = fmt.Errorf("Failed to generate random roof shape: %w", err)
		return "", err
	}
	shape, err := getRandomRoofShape()
	if err != nil {
		err = fmt.Errorf("Failed to generate random roof style: %w", err)
		return "", err
	}

	return shape + " and made of " + material, nil
}

func getRandomRoofShape() (string, error) {
	shapes := map[string]int{
		"flat":            20,
		"slanted":         6,
		"conical":         3,
		"steeply slanted": 1,
		"onion-shaped":    1,
	}

	shape, err := random.StringFromThresholdMap(shapes)
	if err != nil {
		err = fmt.Errorf("Failed to generate random roof shape: %w", err)
		return "", err
	}

	return shape, nil
}

func getRandomRoofMaterial() (string, error) {
	materials := map[string]int{
		"beaten metal":    1,
		"ceramic tiles":   5,
		"stone tiles":     3,
		"thatch":          15,
		"wooden shingles": 5,
		"wooden tiles":    3,
	}

	material, err := random.StringFromThresholdMap(materials)
	if err != nil {
		err = fmt.Errorf("Failed to generate random roof material: %w", err)
		return "", err
	}
	return material, nil
}
