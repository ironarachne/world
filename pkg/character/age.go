package character

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/math"
	"github.com/ironarachne/world/pkg/race"
)

func getAgeFromParents(parents Couple) (int, race.AgeCategory) {
	lowestAge := math.Min(parents.Partner1.Age, parents.Partner2.Age)

	childAge := rand.Intn(lowestAge-18) + 1
	childAgeCategory := parents.Partner1.Race.GetAgeCategoryFromAge(childAge)

	return childAge, childAgeCategory
}

// ChangeAge changes the age and age category of a character
func (character Character) ChangeAge(newAge int) Character {
	c := character.Race.GetAgeCategoryFromAge(newAge)
	character.Age = newAge
	character.AgeCategory = c
	character.Height = character.randomHeight()
	character.Weight = character.randomWeight()

	return character
}
