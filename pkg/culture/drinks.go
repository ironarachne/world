package culture

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/climate"
	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/slices"
)

// Drink is an alcoholic drink
type Drink struct {
	Name        string
	Strength    string
	Base        string
	Ingredients []string
}

func (culture Culture) generateDrinks() []Drink {
	var drink Drink
	var drinks []Drink
	var types []string
	var ingredient string
	var numberOfIngredients int
	var possibleIngredients []string

	spices := climate.ListResourcesOfType("spice", culture.HomeClimate.Resources)
	fruit := climate.ListResourcesOfType("fruit", culture.HomeClimate.Resources)
	grain := climate.ListResourcesOfType("grain", culture.HomeClimate.Resources)

	strengths := []string{
		"very strong",
		"strong",
		"weak",
		"very weak",
	}

	kinds := []string{
		"wine",
		"liquor",
		"ale",
		"beer",
		"brandy",
	}

	if len(fruit) > 0 {
		for _, f := range fruit {
			types = append(types, f.Name)
			possibleIngredients = append(possibleIngredients, f.Name)
		}
	}

	if len(grain) > 0 {
		for _, g := range grain {
			types = append(types, g.Name)
		}
	}

	if len(spices) > 0 {
		for _, s := range spices {
			possibleIngredients = append(possibleIngredients, s.Name)
		}
	}

	if len(types) > 0 {
		for i := 0; i < 2; i++ {
			drink = Drink{}
			drink.Base = random.String(types)
			drink.Strength = random.String(strengths)
			numberOfIngredients = rand.Intn(4) + 1
			drink.Ingredients = append(drink.Ingredients, drink.Base)
			for i := 0; i < numberOfIngredients; i++ {
				ingredient = random.String(possibleIngredients)
				if !slices.StringIn(ingredient, drink.Ingredients) {
					drink.Ingredients = append(drink.Ingredients, ingredient)
				}
			}
			drink.Name = drink.Base + " " + random.String(kinds)
			drinks = append(drinks, drink)
		}
	}

	return drinks
}
