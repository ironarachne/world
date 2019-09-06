package country

import (
	"fmt"
	"math/rand"

	"github.com/ironarachne/world/pkg/culture"
	"github.com/ironarachne/world/pkg/grid"
	"github.com/ironarachne/world/pkg/heraldry"
	"github.com/ironarachne/world/pkg/region"
	"github.com/ironarachne/world/pkg/worldmap"
)

// Country is a geographic and political area
type Country struct {
	Name            string
	DominantCulture culture.Culture
	Government      Government
	Regions         []region.Region
	Capital         string
	Heraldry        heraldry.Heraldry
}

// Generate procedurally generates a country
func Generate() (Country, error) {
	regions := []region.Region{}
	country := Country{}

	dominantCulture, err := culture.Random()
	if err != nil {
		err = fmt.Errorf("Could not generate country: %w", err)
		return Country{}, err
	}
	country.DominantCulture = dominantCulture
	government, err := country.getNewMonarchy()
	if err != nil {
		err = fmt.Errorf("Could not generate country: %w", err)
		return Country{}, err
	}
	country.Government = government
	country.Heraldry = heraldry.GenerateHeraldry()
	name, err := country.DominantCulture.Language.RandomName()
	if err != nil {
		err = fmt.Errorf("Could not generate country: %w", err)
		return Country{}, err
	}
	country.Name = name

	size := rand.Intn(10) + 4

	for i := 0; i < size; i++ {
		r, err := region.Generate(country.DominantCulture.HomeClimate, country.DominantCulture)
		if err != nil {
			err = fmt.Errorf("Could not generate country: %w", err)
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

// GetAllTiles returns a slice of all tiles in the country
func (c Country) GetAllTiles(worldMap worldmap.WorldMap) []worldmap.Tile {
	coords := c.GetAllTileCoordinates()

	tiles := worldmap.FindTilesByCoordinates(coords, worldMap.Tiles)

	return tiles
}

// GetAllTileCoordinates returns a slice of all coordinates in the country
func (c Country) GetAllTileCoordinates() []grid.Coordinate {
	coords := []grid.Coordinate{}

	for _, r := range c.Regions {
		for _, d := range r.TilesOccupied {
			coords = append(coords, d)
		}
	}

	return coords
}
