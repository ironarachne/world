/*
Package fish provides fish data sets and logic. It makes use of
the species package to provide an underlying structure and common
elements.
*/
package fish

import (
	"fmt"

	"github.com/ironarachne/world/pkg/species"
)

// All returns all predefined fish
func All() ([]species.Species, error) {
	fish, err := species.Load("fish")
	if err != nil {
		err = fmt.Errorf("failed to load fish: %w", err)
		return []species.Species{}, err
	}

	return fish, nil
}
