package food

import (
	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/slices"
)

func (style Style) randomDishes() []string {
	var dishes []string
	var dish string

	flavors := []string{
		"aromatic",
		"bitter",
		"pungent",
		"salty",
		"saucy",
		"savory",
		"sour",
		"spicy",
		"sweet",
	}

	for i := 0; i < 5; i++ {
		dish = random.String(flavors) + " " + random.String(style.CookingTechniques) + " " + random.String(style.CommonBases) + " with " + random.String(style.CommonSpices)

		if !slices.StringIn(dish, dishes) {
			dishes = append(dishes, dish)
		}
	}

	return dishes
}