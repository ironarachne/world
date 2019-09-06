package plant

import (
	"github.com/ironarachne/world/pkg/resource"
)

func getFibers() []Plant {
	fibers := []Plant{
		{
			Name:           "cotton",
			PluralName:     "cotton",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "cotton",
					Origin:       "cotton",
					MainMaterial: "cotton",
					Tags: []string{
						"fabric fiber",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "flax",
			PluralName:     "flax",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "linen",
					Origin:       "flax",
					MainMaterial: "linen",
					Tags: []string{
						"fabric fiber",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "hemp",
			PluralName:     "hemp",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 9,
			Resources: []resource.Resource{
				{
					Name:         "hemp",
					Origin:       "hemp",
					MainMaterial: "hemp",
					Tags: []string{
						"fabric fiber",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "coir",
			PluralName:     "coir",
			MinHumidity:    6,
			MaxHumidity:    10,
			MinTemperature: 8,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "coir seed",
					Origin:       "coir",
					MainMaterial: "coir",
					Tags: []string{
						"seed",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "papyrus",
			PluralName:     "papyrus",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 9,
			Resources: []resource.Resource{
				{
					Name:         "papyrus",
					Origin:       "papyrus",
					MainMaterial: "papyrus",
					Tags: []string{
						"paper",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "jute",
			PluralName:     "jute",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 9,
			Resources: []resource.Resource{
				{
					Name:         "jute",
					Origin:       "jute",
					MainMaterial: "jute",
					Tags: []string{
						"fabric fiber",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "ramie",
			PluralName:     "ramie",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "ramie",
					Origin:       "ramie",
					MainMaterial: "ramie",
					Tags: []string{
						"fabric fiber",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
	}

	return fibers
}
