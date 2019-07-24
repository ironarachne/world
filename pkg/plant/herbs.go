package plant

import "github.com/ironarachne/world/pkg/resource"

func getHerbs() []Plant {
	herbs := []Plant{
		{
			Name:           "parsley",
			PluralName:     "parsleys",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 3,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "parsley",
					Origin: "parsley",
					Tags: []string{
						"herb",
					},
					Commonality: 5,
				},
			},
		},
		{
			Name:           "brahmi",
			PluralName:     "brahmis",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 3,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "brahmi",
					Origin: "brahmi",
					Tags: []string{
						"herb",
					},
					Commonality: 5,
				},
			},
		},
		{
			Name:           "basil",
			PluralName:     "basils",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 3,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "basil",
					Origin: "basil",
					Tags: []string{
						"herb",
					},
					Commonality: 5,
				},
			},
		},
		{
			Name:           "cilantro",
			PluralName:     "cilantros",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 3,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "cilantro",
					Origin: "cilantro",
					Tags: []string{
						"herb",
					},
					Commonality: 5,
				},
			},
		},
		{
			Name:           "saffron",
			PluralName:     "saffrons",
			MinHumidity:    5,
			MaxHumidity:    10,
			MinTemperature: 6,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "saffron",
					Origin: "saffron",
					Tags: []string{
						"herb",
					},
					Commonality: 3,
				},
			},
		},
		{
			Name:           "rosemary",
			PluralName:     "rosemarys",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 3,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "rosemary",
					Origin: "rosemary",
					Tags: []string{
						"herb",
					},
					Commonality: 5,
				},
			},
		},
		{
			Name:           "thyme",
			PluralName:     "thymes",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 3,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "thyme",
					Origin: "thyme",
					Tags: []string{
						"herb",
					},
					Commonality: 5,
				},
			},
		},
		{
			Name:           "black pepper",
			PluralName:     "black peppers",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 3,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "black pepper",
					Origin: "black pepper",
					Tags: []string{
						"spice",
					},
					Commonality: 5,
				},
			},
		},
		{
			Name:           "oregano",
			PluralName:     "oreganos",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 3,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "oregano",
					Origin: "oregano",
					Tags: []string{
						"herb",
					},
					Commonality: 5,
				},
			},
		},
		{
			Name:           "mint",
			PluralName:     "mints",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 3,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "mint",
					Origin: "mint",
					Tags: []string{
						"herb",
					},
					Commonality: 5,
				},
			},
		},
		{
			Name:           "lemongrass",
			PluralName:     "lemongrasses",
			MinHumidity:    5,
			MaxHumidity:    10,
			MinTemperature: 7,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "lemongrass",
					Origin: "lemongrass",
					Tags: []string{
						"herb",
					},
					Commonality: 5,
				},
			},
		},
		{
			Name:           "coriander",
			PluralName:     "corianders",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 3,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "coriander",
					Origin: "coriander",
					Tags: []string{
						"spice",
					},
					Commonality: 5,
				},
			},
		},
		{
			Name:           "clove",
			PluralName:     "cloves",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 3,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "clove",
					Origin: "clove",
					Tags: []string{
						"spice",
					},
					Commonality: 5,
				},
			},
		},
		{
			Name:           "sage",
			PluralName:     "sages",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 3,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "sage",
					Origin: "sage",
					Tags: []string{
						"herb",
					},
					Commonality: 5,
				},
			},
		},
		{
			Name:           "turmeric",
			PluralName:     "turmerics",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 7,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "turmeric",
					Origin: "turmeric",
					Tags: []string{
						"spice",
					},
					Commonality: 5,
				},
			},
		},
		{
			Name:           "vanilla",
			PluralName:     "vanillas",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 7,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "vanilla",
					Origin: "vanilla",
					Tags: []string{
						"herb",
					},
					Commonality: 3,
				},
			},
		},
		{
			Name:           "galangal",
			PluralName:     "galangals",
			MinHumidity:    5,
			MaxHumidity:    10,
			MinTemperature: 7,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "galangal",
					Origin: "galangal",
					Tags: []string{
						"spice",
					},
					Commonality: 5,
				},
			},
		},
		{
			Name:           "ginger",
			PluralName:     "gingers",
			MinHumidity:    5,
			MaxHumidity:    10,
			MinTemperature: 7,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "ginger",
					Origin: "ginger",
					Tags: []string{
						"spice",
					},
					Commonality: 5,
				},
			},
		},
	}

	return herbs
}
