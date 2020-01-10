package climate

import (
	"fmt"

	"github.com/ironarachne/world/pkg/species"
	"github.com/ironarachne/world/pkg/tree"
)

func (gen Generator) getTrees(humidity int, temperature int) ([]species.Species, error) {
	trees, err := tree.All()
	if err != nil {
		err = fmt.Errorf("failed to get trees for climate: %w", err)
		return []species.Species{}, err
	}
	trees = species.FilterHumidity(humidity, trees)
	trees = species.FilterTemperature(temperature, trees)
	trees = species.ByTagIn(gen.TreeTags, trees)

	if len(trees) > 8 {
		trees = species.Random(8, trees)
	}

	return trees, nil
}
