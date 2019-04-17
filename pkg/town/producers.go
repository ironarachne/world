package town

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/slices"
)

// Producer is a type of person who creates a trade good
type Producer struct {
	Name       string
	GoodsMade  []string
	Needs      []string
	SkillLevel int
}

func getAllProducers() []Producer {
	producers := []Producer{
		Producer{
			Name: "alchemist",
			GoodsMade: []string{
				"potions",
			},
			Needs: []string{
				"herb",
			},
		},
		Producer{
			Name: "apothecary",
			GoodsMade: []string{
				"medicines",
			},
			Needs: []string{
				"medicine",
			},
		},
		Producer{
			Name: "baker",
			GoodsMade: []string{
				"bread",
			},
			Needs: []string{
				"grain",
			},
		},
		Producer{
			Name: "blacksmith",
			GoodsMade: []string{
				"alloys",
				"armor",
				"weapons",
			},
			Needs: []string{
				"metal",
			},
		},
		Producer{
			Name: "brewer",
			GoodsMade: []string{
				"beer",
			},
			Needs: []string{
				"grain",
			},
		},
		Producer{
			Name: "carpenter",
			GoodsMade: []string{
				"furniture",
				"timber",
			},
			Needs: []string{
				"wood",
			},
		},
		Producer{
			Name: "cobbler",
			GoodsMade: []string{
				"boots",
				"shoes",
			},
			Needs: []string{
				"hide",
			},
		},
		Producer{
			Name: "mason",
			GoodsMade: []string{
				"block",
			},
			Needs: []string{
				"stone block",
			},
		},
		Producer{
			Name: "pack animal trainer",
			GoodsMade: []string{
				"pack animal",
			},
			Needs: []string{
				"pack animal",
			},
		},
		Producer{
			Name: "tailor",
			GoodsMade: []string{
				"clothing",
			},
			Needs: []string{
				"fabric",
			},
		},
		Producer{
			Name: "tanner",
			GoodsMade: []string{
				"leather",
			},
			Needs: []string{
				"hide",
			},
		},
		Producer{
			Name: "vintner",
			GoodsMade: []string{
				"wine",
			},
			Needs: []string{
				"fruit",
			},
		},
	}

	return producers
}

func (town Town) getPossibleProducers() []Producer {
	possibleProducers := getAllProducers()

	availableTypes := []string{}
	producers := []Producer{}

	for _, r := range town.Climate.Resources {
		if !slices.StringIn(r.Type, availableTypes) {
			availableTypes = append(availableTypes, r.Type)
		}
	}

	for _, p := range possibleProducers {
		if slices.StringSlicePartlyWithin(p.Needs, availableTypes) {
			producers = append(producers, p)
		}
	}

	return producers
}

func (town Town) getRandomProducers(max int) []Producer {
	producers := []Producer{}
	producer := Producer{}

	possible := town.getPossibleProducers()

	for i := 0; i < max; i++ {
		producer = possible[rand.Intn(len(possible)-1)]
		if !isProducerInSlice(producer, producers) {
			producer.SkillLevel = rand.Intn(10) + 1
			producers = append(producers, producer)
		}
	}

	return producers
}

func isProducerInSlice(producer Producer, slice []Producer) bool {
	names := []string{}

	for _, p := range slice {
		names = append(names, p.Name)
	}

	if slices.StringIn(producer.Name, names) {
		return true
	}

	return false
}

func qualityFromSkillLevel(level int) string {
	qualities := map[int]string{
		1:  "poor",
		2:  "mediocre",
		3:  "average",
		4:  "decent",
		5:  "fine",
		6:  "excellent",
		7:  "exceptional",
		8:  "masterful",
		9:  "wonderful",
		10: "legendary",
	}

	return qualities[level]
}
