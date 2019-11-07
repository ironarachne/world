package climate

import (
	"fmt"

	"github.com/ironarachne/world/pkg/plant"
	"github.com/ironarachne/world/pkg/species"
)

func (climate Climate) getPlants() ([]species.Species, error) {
	allPlants := plant.All()
	plants := climate.getFilteredPlants()

	plants = species.Random(climate.MaxPlants-1, plants)

	randomFabricFiber, err := species.RandomWithResourceTag("fabric fiber", allPlants)
	if err != nil {
		err = fmt.Errorf("Could not populate climate: %w", err)
		return []species.Species{}, err
	}

	randomGrain, err := species.RandomWithResourceTag("grain", allPlants)
	if err != nil {
		err = fmt.Errorf("Could not populate climate: %w", err)
		return []species.Species{}, err
	}

	randomFruit, err := species.RandomWithResourceTag("fruit", allPlants)
	if err != nil {
		err = fmt.Errorf("Could not populate climate: %w", err)
		return []species.Species{}, err
	}

	plants = append(plants, randomFabricFiber)
	plants = append(plants, randomGrain)
	plants = append(plants, randomFruit)

	return plants, nil
}

func (climate Climate) getFilteredPlants() []species.Species {
	plants := plant.All()
	plants = species.FilterHumidity(climate.Humidity, plants)
	plants = species.FilterTemperature(climate.Temperature, plants)

	if len(plants) < 1 {
		panic("Found no plants for climate " + climate.Name)
	}

	return plants
}
