package plant

import "github.com/ironarachne/world/pkg/resource"

func getFruits() []Plant {
	fruits := []Plant{
		Plant{
			Name:           "blackberry",
			PluralName:     "blackberries",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "blackberry",
					Origin:      "blackberry",
					Type:        "fruit",
					Commonality: 5,
				},
			},
		},
		Plant{
			Name:           "blueberry",
			PluralName:     "blueberries",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "blueberry",
					Origin:      "blueberry",
					Type:        "fruit",
					Commonality: 5,
				},
			},
		},
		Plant{
			Name:           "raspberry",
			PluralName:     "raspberries",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "raspberry",
					Origin:      "raspberry",
					Type:        "fruit",
					Commonality: 5,
				},
			},
		},
		Plant{
			Name:           "strawberry",
			PluralName:     "strawberries",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "strawberry",
					Origin:      "strawberry",
					Type:        "fruit",
					Commonality: 5,
				},
			},
		},
		Plant{
			Name:           "boysenberry",
			PluralName:     "boysenberries",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "boysenberry",
					Origin:      "boysenberry",
					Type:        "fruit",
					Commonality: 5,
				},
			},
		},
		Plant{
			Name:           "gooseberry",
			PluralName:     "gooseberries",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "gooseberry",
					Origin:      "gooseberry",
					Type:        "fruit",
					Commonality: 5,
				},
			},
		},
		Plant{
			Name:           "avocado",
			PluralName:     "avocados",
			MinHumidity:    6,
			MaxHumidity:    10,
			MinTemperature: 6,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "avocado",
					Origin:      "avocado",
					Type:        "fruit",
					Commonality: 5,
				},
			},
		},
		Plant{
			Name:           "cantaloupe",
			PluralName:     "cantaloupes",
			MinHumidity:    5,
			MaxHumidity:    10,
			MinTemperature: 6,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "cantaloupe",
					Origin:      "cantaloupe",
					Type:        "fruit",
					Commonality: 5,
				},
			},
		},
		Plant{
			Name:           "dragonfruit",
			PluralName:     "dragonfruits",
			MinHumidity:    5,
			MaxHumidity:    10,
			MinTemperature: 6,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "dragonfruit",
					Origin:      "dragonfruit",
					Type:        "fruit",
					Commonality: 5,
				},
			},
		},
		Plant{
			Name:           "honeydew",
			PluralName:     "honeydews",
			MinHumidity:    5,
			MaxHumidity:    10,
			MinTemperature: 6,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "honeydew",
					Origin:      "honeydew",
					Type:        "fruit",
					Commonality: 5,
				},
			},
		},
		Plant{
			Name:           "pumpkin",
			PluralName:     "pumpkins",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 3,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "pumpkin",
					Origin:      "pumpkin",
					Type:        "fruit",
					Commonality: 5,
				},
			},
		},
		Plant{
			Name:           "watermelon",
			PluralName:     "watermelons",
			MinHumidity:    5,
			MaxHumidity:    10,
			MinTemperature: 6,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "watermelon",
					Origin:      "watermelon",
					Type:        "fruit",
					Commonality: 5,
				},
			},
		},
		Plant{
			Name:           "zuccini",
			PluralName:     "zuccinis",
			MinHumidity:    2,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				resource.Resource{
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
