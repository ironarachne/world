package climate

import (
	"github.com/ironarachne/world/pkg/plant"
)

func (climate Climate) getFilteredPlants() []plant.Plant {
	plants := plant.All()
	plants = filterPlantsForHumidity(climate.Humidity, plants)
	plants = filterPlantsForTemperature(climate.Temperature, plants)

	if len(plants) < 1 {
		panic("Found no plants for climate " + climate.Name)
	}

	return plants
}

func filterPlantsForHumidity(humidity int, plants []plant.Plant) []plant.Plant {
	var filteredPlants []plant.Plant

	for _, a := range plants {
		if a.MinHumidity <= humidity && a.MaxHumidity >= humidity {
			filteredPlants = append(filteredPlants, a)
		}
	}

	return filteredPlants
}

func filterPlantsForTemperature(temperature int, plants []plant.Plant) []plant.Plant {
	var filteredPlants []plant.Plant

	for _, a := range plants {
		if a.MinTemperature <= temperature && a.MaxTemperature >= temperature {
			filteredPlants = append(filteredPlants, a)
		}
	}

	return filteredPlants
}
