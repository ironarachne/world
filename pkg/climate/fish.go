package climate

import (
	"github.com/ironarachne/world/pkg/fish"
	"github.com/ironarachne/world/pkg/species"
)

func (climate Climate) getFish() []species.Species {
	var result []species.Species
	allFish := fish.All()
	potentialFish := []species.Species{}

	if !climate.HasLakes {
		for _, f := range allFish {
			if f.HasTag("lives in lakes") && !f.InSlice(potentialFish) {
				potentialFish = append(potentialFish, f)
			}
		}
	}
	if !climate.HasOcean {
		for _, f := range allFish {
			if f.HasTag("lives in oceans") && !f.InSlice(potentialFish) {
				potentialFish = append(potentialFish, f)
			}
		}
	}
	if !climate.HasRivers {
		for _, f := range allFish {
			if f.HasTag("lives in rivers") && !f.InSlice(potentialFish) {
				potentialFish = append(potentialFish, f)
			}
		}
	}

	if len(potentialFish) > 1 {
		result = species.Random(climate.MaxFish, potentialFish)
	}

	return result
}
