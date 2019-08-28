package animal

import (
	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/size"
)

func getCanines() []Animal {
	animals := []Animal{
		{
			Name:       "coyote",
			PluralName: "coyotes",
		},
		{
			Name:       "fox",
			PluralName: "foxes",
		},
		{
			Name:       "wolf",
			PluralName: "wolves",
		},
	}

	for _, a := range animals {
		a.AnimalType = "mammal"
		a.EatsAnimals = true
		a.EatsPlants = false
		a.IsMount = false
		a.IsPackAnimal = false
		a.IsScavenger = false
		a.MinHumidity = 0
		a.MaxHumidity = 10
		a.MinTemperature = 0
		a.MaxTemperature = 10
		a.Size = size.GetCategoryByName("medium")
		a.Resources = []resource.Resource{
			{
				Name:   a.Name + " hide",
				Origin: a.Name,
				Tags: []string{
					"hide",
				},
				Commonality: 4,
				Value:       4,
			},
			{
				Name:   a.Name + " fangs",
				Origin: a.Name,
				Tags: []string{
					"fangs",
					"teeth",
				},
				Commonality: 4,
				Value:       1,
			},
			{
				Name:   a.Name,
				Origin: a.Name,
				Tags: []string{
					"meat",
				},
				Commonality: 4,
				Value:       1,
			},
			{
				Name:   a.Name,
				Origin: a.Name,
				Tags: []string{
					"sinew",
				},
				Commonality: 5,
				Value:       1,
			},
		}
		a.Tags = []string{
			"animal",
			"canine",
			"carnivore",
			"hide",
		}
	}

	return animals
}
