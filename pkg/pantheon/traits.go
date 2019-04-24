package pantheon

import (
	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/slices"
)

func getAllTraits() []string {
	traits := []string{
		"absent-minded",
		"belligerent",
		"boisterous",
		"bold",
		"callous",
		"clever",
		"creative",
		"cynical",
		"destructive",
		"fae",
		"fierce",
		"greedy",
		"humble",
		"insidious",
		"jittery",
		"jolly",
		"lecherous",
		"lusty",
		"nervous",
		"peaceful",
		"pleasant",
		"proud",
		"quiet",
		"short-tempered",
		"sly",
		"vengeful",
		"vicious",
		"wise",
	}

	return traits
}

func (deity Deity) getRandomTraits() []string {
	possibleTraits := getAllTraits()
	possibleTraits = append(possibleTraits, getAllPersonalitiesForDomains(deity.Domains)...)
	traits := []string{}

	for i := 0; i < 2; i++ {
		// Only add a trait if it isn't already in the PersonalityTraits slice
		trait := random.String(possibleTraits)
		if !slices.StringIn(trait, traits) {
			traits = append(traits, trait)
		}
	}

	return traits
}
