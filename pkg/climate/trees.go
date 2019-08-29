package climate

import (
	"github.com/ironarachne/world/pkg/tree"
)

func (climate Climate) getFilteredTrees() []tree.Tree {
	trees := tree.All()
	trees = filterTreesForHumidity(climate.Humidity, trees)
	trees = filterTreesForTemperature(climate.Temperature, trees)
	trees = filterTreesForType(climate, trees)

	return trees
}

func filterTreesForType(climate Climate, trees []tree.Tree) []tree.Tree {
	var filteredTrees []tree.Tree

	for _, a := range trees {
		if (climate.HasDeciduousTrees && a.IsDeciduous) || (climate.HasConiferousTrees && a.IsConiferous) {
			filteredTrees = append(filteredTrees, a)
		}
	}

	return filteredTrees
}

func filterTreesForHumidity(humidity int, trees []tree.Tree) []tree.Tree {
	var filteredTrees []tree.Tree

	for _, a := range trees {
		if a.MinHumidity <= humidity && a.MaxHumidity >= humidity {
			filteredTrees = append(filteredTrees, a)
		}
	}

	return filteredTrees
}

func filterTreesForTemperature(temperature int, trees []tree.Tree) []tree.Tree {
	var filteredTrees []tree.Tree

	for _, a := range trees {
		if a.MinTemperature <= temperature && a.MaxTemperature >= temperature {
			filteredTrees = append(filteredTrees, a)
		}
	}

	return filteredTrees
}
