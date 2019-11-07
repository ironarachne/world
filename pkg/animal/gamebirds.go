package animal

import (
	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/species"
)

func getGameBirds() []species.Species {
	birds := []species.Species{
		{
			Name:           "goose",
			PluralName:     "geese",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "goose eggs",
					Origin:       "goose",
					MainMaterial: "goose",
					Tags: []string{
						"egg",
					},
					Commonality: 4,
					Value:       1,
				},
				{
					Name:         "goose feathers",
					Origin:       "goose",
					MainMaterial: "goose",
					Tags: []string{
						"feather",
					},
					Commonality: 4,
					Value:       1,
				},
				{
					Name:         "goose",
					Origin:       "goose",
					MainMaterial: "goose",
					Tags: []string{
						"meat",
					},
					Commonality: 4,
					Value:       1,
				},
				{
					Name:         "goose liver",
					Origin:       "goose",
					MainMaterial: "goose",
					Tags: []string{
						"meat",
					},
					Commonality: 4,
					Value:       1,
				},
			},
			Tags: []string{
				"animal",
				"bird",
				"game bird",
				"herbivore",
			},
		},
		{
			Name:           "chicken",
			PluralName:     "chickens",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "chicken eggs",
					Origin:       "chicken",
					MainMaterial: "chicken",
					Tags: []string{
						"egg",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "chicken feathers",
					Origin:       "chicken",
					MainMaterial: "chicken",
					Tags: []string{
						"feather",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "chicken",
					Origin:       "chicken",
					MainMaterial: "chicken",
					Tags: []string{
						"meat",
					},
					Commonality: 4,
					Value:       1,
				},
				{
					Name:         "chicken liver",
					Origin:       "chicken",
					MainMaterial: "chicken",
					Tags: []string{
						"meat",
					},
					Commonality: 4,
					Value:       1,
				},
			},
			Tags: []string{
				"animal",
				"bird",
				"game bird",
				"herbivore",
			},
		},
		{
			Name:           "mudhen",
			PluralName:     "mudhens",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "mudhen eggs",
					Origin:       "mudhen",
					MainMaterial: "mudhen",
					Tags: []string{
						"egg",
					},
					Commonality: 4,
					Value:       1,
				},
				{
					Name:         "mudhen feathers",
					Origin:       "mudhen",
					MainMaterial: "mudhen",
					Tags: []string{
						"feather",
					},
					Commonality: 4,
					Value:       1,
				},
				{
					Name:         "mudhen",
					Origin:       "mudhen",
					MainMaterial: "mudhen",
					Tags: []string{
						"meat",
					},
					Commonality: 4,
					Value:       1,
				},
			},
			Tags: []string{
				"animal",
				"bird",
				"game bird",
				"herbivore",
			},
		},
		{
			Name:           "quail",
			PluralName:     "quails",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "quail eggs",
					Origin:       "quail",
					MainMaterial: "quail",
					Tags: []string{
						"egg",
					},
					Commonality: 4,
					Value:       1,
				},
				{
					Name:         "quail feathers",
					Origin:       "quail",
					MainMaterial: "quail",
					Tags: []string{
						"feather",
					},
					Commonality: 4,
					Value:       1,
				},
				{
					Name:         "quail",
					Origin:       "quail",
					MainMaterial: "quail",
					Tags: []string{
						"meat",
					},
					Commonality: 4,
					Value:       1,
				},
			},
			Tags: []string{
				"animal",
				"bird",
				"game bird",
				"herbivore",
			},
		},
		{
			Name:           "duck",
			PluralName:     "ducks",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "duck eggs",
					Origin:       "duck",
					MainMaterial: "duck",
					Tags: []string{
						"egg",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "duck feathers",
					Origin:       "duck",
					MainMaterial: "duck",
					Tags: []string{
						"feather",
					},
					Commonality: 4,
					Value:       1,
				},
				{
					Name:         "duck",
					Origin:       "duck",
					MainMaterial: "duck",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "duck liver",
					Origin:       "duck",
					MainMaterial: "duck",
					Tags: []string{
						"meat",
					},
					Commonality: 4,
					Value:       1,
				},
			},
			Tags: []string{
				"animal",
				"bird",
				"game bird",
				"herbivore",
			},
		},
	}

	return birds
}
