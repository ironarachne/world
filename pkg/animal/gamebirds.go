package animal

import (
	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/size"
)

func getGameBirds() []Animal {
	birds := []Animal{
		{
			Name:           "goose",
			PluralName:     "geese",
			EatsAnimals:    false,
			EatsPlants:     true,
			IsAquatic:      false,
			IsMount:        false,
			IsPackAnimal:   false,
			IsScavenger:    false,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "goose eggs",
					Origin: "goose",
					Tags: []string{
						"egg",
					},
					Commonality: 4,
					Value:       1,
				},
				{
					Name:   "goose feathers",
					Origin: "goose",
					Tags: []string{
						"feather",
					},
					Commonality: 4,
					Value:       1,
				},
				{
					Name:   "goose",
					Origin: "goose",
					Tags: []string{
						"meat",
					},
					Commonality: 4,
					Value:       1,
				},
				{
					Name:   "goose liver",
					Origin: "goose",
					Tags: []string{
						"meat",
					},
					Commonality: 4,
					Value:       1,
				},
			},
			Size: size.GetCategoryByName("tiny"),
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
			EatsAnimals:    false,
			EatsPlants:     true,
			IsAquatic:      false,
			IsMount:        false,
			IsPackAnimal:   false,
			IsScavenger:    false,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "chicken eggs",
					Origin: "chicken",
					Tags: []string{
						"egg",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:   "chicken feathers",
					Origin: "chicken",
					Tags: []string{
						"feather",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:   "chicken",
					Origin: "chicken",
					Tags: []string{
						"meat",
					},
					Commonality: 4,
					Value:       1,
				},
				{
					Name:   "chicken liver",
					Origin: "chicken",
					Tags: []string{
						"meat",
					},
					Commonality: 4,
					Value:       1,
				},
			},
			Size: size.GetCategoryByName("tiny"),
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
			EatsAnimals:    false,
			EatsPlants:     true,
			IsAquatic:      false,
			IsMount:        false,
			IsPackAnimal:   false,
			IsScavenger:    false,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "mudhen eggs",
					Origin: "mudhen",
					Tags: []string{
						"egg",
					},
					Commonality: 4,
					Value:       1,
				},
				{
					Name:   "mudhen feathers",
					Origin: "mudhen",
					Tags: []string{
						"feather",
					},
					Commonality: 4,
					Value:       1,
				},
				{
					Name:   "mudhen",
					Origin: "mudhen",
					Tags: []string{
						"meat",
					},
					Commonality: 4,
					Value:       1,
				},
			},
			Size: size.GetCategoryByName("tiny"),
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
			EatsAnimals:    false,
			EatsPlants:     true,
			IsAquatic:      false,
			IsMount:        false,
			IsPackAnimal:   false,
			IsScavenger:    false,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "quail eggs",
					Origin: "quail",
					Tags: []string{
						"egg",
					},
					Commonality: 4,
					Value:       1,
				},
				{
					Name:   "quail feathers",
					Origin: "quail",
					Tags: []string{
						"feather",
					},
					Commonality: 4,
					Value:       1,
				},
				{
					Name:   "quail",
					Origin: "quail",
					Tags: []string{
						"meat",
					},
					Commonality: 4,
					Value:       1,
				},
			},
			Size: size.GetCategoryByName("tiny"),
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
			EatsAnimals:    false,
			EatsPlants:     true,
			IsAquatic:      true,
			IsMount:        false,
			IsPackAnimal:   false,
			IsScavenger:    false,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "duck eggs",
					Origin: "duck",
					Tags: []string{
						"egg",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:   "duck feathers",
					Origin: "duck",
					Tags: []string{
						"feather",
					},
					Commonality: 4,
					Value:       1,
				},
				{
					Name:   "duck",
					Origin: "duck",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:   "duck liver",
					Origin: "duck",
					Tags: []string{
						"meat",
					},
					Commonality: 4,
					Value:       1,
				},
			},
			Size: size.GetCategoryByName("tiny"),
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
