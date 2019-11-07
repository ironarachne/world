package animal

import (
	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/species"
)

func getBirds() []species.Species {
	birds := []species.Species{
		{
			Name:           "flamingo",
			PluralName:     "flamingos",
			MinHumidity:    5,
			MaxHumidity:    10,
			MinTemperature: 5,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "flamingo feathers",
					Origin:       "flamingo",
					MainMaterial: "flamingo",
					Tags: []string{
						"feather",
					},
					Commonality: 4,
					Value:       1,
				},
			},
			Tags: []string{
				"animal",
				"bird",
			},
		},
		{
			Name:           "peacock",
			PluralName:     "peacocks",
			MinHumidity:    5,
			MaxHumidity:    10,
			MinTemperature: 5,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "peacock feathers",
					Origin:       "peacock",
					MainMaterial: "peacock",
					Tags: []string{
						"feather",
					},
					Commonality: 4,
					Value:       1,
				},
			},
			Tags: []string{
				"animal",
				"bird",
			},
		},
	}

	return birds
}
