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
					Name:        "goose eggs",
					Origin:      "goose",
					Type:        "egg",
					Commonality: 4,
				},
				{
					Name:        "goose feathers",
					Origin:      "goose",
					Type:        "feather",
					Commonality: 4,
				},
				{
					Name:        "goose",
					Origin:      "goose",
					Type:        "meat",
					Commonality: 4,
				},
				{
					Name:        "goose liver",
					Origin:      "goose",
					Type:        "meat",
					Commonality: 4,
				},
			},
			Size: size.GetCategoryByName("tiny"),
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
					Name:        "chicken eggs",
					Origin:      "chicken",
					Type:        "egg",
					Commonality: 5,
				},
				{
					Name:        "chicken feathers",
					Origin:      "chicken",
					Type:        "feather",
					Commonality: 5,
				},
				{
					Name:        "chicken",
					Origin:      "chicken",
					Type:        "meat",
					Commonality: 4,
				},
				{
					Name:        "chicken liver",
					Origin:      "chicken",
					Type:        "meat",
					Commonality: 4,
				},
			},
			Size: size.GetCategoryByName("tiny"),
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
					Name:        "mudhen eggs",
					Origin:      "mudhen",
					Type:        "egg",
					Commonality: 4,
				},
				{
					Name:        "mudhen feathers",
					Origin:      "mudhen",
					Type:        "feather",
					Commonality: 4,
				},
				{
					Name:        "mudhen",
					Origin:      "mudhen",
					Type:        "meat",
					Commonality: 4,
				},
			},
			Size: size.GetCategoryByName("tiny"),
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
					Name:        "quail eggs",
					Origin:      "quail",
					Type:        "egg",
					Commonality: 4,
				},
				{
					Name:        "quail feathers",
					Origin:      "quail",
					Type:        "feather",
					Commonality: 4,
				},
				{
					Name:        "quail",
					Origin:      "quail",
					Type:        "meat",
					Commonality: 4,
				},
			},
			Size: size.GetCategoryByName("tiny"),
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
					Name:        "duck eggs",
					Origin:      "duck",
					Type:        "egg",
					Commonality: 5,
				},
				{
					Name:        "duck feathers",
					Origin:      "duck",
					Type:        "feather",
					Commonality: 4,
				},
				{
					Name:        "duck",
					Origin:      "duck",
					Type:        "meat",
					Commonality: 5,
				},
				{
					Name:        "duck liver",
					Origin:      "duck",
					Type:        "meat",
					Commonality: 4,
				},
			},
			Size: size.GetCategoryByName("tiny"),
		},
	}

	return birds
}
