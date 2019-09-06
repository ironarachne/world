package plant

import "github.com/ironarachne/world/pkg/resource"

func getSquash() []Plant {
	plants := []Plant{
		{
			Name:           "acorn squash",
			PluralName:     "acorn squashes",
			MinHumidity:    1,
			MaxHumidity:    10,
			MinTemperature: 1,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "acorn squash",
					Origin:       "acorn squash",
					MainMaterial: "acorn squash",
					Tags: []string{
						"squash",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "buttercup squash",
			PluralName:     "buttercup squashes",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 3,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "buttercup squash",
					Origin:       "buttercup squash",
					MainMaterial: "buttercup squash",
					Tags: []string{
						"squash",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "butternut squash",
			PluralName:     "butternut squashes",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 3,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "butternut squash",
					Origin:       "butternut squash",
					MainMaterial: "butternut squash",
					Tags: []string{
						"squash",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "pumpkin",
			PluralName:     "pumpkins",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 1,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "pumpkin",
					Origin:       "pumpkin",
					MainMaterial: "pumpkin",
					Tags: []string{
						"squash",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "pumpkin seed oil",
					Origin:       "pumpkin",
					MainMaterial: "pumpkin",
					Commonality:  2,
					Value:        1,
				},
			},
		},
		{
			Name:           "spaghetti squash",
			PluralName:     "spaghetti squashes",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 3,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "spaghetti squash",
					Origin:       "spaghetti squash",
					MainMaterial: "spaghetti squash",
					Tags: []string{
						"squash",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "zucchini",
			PluralName:     "zucchini",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "zucchini",
					Origin:       "zucchini",
					MainMaterial: "zucchini",
					Tags: []string{
						"squash",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
	}

	return plants
}
