package region

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/ironarachne/world/pkg/climate"
	"github.com/ironarachne/world/pkg/culture"
	"github.com/ironarachne/world/pkg/grid"
	"github.com/ironarachne/world/pkg/organization"
	"github.com/ironarachne/world/pkg/town"
)

// Region is a map region
type Region struct {
	Biome         string
	Capital       string
	Class         Class
	Climate       climate.Climate
	Culture       culture.Culture
	Name          string
	Organizations []organization.Organization
	RulingBody    organization.Organization
	TilesOccupied []grid.Coordinate
	Towns         []town.Town
}

// AssignTiles gives a set of coordinates for tiles to a region
func (region Region) AssignTiles(coordinates []grid.Coordinate) Region {
	placedRegion := region

	placedRegion.TilesOccupied = coordinates

	return placedRegion
}

// Generate generates a region
func Generate(regionClimate climate.Climate, originCulture culture.Culture) (Region, error) {
	var nobleMembers []organization.Member
	region := Region{}

	region.Biome = regionClimate.Name
	region.Climate = regionClimate
	region.Culture = originCulture

	class, err := getRandomWeightedClass()
	if err != nil {
		err = fmt.Errorf("Could not generate region: %w", err)
		return Region{}, err
	}
	region.Class = class

	newTown, err := town.Generate("city", regionClimate, region.Culture)
	if err != nil {
		err = fmt.Errorf("Could not generate region: %w", err)
		return Region{}, err
	}
	region.Towns = append(region.Towns, newTown)

	region.Capital = newTown.Name

	for i := region.Class.MinNumberOfTowns - 1; i < region.Class.MaxNumberOfTowns-1; i++ {
		newTown, err = town.Generate("random", regionClimate, region.Culture)
		if err != nil {
			err = fmt.Errorf("Could not generate region: %w", err)
			return Region{}, err
		}
		region.Towns = append(region.Towns, newTown)
	}

	organizations, err := region.getOrganizations()
	if err != nil {
		err = fmt.Errorf("Could not generate region: %w", err)
		return Region{}, err
	}
	region.Organizations = organizations

	rulingBody, err := region.generateRulingBody()
	if err != nil {
		err = fmt.Errorf("Could not generate region: %w", err)
		return Region{}, err
	}
	region.RulingBody = rulingBody

	regionName, err := region.Culture.Language.RandomFamilyName()
	if err != nil {
		err = fmt.Errorf("Could not generate region: %w", err)
		return Region{}, err
	}
	region.Name = strings.Title(regionName)

	rulerTitle := region.RulingBody.Leader.CharacterData.Title + " of " + region.Name
	region.RulingBody.Leader.CharacterData.Titles = append(region.RulingBody.Leader.CharacterData.Titles, rulerTitle)

	for _, m := range region.RulingBody.NotableMembers {
		if m.Rank.Title == "Heir" {
			m.CharacterData.Titles = append(m.CharacterData.Titles, "Heir of "+region.Name)
		} else if m.Rank.Title == region.Class.RulerTitleFemale || m.Rank.Title == region.Class.RulerTitleMale {
			m.CharacterData.Titles = append(m.CharacterData.Titles, m.CharacterData.Title+" of "+region.Name)
		}

		nobleMembers = append(nobleMembers, m)
	}

	region.RulingBody.NotableMembers = nobleMembers

	return region, nil
}

func (region Region) getOrganizations() ([]organization.Organization, error) {
	organizations := []organization.Organization{}

	numberOfOrgs := rand.Intn(region.Class.MinNumberOfOrganizations+region.Class.MaxNumberOfOrganizations) + region.Class.MinNumberOfOrganizations

	for i := 0; i < numberOfOrgs; i++ {
		org, err := organization.Generate(region.Culture)
		if err != nil {
			err = fmt.Errorf("Could not generate region organizations: %w", err)
			return []organization.Organization{}, err
		}
		organizations = append(organizations, org)
	}

	return organizations, nil
}

// Random generates a completely random region
func Random() (Region, error) {
	randomCulture, err := culture.Random()
	if err != nil {
		err = fmt.Errorf("Could not generate random region: %w", err)
		return Region{}, err
	}

	region, err := Generate(randomCulture.HomeClimate, randomCulture)
	if err != nil {
		err = fmt.Errorf("Could not generate random region: %w", err)
		return Region{}, err
	}

	return region, nil
}
