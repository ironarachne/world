/*
Package plant provides plant implementation of species.Species
*/
package plant

import (
	"fmt"

	"github.com/ironarachne/world/pkg/species"
)

// All returns all predefined plants
func All() ([]species.Species, error) {
	plants, err := species.Load("plants")
	if err != nil {
		err = fmt.Errorf("failed to load plants: %w", err)
		return []species.Species{}, err
	}

	return plants, nil
}
