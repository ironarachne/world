package culture

import (
	"github.com/ironarachne/world/pkg/pantheon"
	"github.com/ironarachne/world/pkg/random"
)

// Religion is a culture's religion
type Religion struct {
	Class              ReligionClass
	Pantheon           pantheon.SimplifiedPantheon
	GatheringPlaceName string
}

// ReligionClass is a type of religion
type ReligionClass struct {
	Name                string
	LeaderTitle         string
	PantheonMaxSize     int
	GatheringPlaceNames []string
}

func (culture Culture) generateReligion() Religion {
	religion := Religion{}

	religion.Class = randomReligionClass()
	religion.GatheringPlaceName = religion.randomGatheringPlaceName()

	if religion.Class.PantheonMaxSize > 0 {
		religion.Pantheon = pantheon.GenerateForDisplay(religion.Class.PantheonMaxSize, culture.Language)
	}

	return religion
}

func (religion Religion) randomGatheringPlaceName() string {
	return random.String(religion.Class.GatheringPlaceNames)
}

func randomReligionClass() ReligionClass {
	weightedClasses := map[string]int{
		"monotheistic":  5,
		"philosophical": 1,
		"polytheistic":  9,
		"shamanistic":   2,
	}

	classes := map[string]ReligionClass{
		"monotheistic": ReligionClass{
			Name:            "monotheistic",
			LeaderTitle:     "priest",
			PantheonMaxSize: 1,
			GatheringPlaceNames: []string{
				"temple",
				"shrine",
				"church",
			},
		},
		"polytheistic": ReligionClass{
			Name:            "polytheistic",
			LeaderTitle:     "priest",
			PantheonMaxSize: 30,
			GatheringPlaceNames: []string{
				"temple",
				"shrine",
				"church",
			},
		},
		"shamanistic": ReligionClass{
			Name:            "shamanistic",
			LeaderTitle:     "shaman",
			PantheonMaxSize: 0,
			GatheringPlaceNames: []string{
				"medicine lodge",
				"sweat lodge",
				"spirit lodge",
			},
		},
		"philosophical": ReligionClass{
			Name:            "philosophical",
			LeaderTitle:     "sage",
			PantheonMaxSize: 0,
			GatheringPlaceNames: []string{
				"great hall",
				"forum",
				"public house",
				"town square",
			},
		},
	}

	className := random.StringFromThresholdMap(weightedClasses)

	class := classes[className]

	return class
}
