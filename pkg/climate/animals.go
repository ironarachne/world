package climate

import (
	"fmt"
	"math/rand"

	"github.com/ironarachne/world/pkg/animal"
	"github.com/ironarachne/world/pkg/fish"
	"github.com/ironarachne/world/pkg/insect"
	"github.com/ironarachne/world/pkg/species"
)

func (gen Generator) getAnimals(humidity int, temperature int) ([]species.Species, error) {
	animals, err := getFilteredAnimals(humidity, temperature)
	if err != nil {
		err = fmt.Errorf("failed to get animals for climate: %w", err)
		return []species.Species{}, err
	}

	hideAnimals := species.ByTag("hide", animals)
	hideAnimal := species.Random(1, hideAnimals)

	animals = species.ByTagIn(gen.AnimalTags, animals)

	numberOfAnimals := rand.Intn(gen.AnimalMax-gen.AnimalMin) + gen.AnimalMin - 1

	animals = species.Random(numberOfAnimals, animals)
	animals = append(animals, hideAnimal...)

	return animals, nil
}

func getFilteredAnimals(humidity int, temperature int) ([]species.Species, error) {
	animals, err := animal.All()
	if err != nil {
		err = fmt.Errorf("failed to get filtered animals for climate: %w", err)
		return []species.Species{}, err
	}
	fishes, err := fish.All()
	if err != nil {
		err = fmt.Errorf("failed to get filtered fish for climate: %w", err)
		return []species.Species{}, err
	}
	animals = append(animals, fishes...)
	insects, err := insect.All()
	if err != nil {
		err = fmt.Errorf("failed to get filtered insect for climate: %w", err)
		return []species.Species{}, err
	}
	animals = append(animals, insects...)
	animals = species.FilterHumidity(humidity, animals)
	animals = species.FilterTemperature(temperature, animals)

	return animals, nil
}
