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
					Name:        "parsley",
					Origin:      "parsley",
					Type:        "herb",
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
					Name:        "brahmi",
					Origin:      "brahmi",
					Type:        "herb",
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
					Name:        "basil",
					Origin:      "basil",
					Type:        "herb",
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
					Name:        "cilantro",
					Origin:      "cilantro",
					Type:        "herb",
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
					Name:        "saffron",
					Origin:      "saffron",
					Type:        "herb",
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
					Name:        "rosemary",
					Origin:      "rosemary",
					Type:        "herb",
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
					Name:        "thyme",
					Origin:      "thyme",
					Type:        "herb",
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
					Name:        "black pepper",
					Origin:      "black pepper",
					Type:        "spice",
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
					Name:        "oregano",
					Origin:      "oregano",
					Type:        "herb",
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
					Name:        "mint",
					Origin:      "mint",
					Type:        "herb",
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
					Name:        "lemongrass",
					Origin:      "lemongrass",
					Type:        "herb",
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
					Name:        "coriander",
					Origin:      "coriander",
					Type:        "spice",
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
					Name:        "clove",
					Origin:      "clove",
					Type:        "spice",
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
					Name:        "sage",
					Origin:      "sage",
					Type:        "herb",
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
					Name:        "turmeric",
					Origin:      "turmeric",
					Type:        "spice",
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
					Name:        "vanilla",
					Origin:      "vanilla",
					Type:        "herb",
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
					Name:        "galangal",
					Origin:      "galangal",
					Type:        "spice",
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
					Name:        "ginger",
					Origin:      "ginger",
					Type:        "spice",
					Commonality: 5,
				},
			},
		},
	}

	return herbs
}
