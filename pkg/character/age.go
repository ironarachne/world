package character

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/math"
	"github.com/ironarachne/world/pkg/random"
)

// AgeCategory is metadata about a general age
type AgeCategory struct {
	Name        string
	MinAge      int
	MaxAge      int
	Commonality int
}

func getAgeCategoryByName(name string) AgeCategory {
	cats := getAllAgeCategories()

	for _, c := range cats {
		if c.Name == name {
			return c
		}
	}

	return AgeCategory{}
}

func getAgeCategoryFromAge(age int) AgeCategory {
	ageCategory := AgeCategory{}

	ageCategories := getAllAgeCategories()

	for _, category := range ageCategories {
		if age >= category.MinAge && age <= category.MaxAge {
			ageCategory = category
		}
	}

	return ageCategory
}

func getAgeFromParents(parents Couple) (int, AgeCategory) {
	lowestAge := math.Min(parents.Partner1.Age, parents.Partner2.Age)

	childAge := rand.Intn(lowestAge-18) + 1
	childAgeCategory := getAgeCategoryFromAge(childAge)

	return childAge, childAgeCategory
}

func getAllAgeCategories() []AgeCategory {
	ageCategories := []AgeCategory{
		AgeCategory{
			Name:        "adult",
			MinAge:      26,
			MaxAge:      69,
			Commonality: 12,
		},
		AgeCategory{
			Name:        "elderly",
			MinAge:      70,
			MaxAge:      110,
			Commonality: 1,
		},
		AgeCategory{
			Name:        "young adult",
			MinAge:      20,
			MaxAge:      25,
			Commonality: 2,
		},
		AgeCategory{
			Name:        "teenager",
			MinAge:      13,
			MaxAge:      19,
			Commonality: 1,
		},
		AgeCategory{
			Name:        "child",
			MinAge:      2,
			MaxAge:      12,
			Commonality: 1,
		},
		AgeCategory{
			Name:        "infant",
			MinAge:      0,
			MaxAge:      1,
			Commonality: 1,
		},
	}

	return ageCategories
}

func getRandomAge(ageCategory AgeCategory) int {
	return rand.Intn(ageCategory.MaxAge-ageCategory.MinAge) + ageCategory.MinAge
}

func getWeightedAgeCategory() AgeCategory {
	ageCategories := getAllAgeCategories()

	weights := map[string]int{}

	for _, c := range ageCategories {
		weights[c.Name] = c.Commonality
	}

	name := random.StringFromThresholdMap(weights)

	for _, c := range ageCategories {
		if c.Name == name {
			return c
		}
	}

	return AgeCategory{}
}

// ChangeAge changes the age and age category of a character
func (character Character) ChangeAge(newAge int) Character {
	c := getAgeCategoryFromAge(newAge)
	character.Age = newAge
	character.AgeCategory = c
	character.Height = character.randomHeight()
	character.Weight = character.randomWeight()

	return character
}
