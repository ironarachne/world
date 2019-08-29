package climate

import (
	"github.com/ironarachne/world/pkg/insect"
)

func (climate Climate) getFilteredInsects() []insect.Insect {
	insects := insect.All()
	insects = filterInsectsForHumidity(climate.Humidity, insects)
	insects = filterInsectsForTemperature(climate.Temperature, insects)

	return insects
}

func filterInsectsForHumidity(humidity int, insects []insect.Insect) []insect.Insect {
	var filteredInsects []insect.Insect

	for _, a := range insects {
		if a.MinHumidity <= humidity && a.MaxHumidity >= humidity {
			filteredInsects = append(filteredInsects, a)
		}
	}

	return filteredInsects
}

func filterInsectsForTemperature(temperature int, insects []insect.Insect) []insect.Insect {
	var filteredInsects []insect.Insect

	for _, a := range insects {
		if a.MinTemperature <= temperature && a.MaxTemperature >= temperature {
			filteredInsects = append(filteredInsects, a)
		}
	}

	return filteredInsects
}
