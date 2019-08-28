package animal

import (
	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/size"
)

func getRaptors() []Animal {
	birds := []Animal{
		{
			Name:           "eagle",
			PluralName:     "eagles",
			EatsAnimals:    true,
			EatsPlants:     false,
			IsAquatic:      false,
			IsMount:        false,
			IsPackAnimal:   false,
			IsScavenger:    true,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "eagle feathers",
					Origin: "eagle",
					Tags: []string{
						"feather",
					},
					Commonality: 4,
					Value:       10,
				},
			},
			Size: size.GetCategoryByName("tiny"),
			Tags: []string{
				"animal",
				"bird",
				"raptor",
				"carnivore",
			},
		},
		{
			Name:           "falcon",
			PluralName:     "falcons",
			EatsAnimals:    true,
			EatsPlants:     false,
			IsAquatic:      false,
			IsMount:        false,
			IsPackAnimal:   false,
			IsScavenger:    true,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "falcon feathers",
					Origin: "falcon",
					Tags: []string{
						"feather",
					},
					Commonality: 4,
					Value:       5,
				},
			},
			Size: size.GetCategoryByName("tiny"),
			Tags: []string{
				"animal",
				"bird",
				"raptor",
				"carnivore",
			},
		},
		{
			Name:           "hawk",
			PluralName:     "hawks",
			EatsAnimals:    true,
			EatsPlants:     false,
			IsAquatic:      false,
			IsMount:        false,
			IsPackAnimal:   false,
			IsScavenger:    true,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "hawk feathers",
					Origin: "hawk",
					Tags: []string{
						"feather",
					},
					Commonality: 4,
					Value:       5,
				},
			},
			Size: size.GetCategoryByName("tiny"),
			Tags: []string{
				"animal",
				"bird",
				"raptor",
				"carnivore",
			},
		},
	}

	return birds
}
