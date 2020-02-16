package geography

import (
	"fmt"
	"github.com/ironarachne/world/pkg/geography/biome"
	"github.com/ironarachne/world/pkg/geography/climate"
	"github.com/ironarachne/world/pkg/geography/region"
	"github.com/ironarachne/world/pkg/geography/season"
	"github.com/ironarachne/world/pkg/mineral"
	"github.com/ironarachne/world/pkg/soil"
	"github.com/ironarachne/world/pkg/species"
)

// Area is a geographic area and all of its component parts
type Area struct {
	Region   region.Region     `json:"region"`
	Climate  climate.Climate   `json:"climate"`
	Biome    biome.Biome       `json:"biome"`
	Seasons  []season.Season   `json:"seasons"`
	Animals  []species.Species `json:"animals"`
	Plants   []species.Species `json:"plants"`
	Minerals []mineral.Mineral `json:"minerals"`
	Soils    []soil.Soil       `json:"soils"`
}

const areaError = "failed to generate geographic area: %w"

// Generate procedurally generates a region, its climate, and its biome
func Generate() (Area, error) {
	r := region.Generate()
	c := climate.Generate(r)
	b, err := biome.Generate(c, r)
	if err != nil {
		err = fmt.Errorf(areaError, err)
		return Area{}, err
	}
	s, err := season.Generate(c, r)
	if err != nil {
		err = fmt.Errorf(areaError, err)
		return Area{}, err
	}

	animals, err := getAnimals(r.Humidity, r.Temperature, b.FaunaPrevalence, b.FaunaTags)
	if err != nil {
		err = fmt.Errorf(areaError, err)
		return Area{}, err
	}

	plants, err := getPlants(r.Humidity, r.Temperature, b.VegetationPrevalence, b.VegetationTags)
	if err != nil {
		err = fmt.Errorf(areaError, err)
		return Area{}, err
	}

	minerals, err := getMinerals()
	if err != nil {
		err = fmt.Errorf(areaError, err)
		return Area{}, err
	}

	soils, err := getSoils(r.NearestOceanDistance, r.Humidity, r.Temperature)

	a := Area{
		Region:   r,
		Climate:  c,
		Biome:    b,
		Seasons:  s,
		Animals:  animals,
		Plants:   plants,
		Minerals: minerals,
		Soils: soils,
	}

	return a, nil
}
