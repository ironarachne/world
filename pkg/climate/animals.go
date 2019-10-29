package climate

import (
	"github.com/ironarachne/world/pkg/animal"
)

func (climate Climate) getAnimals() ([]animal.Species, error) {
	animals := climate.getFilteredAnimals()

	hideAnimals := animal.ByTag("hide", animals)
	hideAnimal := animal.Random(1, hideAnimals)

	animals = animal.Random(climate.MaxAnimals, animals)
	animals = append(animals, hideAnimal...)

	return animals, nil
}

func (climate Climate) getFilteredAnimals() []animal.Species {
	animals := animal.All()
	animals = filterAnimalsForHumidity(climate.Humidity, animals)
	animals = filterAnimalsForTemperature(climate.Temperature, animals)
	animals = climate.filterAnimalsForWater(animals)

	return animals
}

func filterAnimalsForHumidity(humidity int, animals []animal.Species) []animal.Species {
	var filteredAnimals []animal.Species

	for _, a := range animals {
		if a.MinHumidity <= humidity && a.MaxHumidity >= humidity {
			filteredAnimals = append(filteredAnimals, a)
		}
	}

	return filteredAnimals
}

func filterAnimalsForTemperature(temperature int, animals []animal.Species) []animal.Species {
	var filteredAnimals []animal.Species

	for _, a := range animals {
		if a.MinTemperature <= temperature && a.MaxTemperature >= temperature {
			filteredAnimals = append(filteredAnimals, a)
		}
	}

	return filteredAnimals
}

func (climate Climate) filterAnimalsForWater(animals []animal.Species) []animal.Species {
	var filteredAnimals []animal.Species

	for _, a := range animals {
		if a.IsAquatic && (climate.HasLakes || climate.HasOcean || climate.HasRivers || climate.HasWetlands) {
			filteredAnimals = append(filteredAnimals, a)
		} else if !a.IsAquatic {
			filteredAnimals = append(filteredAnimals, a)
		}
	}

	return filteredAnimals
}
