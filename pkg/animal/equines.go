package animal

import (
	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/species"
)

func getEquines() []species.Species {
	animals := []species.Species{
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
		a.MinHumidity = 0
		a.MaxHumidity = 10
		a.MinTemperature = 0
		a.MaxTemperature = 10
		a.Resources = []resource.Resource{
			{
				Name:         a.Name + " hide",
				Origin:       a.Name,
				MainMaterial: a.Name,
				Tags: []string{
					"hide",
				},
				Commonality: 3,
				Value:       1,
			},
			{
				Name:         a.Name + " milk",
				Origin:       a.Name,
				MainMaterial: a.Name,
				Tags: []string{
					"milk",
				},
				Commonality: 3,
				Value:       1,
			},
			{
				Name:         a.Name,
				Origin:       a.Name,
				MainMaterial: a.Name,
				Tags: []string{
					"meat",
				},
				Commonality: 3,
				Value:       1,
			},
			{
				Name:         a.Name,
				Origin:       a.Name,
				MainMaterial: a.Name,
				Tags: []string{
					"mount",
				},
				Commonality: 7,
				Value:       50,
			},
			{
				Name:         a.Name,
				Origin:       a.Name,
				MainMaterial: a.Name,
				Tags: []string{
					"pack animal",
				},
				Commonality: 5,
				Value:       5,
			},
		}
		a.Tags = []string{
			"animal",
			"equine",
			"herbivore",
			"hide",
		}
	}

	return animals
}
