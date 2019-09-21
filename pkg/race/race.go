package race

import (
	"fmt"
	"math/rand"

	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/species"
)

// All returns all races
func All() []species.Species {
	races := []species.Species{}

	races = append(races, getDwarves()...)
	races = append(races, getElves()...)
	races = append(races, getHalflings()...)
	races = append(races, getHumans()...)
	races = append(races, getGnomes()...)

	return races
}

// ByName returns a specific race by name
func ByName(name string) species.Species {
	races := All()

	for _, r := range races {
		if r.Name == name {
			return r
		}
	}

	return species.Species{}
}

// Random returns a random race from the list
func Random() (species.Species, error) {
	races := All()

	if len(races) == 0 {
		err := fmt.Errorf("tried to get random race from slice of zero races")
		return species.Species{}, err
	}

	if len(races) == 1 {
		return races[0], nil
	}

	race := races[rand.Intn(len(races))]

	return race, nil
}

// RandomSimplified returns a random simplified race
func RandomSimplified() (species.Simplified, error) {
	race, err := Random()
	if err != nil {
		err = fmt.Errorf("Failed to generate random simplified race: %w", err)
		return species.Simplified{}, err
	}

	sr := race.Simplify()

	return sr, nil
}

// RandomWeighted returns a random race, taking commonality into account
func RandomWeighted() (species.Species, error) {
	races := All()

	weights := map[string]int{}

	for _, c := range races {
		weights[c.Name] = c.Commonality
	}

	name, err := random.StringFromThresholdMap(weights)
	if err != nil {
		err = fmt.Errorf("Failed to get random weighted race: %w", err)
		return species.Species{}, err
	}

	for _, c := range races {
		if c.Name == name {
			return c, nil
		}
	}

	err = fmt.Errorf("failed to get random weighted race")
	return species.Species{}, err
}
