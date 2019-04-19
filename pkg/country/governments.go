package country

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/character"
	"github.com/ironarachne/world/pkg/heraldry"
)

// Government is a political structure
type Government struct {
	Type    string
	Leaders []character.Character
}

func (country Country) getNewMonarchy() Government {
	government := Government{}
	government.Type = "monarchy"

	monarch := character.GenerateCharacterOfCulture(country.DominantCulture)
	monarch.ChangeAge(rand.Intn(30) + 20)

	if monarch.Gender == "male" {
		monarch.Title = "King"
		monarch.Profession = "monarch"
	} else {
		monarch.Title = "Queen"
		monarch.Profession = "monarch"
	}

	monarch.Heraldry = heraldry.GenerateHeraldry()

	government.Leaders = append(government.Leaders, monarch)

	return government
}
