package country

import (
	"github.com/ironarachne/world/pkg/profession"
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

	monarch := character.Generate(country.DominantCulture)
	monarch.ChangeAge(rand.Intn(30) + 20)

	if monarch.Gender.Name == "male" {
		monarch.Title = "King"
		monarch.Profession = profession.ByName("noble")
	} else {
		monarch.Title = "Queen"
		monarch.Profession = profession.ByName("noble")
	}

	monarch.Heraldry = heraldry.GenerateHeraldry()

	government.Leaders = append(government.Leaders, monarch)

	return government
}
