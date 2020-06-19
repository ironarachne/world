package food

import (
	"context"
	"fmt"

	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/slices"
	"github.com/ironarachne/world/pkg/words"
)

const dishError = "failed to generate dish: %w"

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

func (style Style) randomMainDish(ctx context.Context) (string, error) {
	var base string
	var dish string
	var flavorProfile string
	var spices string
	var treatment string

	allFlavors := getFlavors()
	flavors, err := random.StringSubset(ctx, allFlavors, 2)
	if err != nil {
		err = fmt.Errorf(dishError, err)
		return "", err
	}

	flavorProfile = words.CombinePhrases(flavors)
	base, err = random.String(ctx, style.CommonBases)
	if err != nil {
		err = fmt.Errorf(dishError, err)
		return "", err
	}
	treatment, err = style.getRandomTreatment(ctx)
	if err != nil {
		err = fmt.Errorf(dishError, err)
		return "", err
	}

	allSpices, err := random.StringSubset(ctx, style.CommonSpices, 3)
	if err != nil {
		err = fmt.Errorf(dishError, err)
		return "", err
	}
	spices = words.CombinePhrases(allSpices)

	technique, err := random.String(ctx, style.CookingTechniques)
	if err != nil {
		err = fmt.Errorf(dishError, err)
		return "", err
	}
	dish = flavorProfile + " " + technique + " " + base + treatment

	spiceChance := random.Intn(ctx, 10)

	if spiceChance > 4 {
		dish += " with " + spices
	}

	return dish, nil
}

func (style Style) randomMainDishes(ctx context.Context) ([]string, error) {
	var dish string
	var dishes []string
	var err error

	for i := 0; i < 5; i++ {
		dish, err = style.randomMainDish(ctx)
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
