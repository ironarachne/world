/*
Package monster provides monster implementations of the species.Species struct
*/
package monster

import (
	"fmt"

	"github.com/ironarachne/world/pkg/species"
)

// All returns all predefined monsters
func All() ([]species.Species, error) {
	monsters, err := species.Load("monsters")
	if err != nil {
		err = fmt.Errorf("failed to load monsters: %w", err)
		return []species.Species{}, err
	}

	return monsters, nil
}
