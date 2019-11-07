package plant

import (
	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/species"
)

func getMelons() []species.Species {
	plants := []species.Species{
		{
			Name:           "cantaloupe",
			PluralName:     "cantaloupes",
			MinHumidity:    5,
			MaxHumidity:    10,
			MinTemperature: 6,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "cantaloupe",
					Origin:       "cantaloupe",
					MainMaterial: "cantaloupe",
					Tags: []string{
						"fruit",
					},
					Commonality: 5,
					Value:       1,
				},
			},
			Tags: []string{
				"melon",
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
					Name:         "honeydew",
					Origin:       "honeydew",
					MainMaterial: "honeydew",
					Tags: []string{
						"fruit",
						"melon",
					},
					Commonality: 5,
					Value:       1,
				},
			},
			Tags: []string{
				"melon",
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
					Name:         "watermelon",
					Origin:       "watermelon",
					MainMaterial: "watermelon",
					Tags: []string{
						"fruit",
						"melon",
					},
					Commonality: 5,
					Value:       1,
				},
			},
			Tags: []string{
				"melon",
			},
		},
	}

	return plants
}
