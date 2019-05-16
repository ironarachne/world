package region

import (
	"math/rand"
	"strings"

	"github.com/ironarachne/world/pkg/character"
	"github.com/ironarachne/world/pkg/climate"
	"github.com/ironarachne/world/pkg/culture"
	"github.com/ironarachne/world/pkg/grid"
	"github.com/ironarachne/world/pkg/organization"
	"github.com/ironarachne/world/pkg/town"
)

// Region is a map region
type Region struct {
	Biome         string
	Culture       culture.Culture
	Climate       climate.Climate
	Capital       string
	Class         Class
	Name          string
	Ruler         character.Character
	Towns         []town.Town
	Organizations []organization.Organization
	TilesOccupied []grid.Coordinate
}

// AssignTiles gives a set of coordinates for tiles to a region
func (region Region) AssignTiles(coordinates []grid.Coordinate) Region {
	placedRegion := region

	placedRegion.TilesOccupied = coordinates

	return placedRegion
}

// Generate generates a random region
func Generate(regionType string) Region {
	region := Region{}
	biome := climate.Climate{}

	if regionType == "random" {
		biome = climate.Generate()
	} else {
		biome = climate.GetClimate(regionType)
	}

	regionType = biome.Name

	region.Biome = biome.Name
	region.Climate = biome
	region.Culture = culture.Generate()
	region.Culture = region.Culture.SetClimate(region.Biome)

	region.Class = getRandomWeightedClass()

	newTown := town.Generate("city", regionType)
	newTown = town.SetCulture(region.Culture, newTown)
	region.Towns = append(region.Towns, newTown)

	region.Capital = newTown.Name

	for i := region.Class.MinNumberOfTowns - 1; i < region.Class.MaxNumberOfTowns-1; i++ {
		newTown = town.Generate("random", regionType)
		newTown = town.SetCulture(region.Culture, newTown)
		region.Towns = append(region.Towns, newTown)
	}

	region.Organizations = region.getOrganizations()

	region.Ruler = region.generateRuler()

	regionName := region.Culture.Language.RandomName()
	region.Name = strings.Title(regionName)

	return region
}

func (region Region) getOrganizations() []organization.Organization {
	var org organization.Organization
	organizations := []organization.Organization{}

	numberOfOrgs := rand.Intn(3) + 1

	for i := 0; i < numberOfOrgs; i++ {
		org = organization.Generate()
		org = org.SetCulture(region.Culture)
		organizations = append(organizations, org)
	}

	return organizations
}

// SetCulture sets the culture for a region and derived values
func (region Region) SetCulture(culture.Culture) Region {
	region.Culture = region.Culture.SetClimate(region.Biome)
	newTown := town.Generate("city", region.Biome)
	newTown = town.SetCulture(region.Culture, newTown)
	region.Towns = append(region.Towns, newTown)

	region.Capital = newTown.Name

	for i := region.Class.MinNumberOfTowns - 1; i < region.Class.MaxNumberOfTowns-1; i++ {
		newTown = town.Generate("random", region.Biome)
		newTown = town.SetCulture(region.Culture, newTown)
		region.Towns = append(region.Towns, newTown)
	}

	region.Organizations = region.getOrganizations()

	region.Ruler = region.generateRuler()

	regionName := region.Culture.Language.RandomName()
	region.Name = strings.Title(regionName)

	return region
}
