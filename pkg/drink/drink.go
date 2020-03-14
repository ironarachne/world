/*
Package drink provides methods and tools for generating
fantasy drink styles.
*/
package drink

import (
	"fmt"
	"math/rand"

	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/slices"
	"github.com/ironarachne/world/pkg/words"
)

// Drink is an alcoholic drink
type Drink struct {
	Description string   `json:"description"`
	Appearance  string   `json:"appearance"`
	Type        string   `json:"type"`
	Strength    string   `json:"strength"`
	Base        string   `json:"base"`
	Ingredients []string `json:"ingredients"`
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
func Random(resources []resource.Resource) (Drink, error) {
	var base resource.Resource
	var ingredient string
	var ingredients []string
	var numberOfIngredients int
	var pattern Pattern
	var possibleIngredients []string
	var err error

	spices := resource.ByTag("spice", resources)
	herbs := resource.ByTag("herb", resources)
	fruit := resource.ByTag("fruit", resources)

	patterns := getValidPatterns(resources)

	if len(patterns) == 0 {
		err = fmt.Errorf("not enough valid drink patterns for given resources")
		return Drink{}, err
	} else if len(patterns) == 1 {
		pattern = patterns[0]
	} else {
		pattern = patterns[rand.Intn(len(patterns))]
	}

	appearance, err := random.String(pattern.Descriptors)
	if err != nil {
		err = fmt.Errorf("failed to generate random alcoholic drink: %w", err)
		return Drink{}, err
	}

	baseOptions := resource.ByTag(pattern.RequiredBase, resources)

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
		ingredient, err = random.String(possibleIngredients)
		if err != nil {
			err = fmt.Errorf("failed to get ingredient for random alcoholic drink: %w", err)
			return Drink{}, err
		}
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

	return drink, nil
}

// RandomSet generates some alcoholic drinks
func RandomSet(numberOfDrinks int, resources []resource.Resource) ([]Drink, error) {
	var drink Drink
	var drinks []Drink
	var err error

	for i := 0; i < numberOfDrinks; i++ {
		drink, err = Random(resources)
		if err != nil {
			err = fmt.Errorf("failed to generate alcoholic drinks: %w", err)
			return []Drink{}, err
		}
		drinks = append(drinks, drink)
	}

	return drinks, nil
}
