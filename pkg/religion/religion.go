package religion

import (
	"github.com/ironarachne/world/pkg/language"
	"github.com/ironarachne/world/pkg/pantheon"
	"github.com/ironarachne/world/pkg/random"
)

// Religion is a culture's religion
type Religion struct {
	Class          Class
	Pantheon       pantheon.Pantheon
	GatheringPlace string
}

// Generate procedurally generates a religion
func Generate(originLanguage language.Language) Religion {
	religion := Religion{}

	religion.Class = getWeightedClass()
	religion.GatheringPlace = religion.randomGatheringPlace()

	if religion.Class.PantheonMaxSize > 0 {
		religion.Pantheon = pantheon.Generate(religion.Class.PantheonMinSize, religion.Class.PantheonMaxSize, originLanguage)
	}

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
