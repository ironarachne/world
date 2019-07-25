package animal

import (
	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/size"
)

func getSnakes() []Animal {
	animals := []Animal{
		{
			Name:       "asp",
			PluralName: "asps",
		},
		{
			Name:       "cobra",
			PluralName: "cobras",
		},
		{
			Name:       "rattlesnake",
			PluralName: "rattlesnakes",
		},
		{
			Name:       "black mamba",
			PluralName: "black mambas",
		},
		{
			Name:       "coral snake",
			PluralName: "coral snakes",
		},
		{
			Name:       "boa constrictor",
			PluralName: "boa constrictors",
		},
	}

	for _, a := range animals {
		a.AnimalType = "snake"
		a.EatsAnimals = true
		a.EatsPlants = false
		a.IsMount = false
		a.IsPackAnimal = false
		a.IsScavenger = false
		a.MinHumidity = 0
		a.MaxHumidity = 10
		a.MinTemperature = 6
		a.MaxTemperature = 10
		a.Size = size.GetCategoryByName("tiny")
		a.Resources = []resource.Resource{
			{
				Name:   a.Name + " hide",
				Origin: a.Name,
				Tags: []string{
					"hide",
				},
				Commonality: 3,
			},
			{
				Name:   a.Name + " fangs",
				Origin: a.Name,
				Tags: []string{
					"fangs",
					"teeth",
				},
				Commonality: 3,
			},
			{
				Name:   a.Name,
				Origin: a.Name,
				Tags: []string{
					"meat",
				},
				Commonality: 3,
			},
			{
				Name:   a.Name + " heart",
				Origin: a.Name,
				Tags: []string{
					"meat",
				},
				Commonality: 1,
			},
			{
				Name:   a.Name + " venom",
				Origin: a.Name,
				Tags: []string{
					"venom",
				},
				Commonality: 3,
			},
		}
		a.Tags = []string{
			"animal",
			"reptile",
			"snake",
			"carnivore",
		}
	}

	return animals
}
