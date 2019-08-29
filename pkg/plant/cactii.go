package plant

import "github.com/ironarachne/world/pkg/resource"

func getCactii() []Plant {
	plants := []Plant{
		{
			Name:           "agave",
			PluralName:     "agave",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 6,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "agave fiber",
					Origin: "agave",
					Tags: []string{
						"fiber",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:   "agave needle",
					Origin: "agave",
					Tags: []string{
						"needle",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:   "agave sap",
					Origin: "agave",
					Tags: []string{
						"sap",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "pitahaya",
			PluralName:     "pitahayas",
			MinHumidity:    5,
			MaxHumidity:    10,
			MinTemperature: 6,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "dragonfruit",
					Origin: "pitahaya",
					Tags: []string{
						"fruit",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "saguaro",
			PluralName:     "saguaro",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 7,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "saguaro fiber",
					Origin: "saguaro",
					Tags: []string{
						"fiber",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:   "saguaro",
					Origin: "saguaro",
					Tags: []string{
						"fruit",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:   "saguaro needle",
					Origin: "saguaro",
					Tags: []string{
						"needle",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:   "saguaro sap",
					Origin: "saguaro",
					Tags: []string{
						"sap",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
	}

	return plants
}
