package food

import (
	"fmt"
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

func (style Style) randomMainDish() (string, error) {
	var base string
	var dish string
	var flavorProfile string
	var spices string
	var treatment string

	allFlavors := getFlavors()
	flavors, err := random.StringSubset(allFlavors, 2)
	if err != nil {
		err = fmt.Errorf("Could not generate dish: %w", err)
		return "", err
	}

	flavorProfile = words.CombinePhrases(flavors)
	base, err = random.String(style.CommonBases)
	if err != nil {
		err = fmt.Errorf("Could not generate dish: %w", err)
		return "", err
	}
	treatment, err = style.getRandomTreatment()
	if err != nil {
		err = fmt.Errorf("Could not generate dish: %w", err)
		return "", err
	}

	allSpices, err := random.StringSubset(style.CommonSpices, 3)
	if err != nil {
		err = fmt.Errorf("Could not generate dish: %w", err)
		return "", err
	}
	spices = words.CombinePhrases(allSpices)

	technique, err := random.String(style.CookingTechniques)
	if err != nil {
		err = fmt.Errorf("Could not generate dish: %w", err)
		return "", err
	}
	dish = flavorProfile + " " + technique + " " + base + treatment

	spiceChance := rand.Intn(10)

	if spiceChance > 4 {
		dish += " with " + spices
	}

	return dish, nil
}

func (style Style) randomMainDishes() ([]string, error) {
	var dish string
	var dishes []string
	var err error

	for i := 0; i < 5; i++ {
		dish, err = style.randomMainDish()
		if err != nil {
			err = fmt.Errorf("Could not generate dishes: %w", err)
			return []string{}, err
		}

		if !slices.StringIn(dish, dishes) {
			dishes = append(dishes, dish)
		}
	}

	return dishes, nil
}
