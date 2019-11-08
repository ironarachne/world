package climate

import (
	"github.com/ironarachne/world/pkg/soil"
)

func (gen Generator) getSoils() []soil.Soil {
	soils := soil.All()

	soils = soil.ByTagIn(gen.SoilTags, soils)

	filtered := soil.Random(2, soils)

	return filtered
}
