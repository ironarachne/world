package insect

import "github.com/ironarachne/world/pkg/resource"

func getInsects() []Insect {
	insects := []Insect{
		{
			Name:           "biting fly",
			PluralName:     "biting flies",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Tags: []string{
				"fly",
				"herbivore",
			},
		},
		{
			Name:           "cricket",
			PluralName:     "crickets",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "cricket",
					Origin:       "cricket",
					MainMaterial: "cricket",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
					Value:       1,
				},
			},
			Tags: []string{
				"food",
				"herbivore",
			},
		},
		{
			Name:           "dragonfly",
			PluralName:     "dragonflies",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Tags: []string{
				"fly",
				"carnivore",
			},
		},
		{
			Name:           "dwarf honey bee",
			PluralName:     "dwarf honey bees",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "beeswax",
					Origin:       "dwarf honey bee",
					MainMaterial: "beeswax",
					Tags: []string{
						"wax",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "honey",
					Origin:       "dwarf honey bee",
					MainMaterial: "honey",
					Tags: []string{
						"food",
						"sweetener",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "honeycomb",
					Origin:       "dwarf honey bee",
					MainMaterial: "honeycomb",
					Tags: []string{
						"food",
					},
					Commonality: 3,
					Value:       2,
				},
			},
			Tags: []string{
				"bee",
				"herbivore",
				"pollinator",
			},
		},
		{
			Name:           "earthworm",
			PluralName:     "earthworms",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Tags: []string{
				"worm",
				"herbivore",
			},
		},
		{
			Name:           "giant honey bee",
			PluralName:     "giant honey bees",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "beeswax",
					Origin:       "giant honey bee",
					MainMaterial: "beeswax",
					Tags: []string{
						"wax",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "honey",
					Origin:       "giant honey bee",
					MainMaterial: "honey",
					Tags: []string{
						"food",
						"sweetener",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "honeycomb",
					Origin:       "giant honey bee",
					MainMaterial: "honeycomb",
					Tags: []string{
						"food",
					},
					Commonality: 3,
					Value:       2,
				},
			},
			Tags: []string{
				"bee",
				"herbivore",
				"pollinator",
			},
		},
		{
			Name:           "housefly",
			PluralName:     "houseflies",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Tags: []string{
				"fly",
				"herbivore",
			},
		},
		{
			Name:           "honey bee",
			PluralName:     "honey bees",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "beeswax",
					Origin:       "honey bee",
					MainMaterial: "beeswax",
					Tags: []string{
						"wax",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "honey",
					Origin:       "honey bee",
					MainMaterial: "honey",
					Tags: []string{
						"food",
						"sweetener",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "honeycomb",
					Origin:       "honey bee",
					MainMaterial: "honeycomb",
					Tags: []string{
						"food",
					},
					Commonality: 3,
					Value:       2,
				},
			},
			Tags: []string{
				"bee",
				"herbivore",
				"pollinator",
			},
		},
		{
			Name:           "mosquito",
			PluralName:     "mosquitoes",
			MinHumidity:    5,
			MaxHumidity:    10,
			MinTemperature: 5,
			MaxTemperature: 10,
			Tags: []string{
				"pest",
				"carnivore",
			},
		},
		{
			Name:           "silkworm",
			PluralName:     "silkworms",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "silk",
					Origin:       "silkworm",
					MainMaterial: "silk",
					Tags: []string{
						"fabric fiber",
					},
					Commonality: 3,
					Value:       20,
				},
			},
			Tags: []string{
				"worm",
				"herbivore",
			},
		},
	}

	return insects
}
