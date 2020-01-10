package climate

import (
	"fmt"

	"github.com/ironarachne/world/pkg/soil"
)

func (gen Generator) getSoils() ([]soil.Soil, error) {
	soils, err := soil.All()
	if err != nil {
		err = fmt.Errorf("failed to get soils for climate: %w", err)
		return []soil.Soil{}, err
	}

	soils = soil.ByTagIn(gen.SoilTags, soils)

	filtered := soil.Random(2, soils)

	return filtered, nil
}
