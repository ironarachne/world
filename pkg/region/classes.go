package region

import (
	"context"
	"fmt"

	"github.com/ironarachne/world/pkg/random"
)

// Class is a class of region
type Class struct {
	MaxNumberOfTowns         int    `json:"max_number_of_towns"`
	MinNumberOfTowns         int    `json:"min_number_of_towns"`
	Name                     string `json:"name"`
	RulerTitleFemale         string `json:"ruler_title_female"`
	RulerTitleMale           string `json:"ruler_title_male"`
	Commonality              int    `json:"commonality"`
	NumberOfTiles            int    `json:"number_of_tiles"`
	MinNumberOfOrganizations int    `json:"min_number_of_organizations"`
	MaxNumberOfOrganizations int    `json:"max_number_of_organizations"`
}

func getAllClasses() []Class {
	classes := []Class{
		{
			Name:                     "Barony",
			RulerTitleFemale:         "Baroness",
			RulerTitleMale:           "Baron",
			MinNumberOfTowns:         2,
			MaxNumberOfTowns:         4,
			Commonality:              3,
			NumberOfTiles:            8,
			MinNumberOfOrganizations: 2,
			MaxNumberOfOrganizations: 5,
		},
		{
			Name:                     "County",
			RulerTitleFemale:         "Countess",
			RulerTitleMale:           "Count",
			MinNumberOfTowns:         1,
			MaxNumberOfTowns:         3,
			Commonality:              4,
			NumberOfTiles:            4,
			MinNumberOfOrganizations: 1,
			MaxNumberOfOrganizations: 4,
		},
		{
			Name:                     "Duchy",
			RulerTitleFemale:         "Duchess",
			RulerTitleMale:           "Duke",
			MinNumberOfTowns:         3,
			MaxNumberOfTowns:         6,
			Commonality:              2,
			NumberOfTiles:            16,
			MinNumberOfOrganizations: 3,
			MaxNumberOfOrganizations: 6,
		},
		{
			Name:                     "March",
			RulerTitleFemale:         "Marchioness",
			RulerTitleMale:           "Marquess",
			MinNumberOfTowns:         2,
			MaxNumberOfTowns:         4,
			Commonality:              2,
			NumberOfTiles:            2,
			MinNumberOfOrganizations: 1,
			MaxNumberOfOrganizations: 3,
		},
		{
			Name:                     "Viscounty",
			RulerTitleFemale:         "Viscountess",
			RulerTitleMale:           "Viscount",
			MinNumberOfTowns:         1,
			MaxNumberOfTowns:         2,
			Commonality:              1,
			NumberOfTiles:            1,
			MinNumberOfOrganizations: 1,
			MaxNumberOfOrganizations: 2,
		},
	}

	return classes
}

func getRandomWeightedClass(ctx context.Context) (Class, error) {
	classes := getAllClasses()

	weights := map[string]int{}

	for _, c := range classes {
		weights[c.Name] = c.Commonality
	}

	name, err := random.StringFromThresholdMap(ctx, weights)
	if err != nil {
		err = fmt.Errorf("failed to get random weighted region class: %w", err)
		return Class{}, err
	}

	for _, c := range classes {
		if c.Name == name {
			return c, nil
		}
	}

	err = fmt.Errorf("failed to get random weighted region class")
	return Class{}, err
}
