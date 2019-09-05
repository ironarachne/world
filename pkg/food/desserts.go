package food

import (
	"fmt"

	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/slices"
)

func (style Style) randomDessert() (string, error) {
	var dessert string

	main, err := random.String(style.CommonDessertBases)
	if err != nil {
		err = fmt.Errorf("Could not generate dessert: %w", err)
		return "", err
	}
	treatment, err := random.String(style.DessertTypes)
	if err != nil {
		err = fmt.Errorf("Could not generate drink: %w", err)
		return "", err
	}

	if treatment == "fruit" {
		dessert = main
	} else {
		dessert = main + " " + treatment
	}

	return dessert, nil
}

func (style Style) randomDesserts() ([]string, error) {
	var dessert string
	var desserts []string
	var err error

	for i := 0; i < 5; i++ {
		dessert, err = style.randomDessert()
		if err != nil {
			err = fmt.Errorf("Could not generate desserts: %w", err)
			return []string{}, err
		}

		if !slices.StringIn(dessert, desserts) {
			desserts = append(desserts, dessert)
		}
	}

	return desserts, nil
}
