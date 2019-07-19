package drink

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/slices"
	"github.com/ironarachne/world/pkg/words"
)

// Drink is an alcoholic drink
type Drink struct {
	Description string
	Appearance  string
	Type        string
	Strength    string
	Base        string
	Ingredients []string
}

func (drink Drink) describe() string {
	description := ""

	if drink.Strength != "" {
		description += drink.Strength + " "
	}

	description += drink.Base + " " + drink.Type

	if len(drink.Ingredients) > 0 {
		description += " made with " + words.CombinePhrases(drink.Ingredients)
	}

	phraseChance := rand.Intn(10)

	phrase := "that's"

	if phraseChance > 6 {
		phrase = "that is"
	}

	description += " " + phrase + " " + drink.Appearance

	return description
}

func getRandomStrength(minimum int) string {
	var strengths []string
	allStrengths := []string{
		"very weak",
		"weak",
		"",
		"strong",
		"very strong",
	}

	for i, s := range allStrengths {
		if i >= minimum {
			strengths = append(strengths, s)
		}
	}

	strength := strengths[rand.Intn(len(strengths))]

	return strength
}

// Random returns a random alcoholic drink
func Random(resources []resource.Resource) Drink {
	var base resource.Resource
	var ingredient string
	var ingredients []string
	var numberOfIngredients int
	var pattern Pattern
	var possibleIngredients []string

	spices := resource.ListOfType("spice", resources)
	herbs := resource.ListOfType("herb", resources)
	fruit := resource.ListOfType("fruit", resources)

	patterns := getValidPatterns(resources)

	if len(patterns) == 0 {
		panic("Not enough valid drink patterns!")
	} else if len(patterns) == 1 {
		pattern = patterns[0]
	} else {
		pattern = patterns[rand.Intn(len(patterns))]
	}

	appearance := random.String(pattern.Descriptors)

	baseOptions := resource.ListOfType(pattern.RequiredBase, resources)

	if len(baseOptions) == 1 {
		base = baseOptions[0]
	} else {
		base = baseOptions[rand.Intn(len(baseOptions))]
	}

	strength := getRandomStrength(pattern.BaseStrength)

	if len(fruit) > 0 {
		for _, f := range fruit {
			possibleIngredients = append(possibleIngredients, f.Name)
		}
	}

	if len(herbs) > 0 {
		for _, h := range herbs {
			possibleIngredients = append(possibleIngredients, h.Name)
		}
	}

	if len(spices) > 0 {
		for _, s := range spices {
			possibleIngredients = append(possibleIngredients, s.Name)
		}
	}

	numberOfIngredients = rand.Intn(4)

	for i := 0; i < numberOfIngredients; i++ {
		ingredient = random.String(possibleIngredients)
		if !slices.StringIn(ingredient, ingredients) {
			ingredients = append(ingredients, ingredient)
		}
	}

	drink := Drink{
		Appearance:  appearance,
		Type:        pattern.Name,
		Strength:    strength,
		Base:        base.Name,
		Ingredients: ingredients,
	}

	drink.Description = drink.describe()

	return drink
}

// RandomSet generates some alcoholic drinks
func RandomSet(numberOfDrinks int, resources []resource.Resource) []Drink {
	var drink Drink
	var drinks []Drink

	for i := 0; i < numberOfDrinks; i++ {
		drink = Random(resources)
		drinks = append(drinks, drink)
	}

	return drinks
}
