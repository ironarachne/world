package region

import "github.com/ironarachne/world/pkg/random"

// Class is a class of region
type Class struct {
	MaxNumberOfTowns int
	MinNumberOfTowns int
	Name             string
	RulerTitleFemale string
	RulerTitleMale   string
	Commonality      int
}

func getAllClasses() []Class {
	classes := []Class{
		Class{
			Name:             "Barony",
			RulerTitleFemale: "Baroness",
			RulerTitleMale:   "Baron",
			MinNumberOfTowns: 2,
			MaxNumberOfTowns: 4,
			Commonality:      3,
		},
		Class{
			Name:             "County",
			RulerTitleFemale: "Countess",
			RulerTitleMale:   "Count",
			MinNumberOfTowns: 1,
			MaxNumberOfTowns: 3,
			Commonality:      4,
		},
		Class{
			Name:             "Duchy",
			RulerTitleFemale: "Duchess",
			RulerTitleMale:   "Duke",
			MinNumberOfTowns: 3,
			MaxNumberOfTowns: 6,
			Commonality:      2,
		},
		Class{
			Name:             "March",
			RulerTitleFemale: "Marchioness",
			RulerTitleMale:   "Marquess",
			MinNumberOfTowns: 2,
			MaxNumberOfTowns: 4,
			Commonality:      2,
		},
		Class{
			Name:             "Viscounty",
			RulerTitleFemale: "Viscountess",
			RulerTitleMale:   "Viscount",
			MinNumberOfTowns: 1,
			MaxNumberOfTowns: 2,
			Commonality:      1,
		},
	}

	return classes
}

func getRandomWeightedClass() Class {
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
