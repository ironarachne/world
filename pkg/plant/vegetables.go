package plant

import (
	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/species"
)

func getVegetables() []species.Species {
	vegetables := []species.Species{
		{
			Name:           "potato",
			PluralName:     "potatoes",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "potato",
					Origin:       "potato",
					MainMaterial: "potato",
					Tags: []string{
						"root vegetable",
						"vegetable",
					},
					Commonality: 5,
					Value:       1,
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
					Name:         "onion",
					Origin:       "onion",
					MainMaterial: "onion",
					Tags: []string{
						"root vegetable",
						"vegetable",
					},
					Commonality: 5,
					Value:       1,
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
					Name:         "beet",
					Origin:       "beet",
					MainMaterial: "beet",
					Tags: []string{
						"root vegetable",
						"vegetable",
					},
					Commonality: 5,
					Value:       1,
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
					Name:         "radish",
					Origin:       "radish",
					MainMaterial: "radish",
					Tags: []string{
						"root vegetable",
						"vegetable",
					},
					Commonality: 5,
					Value:       1,
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
					Name:         "yam",
					Origin:       "yam",
					MainMaterial: "yam",
					Tags: []string{
						"root vegetable",
						"vegetable",
					},
					Commonality: 5,
					Value:       1,
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
					Name:         "carrot",
					Origin:       "carrot",
					MainMaterial: "carrot",
					Tags: []string{
						"root vegetable",
						"vegetable",
					},
					Commonality: 5,
					Value:       1,
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
					Name:         "eggplant",
					Origin:       "eggplant",
					MainMaterial: "eggplant",
					Tags: []string{
						"vegetable",
					},
					Commonality: 5,
					Value:       1,
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
					Name:         "kohlrabi",
					Origin:       "kohlrabi",
					MainMaterial: "kohlrabi",
					Tags: []string{
						"root vegetable",
						"vegetable",
					},
					Commonality: 5,
					Value:       1,
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
					Name:         "cabbage",
					Origin:       "cabbage",
					MainMaterial: "cabbage",
					Tags: []string{
						"vegetable",
					},
					Commonality: 5,
					Value:       1,
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
					Name:         "lettuce",
					Origin:       "lettuce",
					MainMaterial: "lettuce",
					Tags: []string{
						"vegetable",
					},
					Commonality: 5,
					Value:       1,
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
					Name:         "fiddlehead",
					Origin:       "fiddlehead",
					MainMaterial: "fiddlehead",
					Tags: []string{
						"vegetable",
					},
					Commonality: 5,
					Value:       1,
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
					Name:         "broccoli",
					Origin:       "broccoli",
					MainMaterial: "broccoli",
					Tags: []string{
						"vegetable",
					},
					Commonality: 5,
					Value:       1,
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
					Name:         "cauliflower",
					Origin:       "cauliflower",
					MainMaterial: "cauliflower",
					Tags: []string{
						"vegetable",
					},
					Commonality: 5,
					Value:       1,
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
					Name:         "tomato",
					Origin:       "tomato",
					MainMaterial: "tomato",
					Tags: []string{
						"vegetable",
					},
					Commonality: 5,
					Value:       1,
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
					Name:         "bell pepper",
					Origin:       "bell pepper",
					MainMaterial: "bell pepper",
					Tags: []string{
						"vegetable",
					},
					Commonality: 5,
					Value:       1,
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
					Name:         "bok choy",
					Origin:       "bok choy",
					MainMaterial: "bok choy",
					Tags: []string{
						"vegetable",
					},
					Commonality: 5,
					Value:       1,
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
					Name:         "leek",
					Origin:       "leek",
					MainMaterial: "leek",
					Tags: []string{
						"vegetable",
					},
					Commonality: 5,
					Value:       1,
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
					Name:         "brussel sprouts",
					Origin:       "brussel sprout",
					MainMaterial: "brussel sprout",
					Tags: []string{
						"vegetable",
					},
					Commonality: 5,
					Value:       1,
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
					Name:         "green beans",
					Origin:       "green bean",
					MainMaterial: "green bean",
					Tags: []string{
						"vegetable",
					},
					Commonality: 5,
					Value:       1,
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
					Name:         "black beans",
					Origin:       "black bean",
					MainMaterial: "black bean",
					Tags: []string{
						"legume",
						"vegetable",
					},
					Commonality: 5,
					Value:       1,
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
					Name:         "pinto beans",
					Origin:       "pinto bean",
					MainMaterial: "pinto bean",
					Tags: []string{
						"legume",
						"vegetable",
					},
					Commonality: 5,
					Value:       1,
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
					Name:         "lima beans",
					Origin:       "lima bean",
					MainMaterial: "lima bean",
					Tags: []string{
						"legume",
						"vegetable",
					},
					Commonality: 5,
					Value:       1,
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
					Name:         "kidney beans",
					Origin:       "kidney bean",
					MainMaterial: "kidney bean",
					Tags: []string{
						"legume",
						"vegetable",
					},
					Commonality: 5,
					Value:       1,
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
					Name:         "corn",
					Origin:       "corn",
					MainMaterial: "corn",
					Tags: []string{
						"vegetable",
					},
					Commonality: 5,
					Value:       1,
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
					Name:         "jalapeno",
					Origin:       "jalapeno",
					MainMaterial: "jalapeno",
					Tags: []string{
						"hot pepper",
						"vegetable",
					},
					Commonality: 5,
					Value:       1,
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
					Name:         "habanero",
					Origin:       "habanero",
					MainMaterial: "habanero",
					Tags: []string{
						"hot pepper",
						"vegetable",
					},
					Commonality: 5,
					Value:       1,
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
					Name:         "serrano",
					Origin:       "serrano",
					MainMaterial: "serrano",
					Tags: []string{
						"hot pepper",
						"vegetable",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
	}

	return vegetables
}
