/*
Package region provides tools and structures for fantasy regions of a world
*/
package region

import (
	"context"
	"fmt"
	"strings"

	"github.com/ironarachne/world/pkg/geography"
	"github.com/ironarachne/world/pkg/random"

	"github.com/ironarachne/world/pkg/culture"
	"github.com/ironarachne/world/pkg/geometry"
	"github.com/ironarachne/world/pkg/organization"
	"github.com/ironarachne/world/pkg/town"
)

const regionError = "failed to generate region: %w"

// Region is a map region
type Region struct {
	AreaDescription string                      `json:"area_description"`
	Biome           string                      `json:"biome"`
	Capital         string                      `json:"capital"`
	Class           Class                       `json:"class"`
	Geography       geography.Area              `json:"geography"`
	Culture         culture.Culture             `json:"culture"`
	Name            string                      `json:"name"`
	Organizations   []organization.Organization `json:"organizations"`
	RulingBody      organization.Organization   `json:"ruling_body"`
	TilesOccupied   []geometry.Point            `json:"tiles_occupied"`
	Towns           []town.Town                 `json:"town"`
}

const regionGenerationError = "failed to generate random region: %w"

// AssignTiles gives a set of coordinates for tiles to a region
func (region Region) AssignTiles(coordinates []geometry.Point) Region {
	placedRegion := region

	placedRegion.TilesOccupied = coordinates

	return placedRegion
}

// Generate generates a region
func Generate(ctx context.Context, area geography.Area, originCulture culture.Culture) (Region, error) {
	var nobleMembers []organization.Member
	region := Region{}

	region.AreaDescription = area.Description
	region.Biome = area.Biome.Name
	region.Geography = area
	region.Culture = originCulture

	class, err := getRandomWeightedClass(ctx)
	if err != nil {
		err = fmt.Errorf(regionError, err)
		return Region{}, err
	}
	region.Class = class

	newTown, err := town.Generate(ctx, "city", area, region.Culture)
	if err != nil {
		err = fmt.Errorf(regionError, err)
		return Region{}, err
	}
	region.Towns = append(region.Towns, newTown)

	region.Capital = newTown.Name

	for i := region.Class.MinNumberOfTowns - 1; i < region.Class.MaxNumberOfTowns-1; i++ {
		newTown, err = town.Generate(ctx, "random", area, region.Culture)
		if err != nil {
			err = fmt.Errorf(regionError, err)
			return Region{}, err
		}
		region.Towns = append(region.Towns, newTown)
	}

	organizations, err := region.getOrganizations(ctx)
	if err != nil {
		err = fmt.Errorf(regionError, err)
		return Region{}, err
	}
	region.Organizations = organizations

	rulingBody, err := region.generateRulingBody(ctx)
	if err != nil {
		err = fmt.Errorf(regionError, err)
		return Region{}, err
	}
	region.RulingBody = rulingBody

	regionName, err := region.Culture.Language.RandomFamilyName(ctx)
	if err != nil {
		err = fmt.Errorf(regionError, err)
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

func (region Region) getOrganizations(ctx context.Context) ([]organization.Organization, error) {
	organizations := []organization.Organization{}

	numberOfOrgs := random.Intn(ctx, region.Class.MinNumberOfOrganizations+region.Class.MaxNumberOfOrganizations) + region.Class.MinNumberOfOrganizations

	for i := 0; i < numberOfOrgs; i++ {
		org, err := organization.Generate(ctx, region.Culture)
		if err != nil {
			err = fmt.Errorf("failed to generate region organizations: %w", err)
			return []organization.Organization{}, err
		}
		organizations = append(organizations, org)
	}

	return organizations, nil
}

// Random generates a completely random region
func Random(ctx context.Context) (Region, error) {
	area, err := geography.Generate(ctx)
	if err != nil {
		err = fmt.Errorf(regionGenerationError, err)
		return Region{}, err
	}

	randomCulture, err := culture.Generate(ctx, area)
	if err != nil {
		err = fmt.Errorf(regionGenerationError, err)
		return Region{}, err
	}

	region, err := Generate(ctx, area, randomCulture)
	if err != nil {
		err = fmt.Errorf(regionGenerationError, err)
		return Region{}, err
	}

	return region, nil
}
