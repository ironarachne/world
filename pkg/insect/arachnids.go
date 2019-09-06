package insect

import "github.com/ironarachne/world/pkg/resource"

func getArachnids() []Insect {
	insects := []Insect{
		{
			Name:           "black widow",
			PluralName:     "black widows",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 7,
			Resources: []resource.Resource{
				{
					Name:         "black widow venom",
					Origin:       "black widow",
					MainMaterial: "black widow",
					Tags: []string{
						"venom",
					},
					Commonality: 3,
					Value:       20,
				},
			},
			Tags: []string{
				"spider",
				"arachnid",
				"carnivore",
			},
		},
		{
			Name:           "brown recluse",
			PluralName:     "brown recluses",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 7,
			Resources: []resource.Resource{
				{
					Name:         "brown recluse venom",
					Origin:       "brown recluse",
					MainMaterial: "brown recluse",
					Tags: []string{
						"venom",
					},
					Commonality: 3,
					Value:       20,
				},
			},
			Tags: []string{
				"spider",
				"arachnid",
				"carnivore",
			},
		},
		{
			Name:           "emperor scorpion",
			PluralName:     "emperor scorpions",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 7,
			Resources: []resource.Resource{
				{
					Name:         "emperor scorpion venom",
					Origin:       "emperor scorpion",
					MainMaterial: "emperor scorpion",
					Tags: []string{
						"venom",
					},
					Commonality: 2,
					Value:       50,
				},
			},
			Tags: []string{
				"scorpion",
				"arachnid",
				"carnivore",
			},
		},
		{
			Name:           "house spider",
			PluralName:     "house spiders",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 7,
			Tags: []string{
				"spider",
				"arachnid",
				"carnivore",
			},
		},
		{
			Name:           "scorpion",
			PluralName:     "scorpions",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 7,
			Resources: []resource.Resource{
				{
					Name:         "scorpion venom",
					Origin:       "scorpion",
					MainMaterial: "scorpion",
					Tags: []string{
						"venom",
					},
					Commonality: 3,
					Value:       20,
				},
			},
			Tags: []string{
				"scorpion",
				"arachnid",
				"carnivore",
			},
		},
		{
			Name:           "tick",
			PluralName:     "ticks",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Tags: []string{
				"arachnid",
				"carnivore",
			},
		},
	}

	return insects
}
