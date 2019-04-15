package pantheon

import "github.com/ironarachne/random"

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
	traits := []string{}

	for i := 0; i < 2; i++ {
		// Only add a trait if it isn't already in the PersonalityTraits slice
		trait := random.Item(possibleTraits)
		if !inSlice(trait, traits) {
			traits = append(traits, trait)
		}
	}

	return traits
}
