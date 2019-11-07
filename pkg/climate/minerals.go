package climate

import (
	"fmt"

	"github.com/ironarachne/world/pkg/mineral"
)

func (gen Generator) getMinerals() ([]mineral.Mineral, error) {
	var minerals []mineral.Mineral

	metals := mineral.Metals()

	filteredMetals, err := mineral.RandomWeightedSet(5, metals)
	if err != nil {
		err = fmt.Errorf("could not get minerals: %w", err)
		return []mineral.Mineral{}, err
	}

	gems := mineral.Gems()
	filteredGems := mineral.Random(3, gems)

	stones := mineral.Stones()
	filteredStones := mineral.Random(2, stones)
	other := mineral.OtherMinerals()

	minerals = append(minerals, filteredMetals...)
	minerals = append(minerals, filteredGems...)
	minerals = append(minerals, filteredStones...)
	minerals = append(minerals, other...)

	return minerals, nil
}
