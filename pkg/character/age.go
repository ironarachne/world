package character

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/age"
	"github.com/ironarachne/world/pkg/math"
)

func getAgeFromParents(parents Couple) (int, age.Category) {
	lowestAge := math.Min(parents.Partner1.Age, parents.Partner2.Age)

	childAge := rand.Intn(lowestAge-18) + 1
	childAgeCategory := age.GetCategoryFromAge(childAge, parents.Partner1.Race.AgeCategories)

	return childAge, childAgeCategory
}

// ChangeAge changes the age and age category of a character
func (character Character) ChangeAge(newAge int) Character {
	c := age.GetCategoryFromAge(newAge, character.Race.AgeCategories)
	character.Age = newAge
	character.AgeCategory = c
	character.Height = age.GetRandomHeight(character.Gender.Name, character.AgeCategory)
	character.Weight = age.GetRandomWeight(character.Gender.Name, character.AgeCategory)

	return character
}
