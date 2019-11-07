package climate

import (
	"github.com/ironarachne/world/pkg/species"
	"github.com/ironarachne/world/pkg/tree"
)

func (climate Climate) getFilteredTrees() []species.Species {
	trees := tree.All()
	trees = species.FilterHumidity(climate.Humidity, trees)
	trees = species.FilterTemperature(climate.Temperature, trees)
	trees = filterTreesForType(climate, trees)

	return trees
}

func filterTreesForType(climate Climate, trees []species.Species) []species.Species {
	filteredTrees := trees

	if !climate.HasDeciduousTrees {
		filteredTrees = species.ExcludeTag("deciduous", filteredTrees)
	}

	if !climate.HasConiferousTrees {
		filteredTrees = species.ExcludeTag("coniferous", filteredTrees)
	}

	return filteredTrees
}
