package buildings

import (
	"context"
	"fmt"

	"github.com/ironarachne/world/pkg/random"
)

func getRandomDoorMaterial(ctx context.Context) (string, error) {
	materials := []string{
		"light wood",
		"thick wood",
		"thin wood",
		"heavy wood",
	}

	material, err := random.String(ctx, materials)
	if err != nil {
		err = fmt.Errorf("Could not find door material: %w", err)
		return "", err
	}

	return material, nil
}

func getRandomDoorStyle(ctx context.Context) (string, error) {
	material, err := getRandomDoorMaterial(ctx)
	if err != nil {
		err = fmt.Errorf("Could not generate door style: %w", err)
		return "", err
	}
	shape, err := getRandomDoorShape(ctx)
	if err != nil {
		err = fmt.Errorf("Could not generate door style: %w", err)
		return "", err
	}
	style := "made of " + material + " and " + shape

	return style, nil
}

func getRandomDoorShape(ctx context.Context) (string, error) {
	shapes := []string{
		"braced with metal",
		"broad",
		"carved with simple shapes",
		"curved on top",
		"narrow",
		"often possessed of a window in the middle",
		"onion-shaped",
		"ornately carved",
		"rectangular",
		"rounded on top",
		"tall",
	}

	shape, err := random.String(ctx, shapes)
	if err != nil {
		err = fmt.Errorf("Could not get door shape: %w", err)
		return "", err
	}

	return shape, nil
}
