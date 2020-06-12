/*
Package country provides structures and tools for generating fantasy countries.
*/
package country

import (
	"context"
	"fmt"

	"github.com/ironarachne/world/pkg/geography"
	"github.com/ironarachne/world/pkg/random"

	"github.com/ironarachne/world/pkg/culture"
	"github.com/ironarachne/world/pkg/geometry"
	"github.com/ironarachne/world/pkg/heraldry"
	"github.com/ironarachne/world/pkg/region"
)

const countryError = "failed to generate country: %w"

// Country is a geographic and political area
type Country struct {
	Name            string
	DominantCulture culture.Culture
	Government      Government
	Regions         []region.Region
	Capital         string
	Heraldry        heraldry.Device
}

// Generate procedurally generates a country
func Generate(ctx context.Context) (Country, error) {
	regions := []region.Region{}
	country := Country{}

	originArea, err := geography.Generate(ctx)
	if err != nil {
		err = fmt.Errorf(countryError, err)
		return Country{}, err
	}

	dominantCulture, err := culture.Generate(ctx, originArea)
	if err != nil {
		err = fmt.Errorf(countryError, err)
		return Country{}, err
	}
	country.DominantCulture = dominantCulture
	government, err := country.getNewMonarchy(ctx)
	if err != nil {
		err = fmt.Errorf(countryError, err)
		return Country{}, err
	}
	country.Government = government
	device, err := heraldry.Generate(ctx)
	if err != nil {
		err = fmt.Errorf(countryError, err)
		return Country{}, err
	}
	country.Heraldry = device
	name, err := country.DominantCulture.Language.RandomFamilyName(ctx)
	if err != nil {
		err = fmt.Errorf(countryError, err)
		return Country{}, err
	}
	country.Name = name

	size := random.Intn(ctx, 10) + 4

	for i := 0; i < size; i++ {
		r, err := region.Generate(ctx, originArea, country.DominantCulture)
		if err != nil {
			err = fmt.Errorf(countryError, err)
			return Country{}, err
		}
		regions = append(regions, r)
	}

	highestPopulation := 0
	capital := ""

	for _, r := range regions {
		for _, t := range r.Towns {
			if t.Population > highestPopulation {
				highestPopulation = t.Population
				capital = t.Name
			}
		}
	}

	country.Capital = capital

	country.Regions = regions

	return country, nil
}

// GetAllTileCoordinates returns a slice of all coordinates in the country
func (c Country) GetAllTileCoordinates() []geometry.Point {
	var coords []geometry.Point

	for _, r := range c.Regions {
		for _, d := range r.TilesOccupied {
			coords = append(coords, d)
		}
	}

	return coords
}
