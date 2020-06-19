/*
Package age implements structures and routines for age calculation and reasoning.
*/
package age

import (
	"context"
	"fmt"

	"github.com/ironarachne/world/pkg/dice"
	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/size"
)

// Category is metadata about a general age
type Category struct {
	Name             string        `json:"name" db:"name"`
	MinAge           int           `json:"age_min" db:"age_min"`
	MaxAge           int           `json:"age_max" db:"age_max"`
	MaleHeightBase   int           `json:"male_height_base" db:"male_height_base"`
	FemaleHeightBase int           `json:"female_height_base" db:"female_height_base"`
	HeightRangeDice  dice.Dice     `json:"height_range_dice" db:"height_range_dice"`
	MaleWeightBase   int           `json:"male_weight_base" db:"male_weight_base"`
	FemaleWeightBase int           `json:"female_weight_base" db:"female_weight_base"`
	WeightModifier   int           `json:"weight_modifier" db:"weight_modifier"`
	WeightRangeDice  dice.Dice     `json:"weight_range_dice" db:"weight_range_dice"`
	SizeCategory     size.Category `json:"size_category" db:"size_category"`
	Commonality      int           `json:"commonality" db:"commonality"`
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
func GetRandomAge(ctx context.Context, ageCategory Category) int {
	return random.Intn(ctx, ageCategory.MaxAge-ageCategory.MinAge) + ageCategory.MinAge
}

// GetRandomHeight returns a random weight given the age category and gender
func GetRandomHeight(ctx context.Context, gender string, ageCategory Category) int {
	var base int
	if gender == "male" {
		base = ageCategory.MaleHeightBase
	} else {
		base = ageCategory.FemaleHeightBase
	}

	result := dice.Roll(ctx, ageCategory.HeightRangeDice)
	height := base + result

	return height
}

// GetRandomWeight returns a random weight given the age category and gender
func GetRandomWeight(ctx context.Context, gender string, ageCategory Category) int {
	var base int
	if gender == "male" {
		base = ageCategory.MaleWeightBase
	} else {
		base = ageCategory.FemaleWeightBase
	}

	result := dice.Roll(ctx, ageCategory.WeightRangeDice)
	weight := base + (ageCategory.WeightModifier * result)

	return weight
}

// GetWeightedAgeCategory returns a random age category for a race
func GetWeightedAgeCategory(ctx context.Context, categories []Category) (Category, error) {
	weights := map[string]int{}

	for _, c := range categories {
		weights[c.Name] = c.Commonality
	}

	name, err := random.StringFromThresholdMap(ctx, weights)
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
