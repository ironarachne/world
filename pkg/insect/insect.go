package insect

import (
	"math/rand"

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

// RandomSubset returns a set number of randomly chosen insects from a slice
func RandomSubset(amount int, from []species.Species) []species.Species {
	var insect species.Species

	insects := []species.Species{}

	if amount > len(from) {
		amount = len(from)
	}

	for i := 0; i < amount; i++ {
		insect = from[rand.Intn(len(from))]
		if !insect.InSlice(insects) {
			insects = append(insects, insect)
		}
	}

	return insects
}
