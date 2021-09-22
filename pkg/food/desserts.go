package food

import (
	"context"
	"fmt"

	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/slices"
)

func (style Style) randomDessert(ctx context.Context) (string, error) {
	var dessert string

	main, err := random.String(ctx, style.CommonDessertBases)
	if err != nil {
		err = fmt.Errorf("failed to generate dessert base: %w", err)
		return "", err
	}
	treatment, err := random.String(ctx, style.DessertTypes)
	if err != nil {
		err = fmt.Errorf("failed to generate dessert type: %w", err)
		return "", err
	}

	if treatment == "fruit" {
		dessert = main
	} else {
		dessert = main + " " + treatment
	}

	return dessert, nil
}

func (style Style) randomDesserts(ctx context.Context) ([]string, error) {
	var dessert string
	var desserts []string
	var err error

	for i := 0; i < 5; i++ {
		dessert, err = style.randomDessert(ctx)
		if err != nil {
			err = fmt.Errorf("failed to generate desserts: %w", err)
			return []string{}, err
		}

		if !slices.StringIn(dessert, desserts) {
			desserts = append(desserts, dessert)
		}
	}

	return desserts, nil
}
