package climate

import (
	"github.com/ironarachne/world/pkg/insect"
	"github.com/ironarachne/world/pkg/species"
)

func (climate Climate) getFilteredInsects() []species.Species {
	insects := insect.All()
	insects = species.FilterHumidity(climate.Humidity, insects)
	insects = species.FilterTemperature(climate.Temperature, insects)

	return insects
}
