package animal

import (
	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/size"
)

func getReptiles() []Animal {
	reptiles := []Animal{
		Animal{
			Name:           "alligator",
			PluralName:     "alligators",
			AnimalType:     "reptile",
			EatsAnimals:    true,
			EatsPlants:     false,
			IsAquatic:      true,
			IsMount:        false,
			IsPackAnimal:   false,
			IsScavenger:    false,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 5,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "alligator hide",
					Origin:      "alligator",
					Type:        "hide",
					Commonality: 4,
				},
				resource.Resource{
					Name:        "alligator fangs",
					Origin:      "alligator",
					Type:        "teeth",
					Commonality: 4,
				},
				resource.Resource{
					Name:        "alligator",
					Origin:      "alligator",
					Type:        "meat",
					Commonality: 4,
				},
			},
			Size: size.GetCategoryByName("medium"),
		},
		Animal{
			Name:           "crocodile",
			PluralName:     "crocodiles",
			AnimalType:     "reptile",
			EatsAnimals:    true,
			EatsPlants:     false,
			IsAquatic:      true,
			IsMount:        false,
			IsPackAnimal:   false,
			IsScavenger:    false,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 5,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "crocodile hide",
					Origin:      "crocodile",
					Type:        "hide",
					Commonality: 4,
				},
				resource.Resource{
					Name:        "crocodile fangs",
					Origin:      "crocodile",
					Type:        "teeth",
					Commonality: 4,
				},
				resource.Resource{
					Name:        "crocodile",
					Origin:      "crocodile",
					Type:        "meat",
					Commonality: 4,
				},
			},
			Size: size.GetCategoryByName("medium"),
		},
		Animal{
			Name:           "monitor lizard",
			PluralName:     "monitor lizards",
			AnimalType:     "reptile",
			EatsAnimals:    true,
			EatsPlants:     false,
			IsAquatic:      false,
			IsMount:        false,
			IsPackAnimal:   false,
			IsScavenger:    false,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 5,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "monitor lizard hide",
					Origin:      "monitor lizard",
					Type:        "hide",
					Commonality: 4,
				},
				resource.Resource{
					Name:        "monitor lizard fangs",
					Origin:      "monitor lizard",
					Type:        "teeth",
					Commonality: 4,
				},
				resource.Resource{
					Name:        "monitor lizard",
					Origin:      "monitor lizard",
					Type:        "meat",
					Commonality: 4,
				},
				resource.Resource{
					Name:        "monitor lizard venom",
					Origin:      "monitor lizard",
					Type:        "venom",
					Commonality: 4,
				},
			},
			Size: size.GetCategoryByName("small"),
		},
	}

	return reptiles
}
