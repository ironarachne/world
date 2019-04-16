package region

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/character"
)

func (region Region) generateRuler() character.Character {
	ruler := character.GenerateCharacterOfCulture(region.Culture)
	ruler.Profession = "noble"
	ruler = ruler.ChangeAge(rand.Intn(40) + 25)

	return ruler
}
