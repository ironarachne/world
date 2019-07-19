package food

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/slices"
	"github.com/ironarachne/world/pkg/words"
)

func getFlavors() []string {
	return []string{
		"aromatic",
		"bitter",
		"pungent",
		"salty",
		"savory",
		"sour",
		"spicy",
		"sweet",
	}
}

func getNoodles(flour string) []string {
	noodles := []string{}

	noodleStyles := []string{
		"wide",
		"flat",
		"thin",
		"long",
		"narrow",
		"spiral",
		"round",
	}

	for _, s := range noodleStyles {
		noodles = append(noodles, s+" "+flour+" noodle")
	}

	return noodles
}

func (style Style) randomMainDish() string {
	var base string
	var dish string
	var flavorProfile string
	var spices string
	var treatment string

	flavors := getFlavors()

	flavorProfile = words.CombinePhrases(random.StringSubset(flavors, 2))
	base = random.String(style.CommonBases)
	treatment = style.getRandomTreatment()
	spices = words.CombinePhrases(random.StringSubset(style.CommonSpices, 3))
	dish = flavorProfile + " " + random.String(style.CookingTechniques) + " " + base + treatment

	spiceChance := rand.Intn(10)

	if spiceChance > 4 {
		dish += " with " + spices
	}

	return dish
}

func (style Style) randomMainDishes() []string {
	var dish string
	var dishes []string

	for i := 0; i < 5; i++ {
		dish = style.randomMainDish()

		if !slices.StringIn(dish, dishes) {
			dishes = append(dishes, dish)
		}
	}

	return dishes
}
