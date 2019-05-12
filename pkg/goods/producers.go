package goods

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/climate"
	"github.com/ironarachne/world/pkg/slices"
)

// Producer is a type of person who creates a trade good
type Producer struct {
	Name             string
	MinCommunitySize int
	Patterns         []Pattern
	SkillLevel       int
}

func getAllProducers() []Producer {
	armor := getArmor()
	weapons := getWeapons()

	blacksmithPatterns := append(armor, weapons...)

	producers := []Producer{
		Producer{
			Name:             "alchemist",
			MinCommunitySize: 500,
			Patterns:         getPotions(),
		},
		Producer{
			Name:             "animal trainer",
			MinCommunitySize: 10,
			Patterns:         getAnimalTrainerGoods(),
		},
		Producer{
			Name:             "apothecary",
			MinCommunitySize: 100,
			Patterns:         getMedicines(),
		},
		Producer{
			Name:             "armorsmith",
			MinCommunitySize: 500,
			Patterns:         getArmor(),
		},
		Producer{
			Name:             "baker",
			MinCommunitySize: 10,
			Patterns:         getBreads(),
		},
		Producer{
			Name:             "blacksmith",
			MinCommunitySize: 20,
			Patterns:         blacksmithPatterns,
		},
		Producer{
			Name:             "brewer",
			MinCommunitySize: 10,
			Patterns:         getBrewed(),
		},
		Producer{
			Name:             "carpenter",
			MinCommunitySize: 10,
			Patterns:         getCarpenterGoods(),
		},
		Producer{
			Name:             "cobbler",
			MinCommunitySize: 10,
			Patterns:         getCobblerGoods(),
		},
		Producer{
			Name:             "mason",
			MinCommunitySize: 100,
			Patterns:         getMasonGoods(),
		},
		Producer{
			Name:             "miner",
			MinCommunitySize: 10,
			Patterns:         getMinerGoods(),
		},
		Producer{
			Name:             "potter",
			MinCommunitySize: 10,
			Patterns:         getPottery(),
		},
		Producer{
			Name:             "tailor",
			MinCommunitySize: 10,
			Patterns:         getClothing(),
		},
		Producer{
			Name:             "tanner",
			MinCommunitySize: 10,
			Patterns:         getTannerGoods(),
		},
		Producer{
			Name:             "vintner",
			MinCommunitySize: 100,
			Patterns:         getWine(),
		},
		Producer{
			Name:             "weaponsmith",
			MinCommunitySize: 500,
			Patterns:         getWeapons(),
		},
	}

	return producers
}

// GetPossibleProducers gets all possible producers for a given set of resources
func GetPossibleProducers(resources []climate.Resource, communitySize int) []Producer {
	possibleProducers := getAllProducers()

	availableTypes := []string{}
	producers := []Producer{}

	patterns := []Pattern{}

	need1Met := false
	need2Met := false
	need3Met := false

	for _, r := range resources {
		if !slices.StringIn(r.Type, availableTypes) {
			availableTypes = append(availableTypes, r.Type)
		}
	}

	for _, p := range possibleProducers {
		if p.MinCommunitySize <= communitySize {
			patterns = []Pattern{}
			for _, i := range p.Patterns {
				need1Met = slices.StringIn(i.Need1, availableTypes)
				need2Met = i.Need2 == "" || i.Need2 != "" && slices.StringIn(i.Need2, availableTypes)
				need3Met = i.Need3 == "" || i.Need3 != "" && slices.StringIn(i.Need3, availableTypes)

				if need1Met && need2Met && need3Met {
					patterns = append(patterns, i)
				}
			}
			p.Patterns = patterns
			if len(patterns) > 0 {
				producers = append(producers, p)
			}
		}
	}

	return producers
}

// GetRandomProducers returns a random number of producers from a given set of producers
func GetRandomProducers(max int, possible []Producer) []Producer {
	producers := []Producer{}
	producer := Producer{}

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
