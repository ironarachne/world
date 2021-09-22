package buildings

import (
	"context"
	"fmt"
	"github.com/ironarachne/world/pkg/random"
)

func getRandomRoofStyle(ctx context.Context) (string, error) {
	material, err := getRandomRoofMaterial(ctx)
	if err != nil {
		err = fmt.Errorf("Failed to generate random roof shape: %w", err)
		return "", err
	}
	shape, err := getRandomRoofShape(ctx)
	if err != nil {
		err = fmt.Errorf("Failed to generate random roof style: %w", err)
		return "", err
	}

	return shape + " and made of " + material, nil
}

func getRandomRoofShape(ctx context.Context) (string, error) {
	shapes := map[string]int{
		"flat":            20,
		"slanted":         6,
		"conical":         3,
		"steeply slanted": 1,
		"onion-shaped":    1,
	}

	shape, err := random.StringFromThresholdMap(ctx, shapes)
	if err != nil {
		err = fmt.Errorf("Failed to generate random roof shape: %w", err)
		return "", err
	}

	return shape, nil
}

func getRandomRoofMaterial(ctx context.Context) (string, error) {
	materials := map[string]int{
		"beaten metal":    1,
		"ceramic tiles":   5,
		"stone tiles":     3,
		"thatch":          15,
		"wooden shingles": 5,
		"wooden tiles":    3,
	}

	material, err := random.StringFromThresholdMap(ctx, materials)
	if err != nil {
		err = fmt.Errorf("Failed to generate random roof material: %w", err)
		return "", err
	}
	return material, nil
}
