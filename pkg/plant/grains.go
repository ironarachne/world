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
		},
	}

	return grains
}
