package character

import (
	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/slices"
)

func getAllTraits() []string {
	traits := []string{
		"addict",
		"anxious",
		"arrogant",
		"brave",
		"bumbling",
		"cautious",
		"confident",
		"contemptuous",
		"devious",
		"drunkard",
		"encouraging",
		"envious",
		"fair",
		"fierce",
		"just",
		"lusty",
		"meticulous",
		"optimistic",
		"persistent",
		"pessimistic",
		"proud",
		"reckless",
		"religious",
		"resilient",
		"sickly",
		"stable",
		"strong",
		"stubborn",
		"trusting",
		"wise",
	}

	return traits
}

func getRandomTrait() string {
	traits := getAllTraits()

	return random.String(traits)
}

func getRandomTraits() []string {
	possibleTraits := getAllTraits()

	traits := []string{}

	for i := 0; i < 2; i++ {
		trait := random.String(possibleTraits)
		for slices.StringIn(trait, traits) {
			trait = random.String(possibleTraits)
		}
		traits = append(traits, trait)
	}

	return traits
}
