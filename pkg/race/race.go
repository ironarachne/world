/*
Package race provides fantasy races as an implementation of species.Species
*/
package race

import (
	"fmt"
	"math/rand"

	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/species"
)

// All returns all predefined races
func All() ([]species.Species, error) {
	races, err := species.Load("races")
	if err != nil {
		err = fmt.Errorf("failed to load races: %w", err)
		return []species.Species{}, err
	}

	return races, nil
}

// ByName returns a specific race by name
func ByName(name string) (species.Species, error) {
	races, err := All()
	if err != nil {
		err = fmt.Errorf("failed to find race by name: %w", err)
		return species.Species{}, err
	}

	for _, r := range races {
		if r.Name == name {
			return r, nil
		}
	}

	return species.Species{}, nil
}

// Random returns a random race from the list
func Random() (species.Species, error) {
	races, err := All()
	if err != nil {
		err = fmt.Errorf("failed to find random race: %w", err)
		return species.Species{}, err
	}

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
	races, err := All()
	if err != nil {
		err = fmt.Errorf("failed to find random weighted race: %w", err)
		return species.Species{}, err
	}

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
