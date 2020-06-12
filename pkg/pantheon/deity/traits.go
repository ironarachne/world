package deity

import (
	"context"
	"fmt"

	"github.com/ironarachne/world/pkg/pantheon/domain"
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

func (deity Deity) getRandomTraits(ctx context.Context) ([]string, error) {
	var possibleTraits []string

	allTraits := getAllTraits()

	for i := 0; i < 4; i++ {
		trait, err := random.String(ctx, allTraits)
		if err != nil {
			err = fmt.Errorf("Could not generate deity traits: %w", err)
			return []string{}, err
		}
		if !slices.StringIn(trait, possibleTraits) {
			possibleTraits = append(possibleTraits, trait)
		}
	}

	domainTraits := domain.AllPersonalitiesForDomains(deity.Domains)
	possibleTraits = append(possibleTraits, domainTraits...)
	traits := []string{}

	for i := 0; i < 2; i++ {
		// Only add a trait if it isn't already in the PersonalityTraits slice
		trait, err := random.String(ctx, possibleTraits)
		if err != nil {
			err = fmt.Errorf("Could not generate deity traits: %w", err)
			return []string{}, err
		}
		if !slices.StringIn(trait, traits) {
			traits = append(traits, trait)
		}
	}

	return traits, nil
}
