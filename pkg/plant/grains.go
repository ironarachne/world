package plant

import "github.com/ironarachne/world/pkg/resource"

func getGrains() []Plant {
	grains := []Plant{
		{
			Name:           "rye",
			PluralName:     "ryes",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 3,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "rye",
					Origin: "rye",
					Tags: []string{
						"grain",
					},
					Commonality: 5,
					Value:       1,
				},
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
					Name:   "oat",
					Origin: "oat",
					Tags: []string{
						"grain",
					},
					Commonality: 5,
					Value:       1,
				},
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
					Name:   "millet",
					Origin: "millet",
					Tags: []string{
						"grain",
					},
					Commonality: 5,
					Value:       1,
				},
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
					Name:   "barley",
					Origin: "barley",
					Tags: []string{
						"grain",
					},
					Commonality: 5,
					Value:       1,
				},
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
					Name:   "wheat",
					Origin: "wheat",
					Tags: []string{
						"grain",
					},
					Commonality: 5,
					Value:       1,
				},
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
					Name:   "rice",
					Origin: "rice",
					Tags: []string{
						"grain",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
	}

	return grains
}
