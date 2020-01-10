/*
Package insect provides insect data sets and logic. It makes use of
the species package to provide an underlying structure and common
elements.
*/
package insect

import (
	"fmt"

	"github.com/ironarachne/world/pkg/species"
)

// All returns all predefined insects
func All() ([]species.Species, error) {
	insects, err := species.Load("insects")
	if err != nil {
		err = fmt.Errorf("failed to load insects: %w", err)
		return []species.Species{}, err
	}

	return insects, nil
}
