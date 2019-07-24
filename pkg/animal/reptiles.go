package animal

import (
	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/size"
)

func getReptiles() []Animal {
	reptiles := []Animal{
		{
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
				{
					Name:   "alligator hide",
					Origin: "alligator",
					Tags: []string{
						"hide",
					},
					Commonality: 4,
				},
				{
					Name:   "alligator fangs",
					Origin: "alligator",
					Tags: []string{
						"fangs",
						"teeth",
					},
					Commonality: 4,
				},
				{
					Name:   "alligator",
					Origin: "alligator",
					Tags: []string{
						"meat",
					},
					Commonality: 4,
				},
			},
			Size: size.GetCategoryByName("medium"),
		},
		{
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
				{
					Name:   "crocodile hide",
					Origin: "crocodile",
					Tags: []string{
						"hide",
					},
					Commonality: 4,
				},
				{
					Name:   "crocodile fangs",
					Origin: "crocodile",
					Tags: []string{
						"fangs",
						"teeth",
					},
					Commonality: 4,
				},
				{
					Name:   "crocodile",
					Origin: "crocodile",
					Tags: []string{
						"meat",
					},
					Commonality: 4,
				},
			},
			Size: size.GetCategoryByName("medium"),
		},
		{
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
				{
					Name:   "monitor lizard hide",
					Origin: "monitor lizard",
					Tags: []string{
						"hide",
					},
					Commonality: 4,
				},
				{
					Name:   "monitor lizard fangs",
					Origin: "monitor lizard",
					Tags: []string{
						"fangs",
						"teeth",
					},
					Commonality: 4,
				},
				{
					Name:   "monitor lizard",
					Origin: "monitor lizard",
					Tags: []string{
						"meat",
					},
					Commonality: 4,
				},
				{
					Name:   "monitor lizard venom",
					Origin: "monitor lizard",
					Tags: []string{
						"venom",
					},
					Commonality: 4,
				},
			},
			Size: size.GetCategoryByName("small"),
		},
	}

	return reptiles
}
