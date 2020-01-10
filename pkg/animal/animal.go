/*
Package animal provides animal data sets and logic. It makes use of
the species package to provide an underlying structure and common
elements.
*/
package animal

import (
	"fmt"

	"github.com/ironarachne/world/pkg/species"
)

// All returns all predefined animals
func All() ([]species.Species, error) {
	animals, err := species.Load("animals")
	if err != nil {
		err = fmt.Errorf("failed to load animals: %w", err)
		return []species.Species{}, err
	}

	return animals, nil
}
