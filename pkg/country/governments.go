package country

import (
	"fmt"
	"math/rand"

	"github.com/ironarachne/world/pkg/profession"

	"github.com/ironarachne/world/pkg/character"
	"github.com/ironarachne/world/pkg/heraldry"
)

// Government is a political structure
type Government struct {
	Type    string
	Leaders []character.Character
}

func (country Country) getNewMonarchy() (Government, error) {
	government := Government{}
	government.Type = "monarchy"

	monarch, err := character.Generate(country.DominantCulture)
	if err != nil {
		err = fmt.Errorf("Could not generate monarch: %w", err)
		return Government{}, err
	}
	monarch.ChangeAge(rand.Intn(30) + 20)

	if monarch.Gender.Name == "male" {
		monarch.Title = "King"
		monarch.Profession = profession.ByName("noble")
	} else {
		monarch.Title = "Queen"
		monarch.Profession = profession.ByName("noble")
	}

	device, err := heraldry.Generate()
	if err != nil {
		err = fmt.Errorf("Could not generate monarch: %w", err)
		return Government{}, err
	}
	monarch.Heraldry = device

	government.Leaders = append(government.Leaders, monarch)

	return government, nil
}
