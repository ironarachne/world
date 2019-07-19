package animal

import (
	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/size"
)

func getGameBirds() []Animal {
	birds := []Animal{
		Animal{
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
				resource.Resource{
					Name:        "goose eggs",
					Origin:      "goose",
					Type:        "egg",
					Commonality: 4,
				},
				resource.Resource{
					Name:        "goose feathers",
					Origin:      "goose",
					Type:        "feather",
					Commonality: 4,
				},
				resource.Resource{
					Name:        "goose",
					Origin:      "goose",
					Type:        "meat",
					Commonality: 4,
				},
				resource.Resource{
					Name:        "goose liver",
					Origin:      "goose",
					Type:        "meat",
					Commonality: 4,
				},
			},
			Size: size.GetCategoryByName("tiny"),
		},
		Animal{
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
				resource.Resource{
					Name:        "chicken eggs",
					Origin:      "chicken",
					Type:        "egg",
					Commonality: 5,
				},
				resource.Resource{
					Name:        "chicken feathers",
					Origin:      "chicken",
					Type:        "feather",
					Commonality: 5,
				},
				resource.Resource{
					Name:        "chicken",
					Origin:      "chicken",
					Type:        "meat",
					Commonality: 4,
				},
				resource.Resource{
					Name:        "chicken liver",
					Origin:      "chicken",
					Type:        "meat",
					Commonality: 4,
				},
			},
			Size: size.GetCategoryByName("tiny"),
		},
		Animal{
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
				resource.Resource{
					Name:        "mudhen eggs",
					Origin:      "mudhen",
					Type:        "egg",
					Commonality: 4,
				},
				resource.Resource{
					Name:        "mudhen feathers",
					Origin:      "mudhen",
					Type:        "feather",
					Commonality: 4,
				},
				resource.Resource{
					Name:        "mudhen",
					Origin:      "mudhen",
					Type:        "meat",
					Commonality: 4,
				},
			},
			Size: size.GetCategoryByName("tiny"),
		},
		Animal{
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
				resource.Resource{
					Name:        "quail eggs",
					Origin:      "quail",
					Type:        "egg",
					Commonality: 4,
				},
				resource.Resource{
					Name:        "quail feathers",
					Origin:      "quail",
					Type:        "feather",
					Commonality: 4,
				},
				resource.Resource{
					Name:        "quail",
					Origin:      "quail",
					Type:        "meat",
					Commonality: 4,
				},
			},
			Size: size.GetCategoryByName("tiny"),
		},
		Animal{
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
				resource.Resource{
					Name:        "duck eggs",
					Origin:      "duck",
					Type:        "egg",
					Commonality: 5,
				},
				resource.Resource{
					Name:        "duck feathers",
					Origin:      "duck",
					Type:        "feather",
					Commonality: 4,
				},
				resource.Resource{
					Name:        "duck",
					Origin:      "duck",
					Type:        "meat",
					Commonality: 5,
				},
				resource.Resource{
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
