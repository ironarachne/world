/*
Package insect provides insect data sets and logic. It makes use of
the species package to provide an underlying structure and common
elements.
*/
package insect

import (
	"github.com/ironarachne/world/pkg/species"
)

// All returns all insects
func All() []species.Species {
	arachnids := getArachnids()
	beetles := getBeetles()

	insects := getInsects()

	insects = append(insects, arachnids...)
	insects = append(insects, beetles...)

	return insects
}
