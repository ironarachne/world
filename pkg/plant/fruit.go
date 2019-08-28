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
					Name:   "blackberry",
					Origin: "blackberry",
					Tags: []string{
						"berry",
						"fruit",
					},
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
					Name:   "blueberry",
					Origin: "blueberry",
					Tags: []string{
						"berry",
						"fruit",
					},
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
					Name:   "raspberry",
					Origin: "raspberry",
					Tags: []string{
						"berry",
						"fruit",
					},
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
					Name:   "strawberry",
					Origin: "strawberry",
					Tags: []string{
						"berry",
						"fruit",
					},
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
					Name:   "boysenberry",
					Origin: "boysenberry",
					Tags: []string{
						"berry",
						"fruit",
					},
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
					Name:   "gooseberry",
					Origin: "gooseberry",
					Tags: []string{
						"berry",
						"fruit",
					},
					Commonality: 5,
				},
			},
		},
		{
			Name:           "juniper",
			PluralName:     "juniper berries",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "juniper",
					Origin: "juniper",
					Tags: []string{
						"berry",
						"fruit",
					},
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
					Name:   "avocado",
					Origin: "avocado",
					Tags: []string{
						"fruit",
					},
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
					Name:   "cantaloupe",
					Origin: "cantaloupe",
					Tags: []string{
						"fruit",
					},
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
					Name:   "dragonfruit",
					Origin: "dragonfruit",
					Tags: []string{
						"fruit",
					},
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
					Name:   "honeydew",
					Origin: "honeydew",
					Tags: []string{
						"fruit",
						"melon",
					},
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
					Name:   "watermelon",
					Origin: "watermelon",
					Tags: []string{
						"fruit",
						"melon",
					},
					Commonality: 5,
				},
			},
		},
	}

	return fruits
}
