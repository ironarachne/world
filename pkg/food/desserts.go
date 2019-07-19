package food

import (
	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/slices"
)

func (style Style) randomDessert() string {
	var dessert string

	main := random.String(style.CommonDessertBases)
	treatment := random.String(style.DessertTypes)

	if treatment == "fruit" {
		dessert = main
	} else {
		dessert = main + " " + treatment
	}

	return dessert
}

func (style Style) randomDesserts() []string {
	var dessert string
	var desserts []string

	for i := 0; i < 5; i++ {
		dessert = style.randomDessert()

		if !slices.StringIn(dessert, desserts) {
			desserts = append(desserts, dessert)
		}
	}

	return desserts
}
