package geography

import (
	"context"
	"fmt"

	"github.com/ironarachne/world/pkg/geography/biome"
	"github.com/ironarachne/world/pkg/geography/climate"
	"github.com/ironarachne/world/pkg/geography/region"
	"github.com/ironarachne/world/pkg/geography/season"
	"github.com/ironarachne/world/pkg/mineral"
	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/soil"
	"github.com/ironarachne/world/pkg/species"
	"github.com/ironarachne/world/pkg/words"
)

// Area is a geographic area and all of its component parts
type Area struct {
	Description string            `json:"description"`
	Region      region.Region     `json:"region"`
	Climate     climate.Climate   `json:"climate"`
	Biome       biome.Biome       `json:"biome"`
	Seasons     []season.Season   `json:"seasons"`
	Animals     []species.Species `json:"animals"`
	Plants      []species.Species `json:"plants"`
	Minerals    []mineral.Mineral `json:"minerals"`
	Soils       []soil.Soil       `json:"soils"`
}

const areaError = "failed to generate geographic area: %w"

// Describe provides a prose description of an area
func (area Area) Describe(ctx context.Context) (string, error) {
	description := "This area is " + words.Pronoun(area.Biome.Name) + " " + area.Biome.Name + ". The region is " + area.Region.Description + ". "

	for _, s := range area.Seasons {
		description += s.Description + " "
	}

	description += "The area's common animals include "
	description += describeSpecies(ctx, area.Animals)
	description += ". Common plants include "

	description += describeSpecies(ctx, area.Plants)

	description += "."

	return description, nil
}

func describeSpecies(ctx context.Context, from []species.Species) string {
	var list []string

	number := random.Intn(ctx, 2) + 3
	if number > len(from) {
		number = len(from)
	}

	selected := species.Random(ctx, number, from)

	for _, a := range selected {
		list = append(list, a.PluralName)
	}

	description := words.CombinePhrases(list)

	return description
}

// Generate procedurally generates a region, its climate, and its biome
func Generate(ctx context.Context) (Area, error) {
	r := region.Generate(ctx)
	c := climate.Generate(ctx, r)
	b, err := biome.Generate(ctx, c, r)
	if err != nil {
		err = fmt.Errorf(areaError, err)
		return Area{}, err
	}
	s, err := season.Generate(ctx, c, r)
	if err != nil {
		err = fmt.Errorf(areaError, err)
		return Area{}, err
	}

	animals, err := getAnimals(ctx, r.Humidity, r.Temperature, b.FaunaPrevalence, b.FaunaTags)
	if err != nil {
		err = fmt.Errorf(areaError, err)
		return Area{}, err
	}

	plants, err := getPlants(ctx, r.Humidity, r.Temperature, b.VegetationPrevalence, b.VegetationTags)
	if err != nil {
		err = fmt.Errorf(areaError, err)
		return Area{}, err
	}

	minerals, err := getMinerals(ctx)
	if err != nil {
		err = fmt.Errorf(areaError, err)
		return Area{}, err
	}

	soils, err := getSoils(ctx, r.NearestOceanDistance, r.Humidity, r.Temperature)

	a := Area{
		Region:   r,
		Climate:  c,
		Biome:    b,
		Seasons:  s,
		Animals:  animals,
		Plants:   plants,
		Minerals: minerals,
		Soils:    soils,
	}

	a.Description, err = a.Describe(ctx)
	if err != nil {
		err = fmt.Errorf(areaError, err)
		return Area{}, err
	}

	return a, nil
}

// GenerateSpecific generates a specific type of area based on parameters
func GenerateSpecific(ctx context.Context, temperature int, humidity int, altitude int, distance int) (Area, error) {
	r := region.GenerateSpecific(ctx, temperature, humidity, altitude, distance)
	c := climate.Generate(ctx, r)
	b, err := biome.Generate(ctx, c, r)
	if err != nil {
		err = fmt.Errorf(areaError, err)
		return Area{}, err
	}
	s, err := season.Generate(ctx, c, r)
	if err != nil {
		err = fmt.Errorf(areaError, err)
		return Area{}, err
	}

	animals, err := getAnimals(ctx, r.Humidity, r.Temperature, b.FaunaPrevalence, b.FaunaTags)
	if err != nil {
		err = fmt.Errorf(areaError, err)
		return Area{}, err
	}

	plants, err := getPlants(ctx, r.Humidity, r.Temperature, b.VegetationPrevalence, b.VegetationTags)
	if err != nil {
		err = fmt.Errorf(areaError, err)
		return Area{}, err
	}

	minerals, err := getMinerals(ctx)
	if err != nil {
		err = fmt.Errorf(areaError, err)
		return Area{}, err
	}

	soils, err := getSoils(ctx, r.NearestOceanDistance, r.Humidity, r.Temperature)

	a := Area{
		Region:   r,
		Climate:  c,
		Biome:    b,
		Seasons:  s,
		Animals:  animals,
		Plants:   plants,
		Minerals: minerals,
		Soils:    soils,
	}

	a.Description, err = a.Describe(ctx)
	if err != nil {
		err = fmt.Errorf(areaError, err)
		return Area{}, err
	}

	return a, nil
}

// GetResources returns all resources for a given area
func (area Area) GetResources() []resource.Resource {
	var resources []resource.Resource

	for _, a := range area.Animals {
		resources = append(resources, a.Resources...)
	}

	for _, p := range area.Plants {
		resources = append(resources, p.Resources...)
	}

	for _, m := range area.Minerals {
		resources = append(resources, m.Resources...)
	}

	for _, s := range area.Soils {
		resources = append(resources, s.Resources...)
	}

	return resources
}
