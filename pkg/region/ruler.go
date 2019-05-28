package region

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/character"
	"github.com/ironarachne/world/pkg/heraldry"
)

func (region Region) generateRuler() character.Character {
	ruler := character.Generate(region.Culture)
	ruler.Profession = "noble"
	ruler = ruler.ChangeAge(rand.Intn(40) + 25)

	ruler.Title = region.Class.RulerTitleFemale
	if ruler.Gender.Name == "male" {
		ruler.Title = region.Class.RulerTitleMale
	}

	ruler.Heraldry = heraldry.GenerateHeraldry()

	return ruler
}
