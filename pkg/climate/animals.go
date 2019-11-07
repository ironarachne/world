package climate

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/animal"
	"github.com/ironarachne/world/pkg/species"
)

func (gen Generator) getAnimals(humidity int, temperature int) ([]species.Species, error) {
	animals := getFilteredAnimals(humidity, temperature)

	hideAnimals := species.ByTag("hide", animals)
	hideAnimal := species.Random(1, hideAnimals)

	animals = species.ByTagIn(gen.AnimalTags, animals)

	numberOfAnimals := rand.Intn(gen.AnimalMax-gen.AnimalMin) + gen.AnimalMin - 1

	animals = species.Random(numberOfAnimals, animals)
	animals = append(animals, hideAnimal...)

	return animals, nil
}

func getFilteredAnimals(humidity int, temperature int) []species.Species {
	animals := animal.All()
	animals = species.FilterHumidity(humidity, animals)
	animals = species.FilterTemperature(temperature, animals)

	return animals
}
