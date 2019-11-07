package plant

import (
	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/species"
)

func getGrains() []species.Species {
	grains := []species.Species{
		{
			Name:           "rye",
			PluralName:     "ryes",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 3,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "rye",
					Origin:       "rye",
					MainMaterial: "rye",
					Tags: []string{
						"grain",
					},
					Commonality: 5,
					Value:       1,
				},
			},
			Tags: []string{
				"grain",
			},
		},
		{
			Name:           "oat",
			PluralName:     "oats",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 3,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "oat",
					Origin:       "oat",
					MainMaterial: "oat",
					Tags: []string{
						"grain",
					},
					Commonality: 5,
					Value:       1,
				},
			},
			Tags: []string{
				"grain",
			},
		},
		{
			Name:           "millet",
			PluralName:     "millets",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 3,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "millet",
					Origin:       "millet",
					MainMaterial: "millet",
					Tags: []string{
						"grain",
					},
					Commonality: 5,
					Value:       1,
				},
			},
			Tags: []string{
				"grain",
			},
		},
		{
			Name:           "barley",
			PluralName:     "barleys",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 3,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "barley",
					Origin:       "barley",
					MainMaterial: "barley",
					Tags: []string{
						"grain",
					},
					Commonality: 5,
					Value:       1,
				},
			},
			Tags: []string{
				"grain",
			},
		},
		{
			Name:           "wheat",
			PluralName:     "wheats",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "wheat",
					Origin:       "wheat",
					MainMaterial: "wheat",
					Tags: []string{
						"grain",
					},
					Commonality: 5,
					Value:       1,
				},
			},
			Tags: []string{
				"grain",
			},
		},
		{
			Name:           "rice",
			PluralName:     "rices",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 3,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "rice",
					Origin:       "rice",
					MainMaterial: "rice",
					Tags: []string{
						"grain",
					},
					Commonality: 5,
					Value:       1,
				},
			},
			Tags: []string{
				"grain",
			},
		},
	}

	return grains
}
