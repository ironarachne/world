package religion

import (
	"github.com/ironarachne/world/pkg/language"
	"github.com/ironarachne/world/pkg/pantheon"
	"github.com/ironarachne/world/pkg/random"
)

// Religion is a religion
type Religion struct {
	Name           string
	CommonName     string
	Class          Class
	Pantheon       pantheon.Pantheon
	GatheringPlace string
	Virtues        []string
}

// Generate procedurally generates a religion
func Generate(originLanguage language.Language) Religion {
	religion := Religion{}

	religion.Class = getWeightedClass()
	religion.GatheringPlace = religion.randomGatheringPlace()

	if religion.Class.PantheonMaxSize > 0 {
		religion.Pantheon = pantheon.Generate(religion.Class.PantheonMinSize, religion.Class.PantheonMaxSize, originLanguage)
	}

	name := originLanguage.RandomName()
	religion.CommonName = name + "ism"
	name = originLanguage.MakePractice(name)
	religion.Name = name

	virtues := getRandomVirtues()
	religion.Virtues = virtues

	return religion
}

func (religion Religion) randomGatheringPlace() string {
	return random.String(religion.Class.GatheringPlaces)
}

// Random generates a completely random religion
func Random() Religion {
	originLanguage := language.Generate()
	religion := Generate(originLanguage)

	return religion
}
