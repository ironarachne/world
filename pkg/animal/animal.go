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
	var animals []species.Species
	var err error
	var loaded []species.Species

	fileNames := []string{
		"bigcats",
		"canids",
		"cervids",
		"equines",
		"flightlessbirds",
		"gamebirds",
		"miscellaneous",
		"pigs",
		"primates",
		"raptors",
		"reptiles",
		"rodents",
		"waterfowl",
	}

	for _, f := range fileNames {
		loaded, err = species.Load("animals/" + f)
		if err != nil {
			err = fmt.Errorf("failed to load animals: %w", err)
			return []species.Species{}, err
		}
		animals = append(animals, loaded...)
	}

	return animals, nil
}
