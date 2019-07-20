package plant

import "github.com/ironarachne/world/pkg/resource"

func getFruits() []Plant {
	fruits := []Plant{
		{
			Name:           "blackberry",
			PluralName:     "blackberries",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:        "blackberry",
					Origin:      "blackberry",
					Type:        "fruit",
					Commonality: 5,
				},
			},
		},
		{
			Name:           "blueberry",
			PluralName:     "blueberries",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:        "blueberry",
					Origin:      "blueberry",
					Type:        "fruit",
					Commonality: 5,
				},
			},
		},
		{
			Name:           "raspberry",
			PluralName:     "raspberries",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:        "raspberry",
					Origin:      "raspberry",
					Type:        "fruit",
					Commonality: 5,
				},
			},
		},
		{
			Name:           "strawberry",
			PluralName:     "strawberries",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:        "strawberry",
					Origin:      "strawberry",
					Type:        "fruit",
					Commonality: 5,
				},
			},
		},
		{
			Name:           "boysenberry",
			PluralName:     "boysenberries",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:        "boysenberry",
					Origin:      "boysenberry",
					Type:        "fruit",
					Commonality: 5,
				},
			},
		},
		{
			Name:           "gooseberry",
			PluralName:     "gooseberries",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:        "gooseberry",
					Origin:      "gooseberry",
					Type:        "fruit",
					Commonality: 5,
				},
			},
		},
		{
			Name:           "avocado",
			PluralName:     "avocados",
			MinHumidity:    6,
			MaxHumidity:    10,
			MinTemperature: 6,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:        "avocado",
					Origin:      "avocado",
					Type:        "fruit",
					Commonality: 5,
				},
			},
		},
		{
			Name:           "cantaloupe",
			PluralName:     "cantaloupes",
			MinHumidity:    5,
			MaxHumidity:    10,
			MinTemperature: 6,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:        "cantaloupe",
					Origin:      "cantaloupe",
					Type:        "fruit",
					Commonality: 5,
				},
			},
		},
		{
			Name:           "dragonfruit",
			PluralName:     "dragonfruits",
			MinHumidity:    5,
			MaxHumidity:    10,
			MinTemperature: 6,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:        "dragonfruit",
					Origin:      "dragonfruit",
					Type:        "fruit",
					Commonality: 5,
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
					Name:        "honeydew",
					Origin:      "honeydew",
					Type:        "fruit",
					Commonality: 5,
				},
			},
		},
		{
			Name:           "pumpkin",
			PluralName:     "pumpkins",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 3,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:        "pumpkin",
					Origin:      "pumpkin",
					Type:        "fruit",
					Commonality: 5,
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
					Name:        "watermelon",
					Origin:      "watermelon",
					Type:        "fruit",
					Commonality: 5,
				},
			},
		},
		{
			Name:           "zuccini",
			PluralName:     "zuccinis",
			MinHumidity:    2,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:        "zuccini",
					Origin:      "zuccini",
					Type:        "fruit",
					Commonality: 5,
				},
			},
		},
	}

	return fruits
}
