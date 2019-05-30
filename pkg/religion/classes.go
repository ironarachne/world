package religion

import "github.com/ironarachne/world/pkg/random"

// Class is a type of religion
type Class struct {
	Name            string
	Commonality     int
	LeaderTitle     string
	PantheonMinSize int
	PantheonMaxSize int
	GatheringPlaces []string
}

func getAllClasses() []Class {
	return []Class{
		Class{
			Name:            "monotheistic",
			Commonality:     5,
			LeaderTitle:     "priest",
			PantheonMinSize: 1,
			PantheonMaxSize: 1,
			GatheringPlaces: []string{
				"temple",
				"shrine",
				"church",
			},
		},
		Class{
			Name:            "duotheistic",
			Commonality:     1,
			LeaderTitle:     "priest",
			PantheonMinSize: 2,
			PantheonMaxSize: 2,
			GatheringPlaces: []string{
				"temple",
				"shrine",
				"church",
			},
		},
		Class{
			Name:            "polytheistic",
			Commonality:     12,
			LeaderTitle:     "priest",
			PantheonMinSize: 10,
			PantheonMaxSize: 30,
			GatheringPlaces: []string{
				"temple",
				"shrine",
				"church",
			},
		},
		Class{
			Name:            "shamanistic",
			Commonality:     2,
			LeaderTitle:     "shaman",
			PantheonMinSize: 0,
			PantheonMaxSize: 0,
			GatheringPlaces: []string{
				"medicine lodge",
				"sweat lodge",
				"spirit lodge",
			},
		},
		Class{
			Name:            "philosophical",
			Commonality:     1,
			LeaderTitle:     "sage",
			PantheonMinSize: 0,
			PantheonMaxSize: 0,
			GatheringPlaces: []string{
				"great hall",
				"forum",
				"public house",
				"town square",
			},
		},
	}
}

func getWeightedClass() Class {
	classes := getAllClasses()

	weights := map[string]int{}

	for _, c := range classes {
		weights[c.Name] = c.Commonality
	}

	name := random.StringFromThresholdMap(weights)

	for _, c := range classes {
		if c.Name == name {
			return c
		}
	}

	return Class{}
}
