package region

import (
	"fmt"
	"math/rand"

	"github.com/ironarachne/world/pkg/profession"

	"github.com/ironarachne/world/pkg/character"
	"github.com/ironarachne/world/pkg/heraldry"
)

func (region Region) generateRuler() (character.Character, error) {
	ruler, err := character.Generate(region.Culture)
	if err != nil {
		err = fmt.Errorf("Could not generate region: %w", err)
		return character.Character{}, err
	}
	ruler.Profession = profession.ByName("noble")
	ruler = ruler.ChangeAge(rand.Intn(40) + 25)

	ruler.Title = region.Class.RulerTitleFemale
	if ruler.Gender.Name == "male" {
		ruler.Title = region.Class.RulerTitleMale
	}

	ruler.Heraldry = heraldry.Generate()

	return ruler, nil
}
