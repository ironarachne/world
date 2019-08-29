package plant

import "github.com/ironarachne/world/pkg/resource"

func getMelons() []Plant {
	plants := []Plant{
		{
			Name:           "cantaloupe",
			PluralName:     "cantaloupes",
			MinHumidity:    5,
			MaxHumidity:    10,
			MinTemperature: 6,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "cantaloupe",
					Origin: "cantaloupe",
					Tags: []string{
						"fruit",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "honeydew",
			PluralName:     "honeydews",
			MinHumidity:    5,
			MaxHumidity:    10,
			MinTemperature: 6,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "honeydew",
					Origin: "honeydew",
					Tags: []string{
						"fruit",
						"melon",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "watermelon",
			PluralName:     "watermelons",
			MinHumidity:    5,
			MaxHumidity:    10,
			MinTemperature: 6,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "watermelon",
					Origin: "watermelon",
					Tags: []string{
						"fruit",
						"melon",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
	}

	return plants
}
