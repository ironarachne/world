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
