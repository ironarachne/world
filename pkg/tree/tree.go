/*
Package tree implements trees
*/
package tree

import (
	"fmt"

	"github.com/ironarachne/world/pkg/species"
)

// All returns all predefined trees
func All() ([]species.Species, error) {
	trees, err := species.Load("plants/trees")
	if err != nil {
		err = fmt.Errorf("could not load trees: %w", err)
		return []species.Species{}, err
	}

	return trees, nil
}
