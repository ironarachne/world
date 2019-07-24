package animal

import (
	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/size"
)

func getEquines() []Animal {
	animals := []Animal{
		{
			Name:       "horse",
			PluralName: "horses",
		},
		{
			Name:       "donkey",
			PluralName: "donkeys",
		},
		{
			Name:       "zebra",
			PluralName: "zebras",
		},
	}

	for _, a := range animals {
		a.AnimalType = "mammal"
		a.IsMount = true
		a.IsPackAnimal = true
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
				Commonality: 3,
			},
			{
				Name:   a.Name + " milk",
				Origin: a.Name,
				Tags: []string{
					"milk",
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
				Name:   a.Name,
				Origin: a.Name,
				Tags: []string{
					"mount",
				},
				Commonality: 7,
			},
			{
				Name:   a.Name,
				Origin: a.Name,
				Tags: []string{
					"pack animal",
				},
				Commonality: 5,
			},
		}
	}

	return animals
}
