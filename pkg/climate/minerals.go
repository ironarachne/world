package climate

import (
	"fmt"

	"github.com/ironarachne/world/pkg/mineral"
)

func (gen Generator) getMinerals() ([]mineral.Mineral, error) {
	var minerals []mineral.Mineral

	all, err := mineral.All()
	if err != nil {
		err = fmt.Errorf("could not get minerals: %w", err)
		return []mineral.Mineral{}, err
	}

	metals := mineral.ByTag("metal", all)

	filteredMetals, err := mineral.RandomWeightedSet(5, metals)
	if err != nil {
		err = fmt.Errorf("could not get minerals: %w", err)
		return []mineral.Mineral{}, err
	}

	gems := mineral.ByTag("gem", all)
	filteredGems := mineral.Random(3, gems)

	stones := mineral.ByTag("stone", all)
	filteredStones := mineral.Random(2, stones)
	salt := mineral.ByTag("salt", all)

	minerals = append(minerals, filteredMetals...)
	minerals = append(minerals, filteredGems...)
	minerals = append(minerals, filteredStones...)
	minerals = append(minerals, salt...)

	return minerals, nil
}
