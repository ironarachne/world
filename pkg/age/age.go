/*
Package age implements structures and routines for age calculation and reasoning.
*/
package age

import (
	"fmt"
	"math/rand"

	"github.com/ironarachne/world/pkg/dice"
	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/size"
)

// Category is metadata about a general age
type Category struct {
	Name             string
	MinAge           int
	MaxAge           int
	MaleHeightBase   int
	FemaleHeightBase int
	HeightRangeDice  dice.Dice
	MaleWeightBase   int
	FemaleWeightBase int
	WeightModifier   int
	WeightRangeDice  dice.Dice
	SizeCategory     size.Category
	Commonality      int
}

// GetCategoryByName returns the age category given a name
func GetCategoryByName(name string, categories []Category) Category {
	for _, c := range categories {
		if c.Name == name {
			return c
		}
	}

	return Category{}
}

// GetCategoryFromAge returns the age category that fits the age for the given set of age categories
func GetCategoryFromAge(years int, categories []Category) Category {
	ageCategory := Category{}

	for _, category := range categories {
		if years >= category.MinAge && years <= category.MaxAge {
			ageCategory = category
		}
	}

	return ageCategory
}

// GetRandomAge returns a random age in years within the age category range
func GetRandomAge(ageCategory Category) int {
	return rand.Intn(ageCategory.MaxAge-ageCategory.MinAge) + ageCategory.MinAge
}

// GetRandomHeight returns a random weight given the age category and gender
func GetRandomHeight(gender string, ageCategory Category) int {
	var base int
	if gender == "male" {
		base = ageCategory.MaleHeightBase
	} else {
		base = ageCategory.FemaleHeightBase
	}

	result := dice.Roll(ageCategory.HeightRangeDice)
	height := base + result

	return height
}

// GetRandomWeight returns a random weight given the age category and gender
func GetRandomWeight(gender string, ageCategory Category) int {
	var base int
	if gender == "male" {
		base = ageCategory.MaleWeightBase
	} else {
		base = ageCategory.FemaleWeightBase
	}

	result := dice.Roll(ageCategory.WeightRangeDice)
	weight := base + (ageCategory.WeightModifier * result)

	return weight
}

// GetWeightedAgeCategory returns a random age category for a race
func GetWeightedAgeCategory(categories []Category) (Category, error) {
	weights := map[string]int{}

	for _, c := range categories {
		weights[c.Name] = c.Commonality
	}

	name, err := random.StringFromThresholdMap(weights)
	if err != nil {
		err = fmt.Errorf("Failed to get weighted age category: %w", err)
		return Category{}, err
	}

	for _, c := range categories {
		if c.Name == name {
			return c, nil
		}
	}

	err = fmt.Errorf("failed to get weighted age category")
	return Category{}, err
}
