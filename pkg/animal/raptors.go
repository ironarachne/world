package animal

import (
	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/size"
)

func getRaptors() []Species {
	birds := []Species{
		{
			Name:           "eagle",
			PluralName:     "eagles",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "eagle feathers",
					Origin:       "eagle",
					MainMaterial: "eagle",
					Tags: []string{
						"feather",
					},
					Commonality: 4,
					Value:       10,
				},
			},
			Size: size.GetCategoryByName("tiny"),
			Tags: []string{
				"animal",
				"bird",
				"raptor",
				"carnivore",
			},
		},
		{
			Name:           "falcon",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "falcon feathers",
					Origin:       "falcon",
					MainMaterial: "falcon",
					Tags: []string{
						"feather",
					},
					Commonality: 4,
					Value:       5,
				},
			},
			Size: size.GetCategoryByName("tiny"),
			Tags: []string{
				"animal",
				"bird",
				"raptor",
				"carnivore",
			},
		},
		{
			Name:           "hawk",
			PluralName:     "hawks",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "hawk feathers",
					Origin:       "hawk",
					MainMaterial: "hawk",
					Tags: []string{
						"feather",
					},
					Commonality: 4,
					Value:       5,
				},
			},
			Size: size.GetCategoryByName("tiny"),
			Tags: []string{
				"animal",
				"bird",
				"raptor",
				"carnivore",
			},
		},
	}

	return birds
}
