package climate

import (
	"github.com/ironarachne/world/pkg/fish"
)

func (climate Climate) getFish() []fish.Fish {
	var result []fish.Fish
	allFish := fish.All()
	potentialFish := []fish.Fish{}

	if !climate.HasLakes {
		for _, f := range allFish {
			if f.LivesInLakes && !fish.InSlice(f, potentialFish) {
				potentialFish = append(potentialFish, f)
			}
		}
	}
	if !climate.HasOcean {
		for _, f := range allFish {
			if f.LivesInOceans && !fish.InSlice(f, potentialFish) {
				potentialFish = append(potentialFish, f)
			}
		}
	}
	if !climate.HasRivers {
		for _, f := range allFish {
			if f.LivesInRivers && !fish.InSlice(f, potentialFish) {
				potentialFish = append(potentialFish, f)
			}
		}
	}

	if len(potentialFish) > 1 {
		result = fish.Random(climate.MaxFish, potentialFish)
	}

	return result
}
