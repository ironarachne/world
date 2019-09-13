package race

import (
	"fmt"
	"math/rand"

	"github.com/ironarachne/world/pkg/random"
)

// AgeCategory is metadata about a general age
type AgeCategory struct {
	Name        string
	MinAge      int
	MaxAge      int
	Commonality int
}

// GetAgeCategoryByName returns the age category for a race given a name
func (race Race) GetAgeCategoryByName(name string) AgeCategory {
	cats := race.AgeCategories

	for _, c := range cats {
		if c.Name == name {
			return c
		}
	}

	return AgeCategory{}
}

// GetAgeCategoryFromAge returns the age category that fits the age for the given race
func (race Race) GetAgeCategoryFromAge(age int) AgeCategory {
	ageCategory := AgeCategory{}

	ageCategories := race.AgeCategories

	for _, category := range ageCategories {
		if age >= category.MinAge && age <= category.MaxAge {
			ageCategory = category
		}
	}

	return ageCategory
}

// GetRandomAge returns a random age in years within the age category range
func GetRandomAge(ageCategory AgeCategory) int {
	return rand.Intn(ageCategory.MaxAge-ageCategory.MinAge) + ageCategory.MinAge
}

// GetWeightedAgeCategory returns a random age category for a race
func (race Race) GetWeightedAgeCategory() (AgeCategory, error) {
	ageCategories := race.AgeCategories

	weights := map[string]int{}

	for _, c := range ageCategories {
		weights[c.Name] = c.Commonality
	}

	name, err := random.StringFromThresholdMap(weights)
	if err != nil {
		err = fmt.Errorf("Failed to get weighted age category: %w", err)
		return AgeCategory{}, err
	}

	for _, c := range ageCategories {
		if c.Name == name {
			return c, nil
		}
	}

	err = fmt.Errorf("Failed to get weighted age category!")
	return AgeCategory{}, err
}
