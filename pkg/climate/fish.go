package climate

import (
	"github.com/ironarachne/world/pkg/animal"
)

func (climate Climate) getFish() []animal.Fish {
	fish := []animal.Fish{}
	allFish := animal.AllFish()
	potentialFish := []animal.Fish{}

	if !climate.HasLakes {
		for _, f := range allFish {
			if f.LivesInLakes && !f.InSlice(potentialFish) {
				potentialFish = append(potentialFish, f)
			}
		}
	}
	if !climate.HasOcean {
		for _, f := range allFish {
			if f.LivesInOceans && !f.InSlice(potentialFish) {
				potentialFish = append(potentialFish, f)
			}
		}
	}
	if !climate.HasRivers {
		for _, f := range allFish {
			if f.LivesInRivers && !f.InSlice(potentialFish) {
				potentialFish = append(potentialFish, f)
			}
		}
	}

	if len(potentialFish) > 1 {
		fish = animal.RandomFish(climate.MaxFish, potentialFish)
	}

	return fish
}
