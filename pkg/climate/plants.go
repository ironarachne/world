package climate

import (
	"fmt"
	"math/rand"

	"github.com/ironarachne/world/pkg/plant"
	"github.com/ironarachne/world/pkg/species"
)

const plantError = "failed to populate climate with plants: %w"

func (gen Generator) getPlants(humidity int, temperature int) ([]species.Species, error) {
	allPlants, err := plant.All()
	if err != nil {
		err = fmt.Errorf(plantError, err)
		return []species.Species{}, err
	}
	plants, err := getFilteredPlants(humidity, temperature)
	if err != nil {
		err = fmt.Errorf(plantError, err)
		return []species.Species{}, err
	}

	numberOfPlants := rand.Intn(gen.AnimalMax-gen.AnimalMin) + gen.AnimalMin

	plants = species.Random(numberOfPlants, plants)

	randomFabricFiber, err := species.RandomWithResourceTag("fabric fiber", allPlants)
	if err != nil {
		err = fmt.Errorf(plantError, err)
		return []species.Species{}, err
	}

	randomGrain, err := species.RandomWithResourceTag("grain", allPlants)
	if err != nil {
		err = fmt.Errorf(plantError, err)
		return []species.Species{}, err
	}

	randomFruit, err := species.RandomWithResourceTag("fruit", allPlants)
	if err != nil {
		err = fmt.Errorf(plantError, err)
		return []species.Species{}, err
	}

	plants = append(plants, randomFabricFiber)
	plants = append(plants, randomGrain)
	plants = append(plants, randomFruit)

	return plants, nil
}

func getFilteredPlants(humidity int, temperature int) ([]species.Species, error) {
	plants, err := plant.All()
	if err != nil {
		err = fmt.Errorf(plantError, err)
		return []species.Species{}, err
	}
	plants = species.FilterHumidity(humidity, plants)
	plants = species.FilterTemperature(temperature, plants)

	return plants, nil
}
