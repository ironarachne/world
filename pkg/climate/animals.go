package climate

import (
	"github.com/ironarachne/world/pkg/animal"
	"github.com/ironarachne/world/pkg/species"
)

func (climate Climate) getAnimals() ([]species.Species, error) {
	animals := climate.getFilteredAnimals()

	hideAnimals := species.ByTag("hide", animals)
	hideAnimal := species.Random(1, hideAnimals)

	animals = species.Random(climate.MaxAnimals, animals)
	animals = append(animals, hideAnimal...)

	return animals, nil
}

func (climate Climate) getFilteredAnimals() []species.Species {
	animals := animal.All()
	animals = species.FilterHumidity(climate.Humidity, animals)
	animals = species.FilterTemperature(climate.Temperature, animals)
	animals = climate.filterAnimalsForWater(animals)

	return animals
}

func (climate Climate) filterAnimalsForWater(animals []species.Species) []species.Species {
	var filteredAnimals []species.Species

	for _, a := range animals {
		if a.HasTag("aquatic") && (climate.HasLakes || climate.HasOcean || climate.HasRivers || climate.HasWetlands) {
			filteredAnimals = append(filteredAnimals, a)
		} else if !a.HasTag("aquatic") {
			filteredAnimals = append(filteredAnimals, a)
		}
	}

	return filteredAnimals
}
