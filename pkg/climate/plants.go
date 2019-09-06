package climate

import (
	"fmt"

	"github.com/ironarachne/world/pkg/plant"
)

func (climate Climate) getPlants() ([]plant.Plant, error) {
	plants := climate.getFilteredPlants()

	plants = plant.Random(climate.MaxPlants-1, plants)

	randomFabricFiber, err := plant.RandomPlantWithResource("fabric fiber")
	if err != nil {
		err = fmt.Errorf("Could not populate climate: %w", err)
		return []plant.Plant{}, err
	}

	randomGrain, err := plant.RandomPlantWithResource("grain")
	if err != nil {
		err = fmt.Errorf("Could not populate climate: %w", err)
		return []plant.Plant{}, err
	}

	randomFruit, err := plant.RandomPlantWithResource("fruit")
	if err != nil {
		err = fmt.Errorf("Could not populate climate: %w", err)
		return []plant.Plant{}, err
	}

	plants = append(plants, randomFabricFiber)
	plants = append(plants, randomGrain)
	plants = append(plants, randomFruit)

	return plants, nil
}

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
