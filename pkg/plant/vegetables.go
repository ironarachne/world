package plant

import "github.com/ironarachne/world/pkg/resource"

func getVegetables() []Plant {
	vegetables := []Plant{
		{
			Name:           "potato",
			PluralName:     "potatoes",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "potato",
					Origin: "potato",
					Tags: []string{
						"root vegetable",
						"vegetable",
					},
					Commonality: 5,
				},
			},
		},
		{
			Name:           "onion",
			PluralName:     "onions",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "onion",
					Origin: "onion",
					Tags: []string{
						"root vegetable",
						"vegetable",
					},
					Commonality: 5,
				},
			},
		},
		{
			Name:           "beet",
			PluralName:     "beets",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "beet",
					Origin: "beet",
					Tags: []string{
						"root vegetable",
						"vegetable",
					},
					Commonality: 5,
				},
			},
		},
		{
			Name:           "radish",
			PluralName:     "radishes",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "radish",
					Origin: "radish",
					Tags: []string{
						"root vegetable",
						"vegetable",
					},
					Commonality: 5,
				},
			},
		},
		{
			Name:           "yam",
			PluralName:     "yams",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "yam",
					Origin: "yam",
					Tags: []string{
						"root vegetable",
						"vegetable",
					},
					Commonality: 5,
				},
			},
		},
		{
			Name:           "carrot",
			PluralName:     "carrots",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "carrot",
					Origin: "carrot",
					Tags: []string{
						"root vegetable",
						"vegetable",
					},
					Commonality: 5,
				},
			},
		},
		{
			Name:           "eggplant",
			PluralName:     "eggplants",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "eggplant",
					Origin: "eggplant",
					Tags: []string{
						"vegetable",
					},
					Commonality: 5,
				},
			},
		},
		{
			Name:           "kohlrabi",
			PluralName:     "kohlrabis",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "kohlrabi",
					Origin: "kohlrabi",
					Tags: []string{
						"root vegetable",
						"vegetable",
					},
					Commonality: 5,
				},
			},
		},
		{
			Name:           "cabbage",
			PluralName:     "cabbages",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "cabbage",
					Origin: "cabbage",
					Tags: []string{
						"vegetable",
					},
					Commonality: 5,
				},
			},
		},
		{
			Name:           "lettuce",
			PluralName:     "lettuces",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "lettuce",
					Origin: "lettuce",
					Tags: []string{
						"vegetable",
					},
					Commonality: 5,
				},
			},
		},
		{
			Name:           "fiddlehead",
			PluralName:     "fiddleheads",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "fiddlehead",
					Origin: "fiddlehead",
					Tags: []string{
						"vegetable",
					},
					Commonality: 5,
				},
			},
		},
		{
			Name:           "broccoli",
			PluralName:     "broccolis",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "broccoli",
					Origin: "broccoli",
					Tags: []string{
						"vegetable",
					},
					Commonality: 5,
				},
			},
		},
		{
			Name:           "cauliflower",
			PluralName:     "cauliflowers",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "cauliflower",
					Origin: "cauliflower",
					Tags: []string{
						"vegetable",
					},
					Commonality: 5,
				},
			},
		},
		{
			Name:           "tomato",
			PluralName:     "tomatoes",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "tomato",
					Origin: "tomato",
					Tags: []string{
						"vegetable",
					},
					Commonality: 5,
				},
			},
		},
		{
			Name:           "bell pepper",
			PluralName:     "bell peppers",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "bell pepper",
					Origin: "bell pepper",
					Tags: []string{
						"vegetable",
					},
					Commonality: 5,
				},
			},
		},
		{
			Name:           "bok choy",
			PluralName:     "bok choys",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "bok choy",
					Origin: "bok choy",
					Tags: []string{
						"vegetable",
					},
					Commonality: 5,
				},
			},
		},
		{
			Name:           "leek",
			PluralName:     "leeks",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "leek",
					Origin: "leek",
					Tags: []string{
						"vegetable",
					},
					Commonality: 5,
				},
			},
		},
		{
			Name:           "brussel sprout",
			PluralName:     "brussel sprouts",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "brussel sprouts",
					Origin: "brussel sprout",
					Tags: []string{
						"vegetable",
					},
					Commonality: 5,
				},
			},
		},
		{
			Name:           "green bean",
			PluralName:     "green beans",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "green beans",
					Origin: "green bean",
					Tags: []string{
						"vegetable",
					},
					Commonality: 5,
				},
			},
		},
		{
			Name:           "black bean",
			PluralName:     "black beans",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "black beans",
					Origin: "black bean",
					Tags: []string{
						"legume",
						"vegetable",
					},
					Commonality: 5,
				},
			},
		},
		{
			Name:           "pinto bean",
			PluralName:     "pinto beans",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "pinto beans",
					Origin: "pinto bean",
					Tags: []string{
						"legume",
						"vegetable",
					},
					Commonality: 5,
				},
			},
		},
		{
			Name:           "lima bean",
			PluralName:     "lima beans",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "lima beans",
					Origin: "lima bean",
					Tags: []string{
						"legume",
						"vegetable",
					},
					Commonality: 5,
				},
			},
		},
		{
			Name:           "kidney bean",
			PluralName:     "kidney beans",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "kidney beans",
					Origin: "kidney bean",
					Tags: []string{
						"legume",
						"vegetable",
					},
					Commonality: 5,
				},
			},
		},
		{
			Name:           "corn",
			PluralName:     "corns",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "corn",
					Origin: "corn",
					Tags: []string{
						"vegetable",
					},
					Commonality: 5,
				},
			},
		},
		{
			Name:           "jalapeno",
			PluralName:     "jalapenos",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "jalapeno",
					Origin: "jalapeno",
					Tags: []string{
						"hot pepper",
						"vegetable",
					},
					Commonality: 5,
				},
			},
		},
		{
			Name:           "habanero",
			PluralName:     "habaneros",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "habanero",
					Origin: "habanero",
					Tags: []string{
						"hot pepper",
						"vegetable",
					},
					Commonality: 5,
				},
			},
		},
		{
			Name:           "serrano",
			PluralName:     "serranos",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "serrano",
					Origin: "serrano",
					Tags: []string{
						"hot pepper",
						"vegetable",
					},
					Commonality: 5,
				},
			},
		},
	}

	return vegetables
}
