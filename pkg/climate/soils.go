package climate

import (
	"github.com/ironarachne/world/pkg/soil"
)

func (climate Climate) getFilteredSoils() []soil.Soil {
	soils := []soil.Soil{}
	agSoils := soil.Agricultural()
	clays := soil.Clays()
	sands := soil.Sands()

	if climate.Humidity == 0 && climate.Temperature == 9 {
		soils = append(soils, sands...)
	} else {
		soils = append(soils, agSoils...)

		if climate.HasLakes {
			soils = append(soils, clays...)
		}
		if climate.HasOcean {
			soils = append(soils, sands...)
		}
	}

	return soils
}
